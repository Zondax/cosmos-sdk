package keys

import "fmt"

type SampleKey struct {
	KeyType string
	Key     []byte

	a int
	b int
}

func (k SampleKey) String() string {
	return fmt.Sprintf("key type: %s - a: %d, b: %d", k.KeyType, k.a, k.b)
}

func (k SampleKey) Bytes() []byte {
	return k.Key
}

func (k SampleKey) Pubkey() PubKey {
	return k
}

func (k SampleKey) Address() []byte {
	return k.Bytes()
}
