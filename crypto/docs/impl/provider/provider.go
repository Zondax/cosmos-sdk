package provider

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ICryptoProviderMetadata interface {
	GetTypeUUID() string
	GetName() string
}

type ICryptoProvider interface {
	GetMetadata() ICryptoProviderMetadata
	ProtoReflect() protoreflect.Message

	GetKeys() ([]byte, []byte, error)
	GetSigner() (ISigner, error)
	GetVerifier() (IVerifier, error)
	GetCipher() (ICipher, error)
	GetHasher() (IHasher, error)
}

type ISigner interface {
	Sign([]byte) ([]byte, error)
}

type IVerifier interface {
	Verify([]byte, []byte) (bool, error)
}

type ICipher interface {
	Encrypt(message []byte) ([]byte, error)
	Decrypt(encryptedMessage []byte) ([]byte, error)
}

type IHasher interface {
	Hash(input []byte) []byte
	CanHashIncrementally() bool
}
