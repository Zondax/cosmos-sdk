package typesNew

import (
	"crypto/rand"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/cypher"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestSigner(t *testing.T) {
	sampleSigner := signer.SampleSigner{}
	hash := make([]byte, 32)
	_, err := rand.Read(hash)
	fmt.Println("msg/hash:", hash)

	keyByteA := append(make([]byte, 10), 1)
	keyA := keys.SampleKey{KeyType: "A", Key: keyByteA}
	signatureA, err := sampleSigner.Sign(hash, keyA)
	require.NoError(t, err)
	require.NotEmpty(t, signatureA.Bytes())

	keyByteB := append(make([]byte, 10), 2)
	keyB := keys.SampleKey{KeyType: "B", Key: keyByteB}
	signatureB, err := sampleSigner.Sign(hash, keyB)
	require.NoError(t, err)
	require.NotEmpty(t, signatureB.Bytes())

	require.NotEqual(t, signatureA.Bytes(), signatureB.Bytes())
	fmt.Println("Key A:", keyA.Bytes())
	fmt.Println("Key B:", keyB.Bytes())
	fmt.Println("Signature A:", signatureA.Bytes())
	fmt.Println("Signature B:", signatureB.Bytes())

	sampleSigner2 := signer.SampleSignerB{}
	signatureA2, err := sampleSigner2.Sign(hash, keyB)
	fmt.Println("Signature A with Signer 2:", signatureA2.Bytes())
	require.NotEqual(t, signatureA.Bytes(), signatureA2.Bytes())
}

func TestCypher(t *testing.T) {
	msg := make([]byte, 32)
	rand.Read(msg)
	fmt.Println("msg:", msg)

	secret := []byte(fmt.Sprint("a_very_dark_secret"))
	fmt.Println("Secrets: ", secret)

	sampleCypher := cypher.SampleCypher{}
	encryptedData, _ := sampleCypher.Encrypt(msg, secret)
	fmt.Println("Encrypted data: ", encryptedData)

	decryptedData, _ := sampleCypher.Decrypt(msg, secret)
	fmt.Println("Decrypted data: ", decryptedData)
}
