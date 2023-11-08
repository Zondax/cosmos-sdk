package provider

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Blob struct {
	data []byte
}

func NewBlob(data []byte) *Blob {
	return &Blob{data: data}
}

type SecureBlob struct {
	blob Blob
}

func NewSecureBlob(data []byte) *SecureBlob {
	return &SecureBlob{blob: Blob{data: data}}
}

func (b *SecureBlob) Zero() {
	for i := range b.blob.data {
		b.blob.data[i] = 0
	}
}

func (b *SecureBlob) UseAndZero(useKey func(*SecureBlob)) {
	defer b.Zero()
	useKey(b)
}

type ICryptoProviderMetadata interface {
	GetType() string
	GetName() string
	GetSlot() string
}

type ICryptoProvider interface {
	GetMetadata() ICryptoProviderMetadata
	ProtoReflect() protoreflect.Message

	GetKeys() (*Blob, *SecureBlob, error)
	GetSigner() (ISigner, error)
	GetVerifier() (IVerifier, error)
	GetCipher() (ICipher, error)
	GetHasher() (IHasher, error)
}

type SignerOption func(*SignerOptions)

type SignerOptions struct {
	options map[string]string
}

// This will be needed to select the desired signing modes in the sdk
func WithOption(key, value string) SignerOption {
	return func(signerOptions *SignerOptions) {
		signerOptions.options[key] = value
	}
}

type ISigner interface {
	Sign(data []byte, options ...SignerOption) ([]byte, error)
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
