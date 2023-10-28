package localSecp256k1

import (
	"crypto/rand"
	"cryptoV2/provider"
	"io"
	"testing"
)

func TestBuilder_FromSeed(t *testing.T) {
	tests := []struct {
		name    string
		want    provider.ICryptoProvider
		wantErr bool
	}{
		{
			name:    "normal",
			want:    &LocalSecp256K1{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privKeyBytes := [32]byte{}
			r := rand.Reader
			_, err := io.ReadFull(r, privKeyBytes[:])
			b := Builder{}
			got, err := b.FromSeed(privKeyBytes[:])
			if (err != nil) != tt.wantErr {
				t.Errorf("FromSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("nil provider")
				return
			}
		})
	}
}
