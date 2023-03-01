package mintkey_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/crypto/bcrypt"
	"github.com/tendermint/tendermint/crypto"
	armor "github.com/tendermint/tendermint/crypto/armor"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys/mintkey"
)

func TestArmorUnarmorPrivKey(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	armor := mintkey.EncryptArmorPrivKey(priv, "passphrase")
	_, err := mintkey.UnarmorDecryptPrivKey(armor, "wrongpassphrase")
	require.Error(t, err)
	decrypted, err := mintkey.UnarmorDecryptPrivKey(armor, "passphrase")
	require.NoError(t, err)
	require.True(t, priv.Equals(decrypted))
}

func TestArmorUnarmorPubKey(t *testing.T) {
	// Select the encryption and storage for your cryptostore
	cstore := keys.NewInMemory()

	// Add keys and see they return in alphabetical order
	info, _, err := cstore.CreateMnemonic("Bob", keys.English, "passphrase", keys.Secp256k1)
	require.NoError(t, err)
	armor := mintkey.ArmorPubKeyBytes(info.GetPubKey().Bytes())
	pubBytes, err := mintkey.UnarmorPubKeyBytes(armor)
	require.NoError(t, err)
	pub, err := cryptoAmino.PubKeyFromBytes(pubBytes)
	require.NoError(t, err)
	require.True(t, pub.Equals(info.GetPubKey()))
}

func TestPdkdf2Encryption(t *testing.T) {
	priv := secp256k1.GenPrivKey()
	armor := mintkey.EncryptArmorPrivKey(priv, "passphrase")
	_, err := mintkey.UnarmorDecryptPrivKey(armor, "wrongpassphrase")
	require.Error(t, err)
	decrypted, err := mintkey.UnarmorDecryptPrivKey(armor, "passphrase")
	require.NoError(t, err)
	require.True(t, priv.Equals(decrypted))
	require.Equal(t, priv, decrypted)
}

func TestBcryptLegacyEncryption(t *testing.T) {
	priv := secp256k1.GenPrivKey()

	saltBytes := crypto.CRandBytes(16)
	key, _ := bcrypt.GenerateFromPassword(saltBytes, []byte("passphrase"), 12)
	key = crypto.Sha256(key) // get 32 bytes
	privKeyBytes := priv.Bytes()
	encBytes := xsalsa20symmetric.EncryptSymmetric(privKeyBytes, key)
	header := map[string]string{
		"kdf":  "bcrypt",
		"salt": fmt.Sprintf("%X", saltBytes),
	}
	armorString := armor.EncodeArmor("TENDERMINT PRIVATE KEY", header, encBytes)

	_, err := mintkey.UnarmorDecryptPrivKey(armorString, "wrongpassphrase")
	require.Error(t, err)
	decrypted, err := mintkey.UnarmorDecryptPrivKey(armorString, "passphrase")
	require.NoError(t, err)
	require.True(t, priv.Equals(decrypted))
	require.Equal(t, priv, decrypted)
}
