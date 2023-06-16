package cmd

import (
	"context"
	reflectionv1beta1 "cosmossdk.io/api/cosmos/base/reflection/v1beta1"
	"cosmossdk.io/tools/rosetta"
	"crypto/tls"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	gogoproto "github.com/cosmos/gogoproto/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"io"
	"io/ioutil"
	math_bits "math/bits"
	"os/exec"
	"strings"
)

// RosettaCommand builds the rosetta root command given
// a protocol buffers serializer/deserializer
func RosettaCommand(ir codectypes.InterfaceRegistry, cdc codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rosetta",
		Short: "spin up a rosetta server",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := rosetta.FromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			protoCodec, ok := cdc.(*codec.ProtoCodec)
			if !ok {
				return fmt.Errorf("exoected *codec.ProtoMarshaler, got: %T", cdc)
			}
			conf.WithCodec(ir, protoCodec)

			err = rosetta.LoadPlugin(ir, "default") // These interfaces are common to all chains
			somethingToTest(ir)
			//err = rosetta.LoadPlugin(ir, conf.Blockchain)
			if err != nil {
				fmt.Printf("[Rosetta]- Error while loading the plugin: %s", err.Error())
				return err
			}

			rosettaSrv, err := rosetta.ServerFromConfig(conf)
			if err != nil {
				fmt.Printf("[Rosetta]- Error while creating server: %s", err.Error())
				return err
			}
			return rosettaSrv.Start()
		},
	}
	rosetta.SetFlags(cmd.Flags())

	return cmd
}

func openClient() (client *grpc.ClientConn, err error) {
	creds := credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS12,
	})
	endpoint := "grpc-evmos-ia.cosmosia.notional.ventures:443"

	client, err = grpc.Dial(endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("ERROR on client", err.Error())
		return nil, err
	}
	return client, err
}

func getInterfaceImplNames(c context.Context, client *grpc.ClientConn) (interfaceImplNames []string, err error) {
	cosmosReflectBetaClient := reflectionv1beta1.NewReflectionServiceClient(client)
	interfacesRes, err := cosmosReflectBetaClient.ListAllInterfaces(c, &reflectionv1beta1.ListAllInterfacesRequest{})

	if err == nil {
		for _, iface := range interfacesRes.InterfaceNames {
			implRes, err := cosmosReflectBetaClient.ListImplementations(c, &reflectionv1beta1.ListImplementationsRequest{
				InterfaceName: iface,
			})
			if err == nil {
				interfaceImplNames = append(interfaceImplNames, implementationMessageNameCleanup(implRes.GetImplementationMessageNames())...)
			}
		}
	}
	return interfaceImplNames, err
}

func implementationMessageNameCleanup(implMessages []string) (cleanImplMessages []string) {
	for _, implMessage := range implMessages {
		cleanImplMessages = append(cleanImplMessages, implMessage[1:])
	}

	return cleanImplMessages
}

func getFdset(client *grpc.ClientConn, c context.Context) (fdSet *descriptorpb.FileDescriptorSet, err error) {
	fdSet = &descriptorpb.FileDescriptorSet{}

	interfaceImplNames, err := getInterfaceImplNames(c, client)
	if err != nil {
		fmt.Println("[ERROR] on getting interfacesResponse implementations: ", err.Error())
		return fdSet, err
	}

	reflectClient, err := grpc_reflection_v1alpha.NewServerReflectionClient(client).ServerReflectionInfo(c)
	if err != nil {
		fmt.Println("[ERROR] reflect client", err.Error())
		return fdSet, err
	}

	fdMap := map[string]*descriptorpb.FileDescriptorProto{}
	waitListServiceRes := make(chan *grpc_reflection_v1alpha.ListServiceResponse)
	wait := make(chan struct{})
	go func() {
		for {
			in, err := reflectClient.Recv()
			if err == io.EOF {
				close(wait)
				return
			}
			if err != nil {
				fmt.Println("[ERROR] Reflection failed on reflectClient:", err.Error())
			}

			switch res := in.MessageResponse.(type) {
			case *grpc_reflection_v1alpha.ServerReflectionResponse_ErrorResponse:
				fmt.Println("[ERROR] Server reflection response:", res.ErrorResponse.String())
			case *grpc_reflection_v1alpha.ServerReflectionResponse_ListServicesResponse:
				waitListServiceRes <- res.ListServicesResponse
			case *grpc_reflection_v1alpha.ServerReflectionResponse_FileDescriptorResponse:
				for _, bz := range res.FileDescriptorResponse.FileDescriptorProto {
					fd := &descriptorpb.FileDescriptorProto{}
					err := proto.Unmarshal(bz, fd)
					if err != nil {
						fmt.Println("[ERROR] error happening while unmarshalling proto message", err.Error())
					}
					fdMap[fd.GetName()] = fd
				}
			}
		}
	}()

	if err = reflectClient.Send(&grpc_reflection_v1alpha.ServerReflectionRequest{
		MessageRequest: &grpc_reflection_v1alpha.ServerReflectionRequest_ListServices{},
	}); err != nil {
		fmt.Println("[ERROR] on ServerRefleciion services", err.Error())
	}

	listServiceRes := <-waitListServiceRes

	for _, response := range listServiceRes.Service {
		err = reflectClient.Send(&grpc_reflection_v1alpha.ServerReflectionRequest{
			MessageRequest: &grpc_reflection_v1alpha.ServerReflectionRequest_FileContainingSymbol{
				FileContainingSymbol: response.Name,
			},
		})
		if err != nil {
			fmt.Println("[ERROR] on ServerRefleciion services", err.Error())
		}
	}

	for _, msgName := range interfaceImplNames {
		fmt.Println("- ", msgName)
		err = reflectClient.Send(&grpc_reflection_v1alpha.ServerReflectionRequest{
			MessageRequest: &grpc_reflection_v1alpha.ServerReflectionRequest_FileContainingSymbol{
				FileContainingSymbol: msgName,
			},
		})
		if err != nil {
			fmt.Println("[ERROR] on getting interfaceImplNames", err.Error())
		}
	}

	if err = reflectClient.CloseSend(); err != nil {
		fmt.Println("[ERROR] on closing reflectClient", err.Error())
	}

	<-wait

	for _, descriptorProto := range fdMap {
		fdSet.File = append(fdSet.File, descriptorProto)
	}

	return fdSet, err
}

func somethingToTest(ir codectypes.InterfaceRegistry) {
	c := context.Background()
	client, err := openClient()
	if err != nil {
		fmt.Println("Error on open client:", err.Error())
	}

	fdSet, err := getFdset(client, c)
	if err != nil {
		fmt.Println("[ERROR] geting Fdset", err.Error())
	}

	for _, descriptorProto := range fdSet.File {
		if descriptorProto != nil {
			RegisterInterface(ir, descriptorProto)
		}
	}
	//bz, err := proto.Marshal(fdSet)
	//if err != nil {
	//	fmt.Println("[ERROR] masrhalling", err.Error())
	//}
	//
	//if err = os.WriteFile("filename", bz, 0o600); err != nil {
	//	fmt.Println("[ERROR] masrhalling", err.Error())
	//}
}

func convertToGogoproto(protoMessage proto.Message) gogoproto.Message {
	var newProtoMessage = new(extendedProtoMessage)
	newProtoMessage.Message = protoMessage
	return *newProtoMessage
}

func RegisterInterface(registry codectypes.InterfaceRegistry, fileDescriptor *descriptorpb.FileDescriptorProto) {
	name := "/" + strings.Replace(fileDescriptor.GetName(), "/", ".", -1)
	//fmt.Println(name)
	protoInterface := convertToGogoproto(fileDescriptor)
	//implementation := fileDescriptor.ProtoMessage
	//fmt.Println(">>>", name, " - ", protoInterface)
	//registry.RegisterInterface(name, &protoInterface, fileDescriptor)
	if name == "/ethermint.crypto.v1.ethsecp256k1.keys.proto" {
		fmt.Println(fileDescriptor.String())
		fmt.Println(fileDescriptor.MessageType[0])
		fmt.Println("***", fileDescriptor.GetMessageType())
		plsWork := convertToGogoproto(fileDescriptor.MessageType[0].ProtoReflect().Interface())
		fmt.Println("plsWork", plsWork)
		//protoFilePath := "testFile"
		//writeProtoFile(fileDescriptor, protoFilePath)
		//generateProtoCode(protoFilePath)
		registry.RegisterInterface(strings.Replace(fileDescriptor.GetName(), "/", ".", -1)[:len(fileDescriptor.GetName())-6], &protoInterface, &PubKey{})
		//registry.RegisterImplementations(&protoInterface, plsWork)
		fmt.Println("Interfaces", registry.ListAllInterfaces())
		fmt.Println("", registry.ListImplementations("ethermint.crypto.v1.ethsecp256k1.keys"))
	}
}

// PubKey defines a type alias for an ecdsa.PublicKey that implements
// Tendermint's PubKey interface. It represents the 33-byte compressed public
// key format.
type PubKey struct {
	// key is the public key in byte form
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *PubKey) String() string {
	//TODO implement me
	return ""
}

func (m *PubKey) Reset()      { *m = PubKey{} }
func (*PubKey) ProtoMessage() {}
func (*PubKey) Descriptor() ([]byte, []int) {
	return nil, []int{0}
}
func (m *PubKey) XXX_Unmarshal(b []byte) error {
	return nil
}
func (m *PubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return nil, nil
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *PubKey) XXX_Size() int {
	return m.Size()
}

func (m *PubKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}
func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func writeProtoFile(fd *descriptorpb.FileDescriptorProto, protoFilePath string) {
	// Method 1: Marshal to Text format and write to a file
	binaryData, err := proto.Marshal(fd)
	if err != nil {
		fmt.Println("Failed to marshal Proto File Descriptor:", err)
	}
	if err := ioutil.WriteFile(protoFilePath+".proto", binaryData, 0644); err != nil {
		fmt.Println("Failed to write Proto File Descriptor to file:", err)
	}
}

func generateProtoCode(protoFilePath string) {

	fmt.Println(protoFilePath + ".pb")
	// Generate the Go code using protoc command
	cmd := exec.Command("protoc", "--go_out=", protoFilePath+".proto")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to generate Go code:", err)
	}

	// Read the generated Go code from the temporary directory
	goCode, err := ioutil.ReadFile(protoFilePath + ".pb.go")
	if err != nil {
		fmt.Println("Failed to read generated Go code:", err)
	}

	// Print the generated Go code
	fmt.Println(string(goCode))
}

type extendedProtoMessage struct {
	proto.Message
}

func (e extendedProtoMessage) Reset() {
}
func (e extendedProtoMessage) String() string {
	return ""
}
func (e extendedProtoMessage) ProtoMessage() {
}
