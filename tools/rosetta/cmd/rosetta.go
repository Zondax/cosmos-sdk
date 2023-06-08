package cmd

import (
	"context"
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1beta1 "cosmossdk.io/api/cosmos/base/reflection/v1beta1"
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"cosmossdk.io/tools/rosetta"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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
			somethingToTest()
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
	endpoint := "evmos-grpc.lavenderfive.com:443"

	client, err = grpc.Dial(endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("ERROR on client", err.Error())
		return nil, err
	}
	return client, err
}

func getFdset(client *grpc.ClientConn, c context.Context) (fdSet *descriptorpb.FileDescriptorSet, err error) {
	fdSet = &descriptorpb.FileDescriptorSet{}
	//reflectionClient := reflectionv1.NewReflectionServiceClient(client)
	//fdRes, err := reflectionClient.FileDescriptors(c, &reflectionv1.FileDescriptorsRequest{})

	var interfaceImplNames []string
	cosmosReflectBetaClient := reflectionv1beta1.NewReflectionServiceClient(client)
	interfacesRes, err := cosmosReflectBetaClient.ListAllInterfaces(c, &reflectionv1beta1.ListAllInterfacesRequest{})

	if err == nil {
		for _, iface := range interfacesRes.InterfaceNames {
			implRes, err := cosmosReflectBetaClient.ListImplementations(c, &reflectionv1beta1.ListImplementationsRequest{
				InterfaceName: iface,
			})
			if err == nil {
				interfaceImplNames = append(interfaceImplNames, implRes.ImplementationMessageNames...)
			}
		}
	} else {
		fmt.Println("[ERROR] on getting interfacesResponse implementations: ", err.Error())
	}

	reflectClient, err := grpc_reflection_v1alpha.NewServerReflectionClient(client).ServerReflectionInfo(c)
	if err != nil {
		fmt.Println("ERROR reflect client", err.Error())
	}

	fdMap := map[string]*descriptorpb.FileDescriptorProto{}
	waitListServiceRes := make(chan *grpc_reflection_v1alpha.ListServiceResponse)
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := reflectClient.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				fmt.Println("[ERROR] Reflection failed on reflectClient:", err.Error())
			}

			switch res := in.MessageResponse.(type) {
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
		fmt.Println("[ERROR] on ServerRefleciio services", err.Error())
	}

	listServiceRes := <-waitListServiceRes

	for _, response := range listServiceRes.Service {
		err = reflectClient.Send(&grpc_reflection_v1alpha.ServerReflectionRequest{
			MessageRequest: &grpc_reflection_v1alpha.ServerReflectionRequest_FileContainingSymbol{
				FileContainingSymbol: response.Name,
			},
		})
		if err != nil {
			fmt.Println("[ERROR] on ServerRefleciio services", err.Error())
		}
	}

	for _, msgName := range interfaceImplNames {
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
		//return nil, err
	}

	<-waitc

	// we loop through all the file descriptor dependencies to capture any file descriptors we haven't loaded yet
	cantFind := map[string]bool{}
	//for {
	//	missing := missingFileDescriptors(fdMap, cantFind)
	//	if len(missing) == 0 {
	//		break
	//	}
	//
	//	err = addMissingFileDescriptors(ctx, client, fdMap, missing)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// mark all deps that we aren't able to resolve as can't find, so we don't keep looping and get a 429 error
	//	for _, dep := range missing {
	//		if fdMap[dep] == nil {
	//			cantFind[dep] = true
	//		}
	//	}
	//}

	for dep := range cantFind {
		fmt.Printf("Warning: can't find %s.\n", dep)
	}

	for _, descriptorProto := range fdMap {
		fdSet.File = append(fdSet.File, descriptorProto)
	}
	fmt.Println("AQUI LLEGAMOS")

	return fdSet, err
}

func somethingToTest() {
	fmt.Println("1 - Setup client")
	c := context.Background()
	client, err := openClient()
	if err != nil {
		fmt.Println("Error on open client:", err.Error())
	}

	fdSet, err := getFdset(client, c)

	bz, err := proto.Marshal(fdSet)
	if err != nil {
		fmt.Println("Error masrhalling", err.Error())
	}

	if err = os.WriteFile("filename", bz, 0o600); err != nil {
		fmt.Println("Error masrhalling", err.Error())
	}

	protofiles, err := protodesc.FileOptions{AllowUnresolvable: true}.NewFiles(fdSet)
	if err != nil {
		fmt.Println("error building protoregistry.Files: %w", err.Error())
	}

	client, err = openClient()
	if err != nil {
		fmt.Println("Error on open client:", err.Error())
	}

	autocliQueryClient := autocliv1.NewQueryClient(client)
	appOptsRes, err := autocliQueryClient.AppOptions(c, &autocliv1.AppOptionsRequest{})
	if err != nil {
		appOptsRes = guessAutocli(protofiles)
	}

	bz, err = proto.Marshal(appOptsRes)
	if err != nil {
		fmt.Println("Error on marshalling appOptsRes:", err.Error())
	}

	if err := os.WriteFile("appOptsFilename", bz, 0o600); err != nil {
		fmt.Println("Error while writing the files:", err.Error())
	}

	//_ := appOptsRes.ModuleOptions
	//} else {
	//	bz, err := os.ReadFile(appOptsFilename)
	//	if err != nil {
	//		return err
	//	}
	//
	//	var appOptsRes autocliv1.AppOptionsResponse
	//	if err := proto.Unmarshal(bz, &appOptsRes); err != nil {
	//		return err
	//	}
	//
	//	c.ModuleOptions = appOptsRes.ModuleOptions
	//}
}

func guessAutocli(files *protoregistry.Files) *autocliv1.AppOptionsResponse {
	fmt.Printf("This chain does not support autocli directly yet. Using some default mappings in the meantime to support a subset of the available services.\n")
	res := map[string]*autocliv1.ModuleOptions{}
	files.RangeFiles(func(descriptor protoreflect.FileDescriptor) bool {
		services := descriptor.Services()
		n := services.Len()
		for i := 0; i < n; i++ {
			service := services.Get(i)
			serviceName := service.FullName()
			mapping, ok := defaultAutocliMappings[serviceName]
			if ok {
				parts := strings.Split(mapping, " ")
				numParts := len(parts)
				if numParts < 2 || numParts > 3 {
					fmt.Printf("Warning: bad mapping %q found for %q\n", mapping, serviceName)
					continue
				}

				modOpts := res[parts[0]]
				if modOpts == nil {
					modOpts = &autocliv1.ModuleOptions{}
					res[parts[0]] = modOpts
				}

				switch parts[1] {
				case "query":
					if modOpts.Query == nil {
						modOpts.Query = &autocliv1.ServiceCommandDescriptor{
							SubCommands: map[string]*autocliv1.ServiceCommandDescriptor{},
						}
					}
					if numParts == 3 {
						modOpts.Query.SubCommands[parts[2]] = &autocliv1.ServiceCommandDescriptor{Service: string(serviceName)}
					} else {
						modOpts.Query.Service = string(serviceName)
					}
				case "tx":
					if modOpts.Tx == nil {
						modOpts.Tx = &autocliv1.ServiceCommandDescriptor{
							SubCommands: map[string]*autocliv1.ServiceCommandDescriptor{},
						}
					}
					if numParts == 3 {
						modOpts.Tx.SubCommands[parts[2]] = &autocliv1.ServiceCommandDescriptor{Service: string(serviceName)}
					} else {
						modOpts.Tx.Service = string(serviceName)
					}
				default:
					fmt.Printf("Warning: bad mapping %q found for %q\n", mapping, serviceName)
					continue
				}
			}
		}
		return true
	})

	return &autocliv1.AppOptionsResponse{ModuleOptions: res}
}

var defaultAutocliMappings = map[protoreflect.FullName]string{
	"cosmos.auth.v1beta1.Query":         "auth query",
	"cosmos.authz.v1beta1.Query":        "authz query",
	"cosmos.bank.v1beta1.Query":         "bank query",
	"cosmos.distribution.v1beta1.Query": "distribution query",
	"cosmos.evidence.v1.Query":          "evidence query",
	"cosmos.feegrant.v1beta1.Query":     "feegrant query",
	"cosmos.gov.v1.Query":               "gov query",
	"cosmos.gov.v1beta1.Query":          "gov query v1beta1",
	"cosmos.group.v1.Query":             "group query",
	"cosmos.mint.v1beta1.Query":         "mint query",
	"cosmos.params.v1beta1.Query":       "params query",
	"cosmos.slashing.v1beta1.Query":     "slashing query",
	"cosmos.staking.v1beta1.Query":      "staking query",
	"cosmos.upgrade.v1.Query":           "upgrade query",
}
