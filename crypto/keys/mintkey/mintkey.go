package mintkey

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/keyerror"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/armor"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"

	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/tendermint/crypto/bcrypt"
	pdkdf2 "golang.org/x/crypto/pbkdf2"
)

const (
	blockTypePrivKey = "TENDERMINT PRIVATE KEY"
	blockTypeKeyInfo = "TENDERMINT KEY INFO"
	blockTypePubKey  = "TENDERMINT PUBLIC KEY"
)

// Make bcrypt security parameter var, so it can be changed within the lcd test
// Making the bcrypt security parameter a var shouldn't be a security issue:
// One can't verify an invalid key by maliciously changing the bcrypt
// parameter during a runtime vulnerability. The main security
// threat this then exposes would be something that changes this during
// runtime before the user creates their key. This vulnerability must
// succeed to update this to that same value before every subsequent call
// to the keys command in future startups / or the attacker must get access
// to the filesystem. However, with a similar threat model (changing
// variables in runtime), one can cause the user to sign a different tx
// than what they see, which is a significantly cheaper attack then breaking
// a bcrypt hash. (Recall that the nonce still exists to break rainbow tables)
// For further notes on security parameter choice, see README.md
var BcryptSecurityParameter = 12

//-----------------------------------------------------------------
// add armor

// Armor the InfoBytes
func ArmorInfoBytes(bz []byte) string {
	return armorBytes(bz, blockTypeKeyInfo)
}

// Armor the PubKeyBytes
func ArmorPubKeyBytes(bz []byte) string {
	return armorBytes(bz, blockTypePubKey)
}

func armorBytes(bz []byte, blockType string) string {
	header := map[string]string{
		"type":    "Info",
		"version": "0.0.0",
	}
	return armor.EncodeArmor(blockType, header, bz)
}

//-----------------------------------------------------------------
// remove armor

// Unarmor the InfoBytes
func UnarmorInfoBytes(armorStr string) (bz []byte, err error) {
	return unarmorBytes(armorStr, blockTypeKeyInfo)
}

// Unarmor the PubKeyBytes
func UnarmorPubKeyBytes(armorStr string) (bz []byte, err error) {
	return unarmorBytes(armorStr, blockTypePubKey)
}

func unarmorBytes(armorStr, blockType string) (bz []byte, err error) {
	bType, header, bz, err := armor.DecodeArmor(armorStr)
	if err != nil {
		return
	}
	if bType != blockType {
		err = fmt.Errorf("unrecognized armor type %q, expected: %q", bType, blockType)
		return
	}
	if header["version"] != "0.0.0" {
		err = fmt.Errorf("unrecognized version: %v", header["version"])
		return
	}
	return
}

//-----------------------------------------------------------------
// encrypt/decrypt with armor

// Encrypt and armor the private key.
func EncryptArmorPrivKey(privKey crypto.PrivKey, passphrase string) string {
	saltBytes, encBytes := encryptPrivKey(privKey, passphrase)
	header := map[string]string{
		"kdf":  "bcrypt",
		"salt": fmt.Sprintf("%X", saltBytes),
	}
	armorStr := armor.EncodeArmor(blockTypePrivKey, header, encBytes)
	return armorStr
}

// encrypt the given privKey with the passphrase using a randomly
// generated salt and the xsalsa20 cipher. returns the salt and the
// encrypted priv key.
func encryptPrivKey(privKey crypto.PrivKey, passphrase string) (saltBytes []byte, encBytes []byte) {
	saltBytes = crypto.CRandBytes(16)
	key := pdkdf2.Key([]byte(passphrase), saltBytes, BcryptSecurityParameter, 60, sha256.New)
	key = crypto.Sha256(key) // get 32 bytes
	privKeyBytes := privKey.Bytes()
	privKeyBytesHash := crypto.Sha256(privKeyBytes)
	privKeyBytes = append(privKeyBytes, privKeyBytesHash...) // Add own hash to differentiate it from old implementation
	return saltBytes, xsalsa20symmetric.EncryptSymmetric(privKeyBytes, key)
}

// Unarmor and decrypt the private key.
func UnarmorDecryptPrivKey(armorStr string, passphrase string) (crypto.PrivKey, error) {
	var privKey crypto.PrivKey
	blockType, header, encBytes, err := armor.DecodeArmor(armorStr)
	if err != nil {
		return privKey, err
	}
	if blockType != blockTypePrivKey {
		return privKey, fmt.Errorf("unrecognized armor type: %v", blockType)
	}
	if header["kdf"] != "bcrypt" {
		return privKey, fmt.Errorf("unrecognized KDF type: %v", header["KDF"])
	}
	if header["salt"] == "" {
		return privKey, fmt.Errorf("missing salt bytes")
	}
	saltBytes, err := hex.DecodeString(header["salt"])
	if err != nil {
		return privKey, fmt.Errorf("error decoding salt: %v", err.Error())
	}
	privKey, err = decryptPrivKey(saltBytes, encBytes, passphrase)
	return privKey, err
}

func decryptPrivKey(saltBytes []byte, encBytes []byte, passphrase string) (privKey crypto.PrivKey, err error) {
	key := pdkdf2.Key([]byte(passphrase), saltBytes, BcryptSecurityParameter, 60, sha256.New)
	key = crypto.Sha256(key) // Get 32 bytes

	privateBytes, err := decryptSymmetric(encBytes, key)
	if err == nil || len(privateBytes) > 32 {
		decryptedBytes := privateBytes[:len(privateBytes)-32]
		decryptedBytesHash := privateBytes[len(privateBytes)-32:] //SHA-256 hash is 32 bytes
		//If the decrypted hash doesn't match the privateBytes hash, then we are working with the old bcrypt algorithm
		if !bytes.Equal(crypto.Sha256(decryptedBytes), decryptedBytesHash) {
			decryptedBytes, err = decryptPrivKeyLegacy(saltBytes, encBytes, passphrase)
		}
	} else {
		privateBytes, err = decryptPrivKeyLegacy(saltBytes, encBytes, passphrase)
	}

	if err != nil {
		return privKey, err
	}
	privKey, err = cryptoAmino.PrivKeyFromBytes(privateBytes)
	return privKey, err
}

func decryptPrivKeyLegacy(saltBytes []byte, encBytes []byte, passphrase string) (decryptedBytes []byte, err error) {
	key, err := bcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)
	if err != nil {
		cmn.Exit("error generating bcrypt key from passphrase: " + err.Error())
	}
	return decryptSymmetric(encBytes, key)
}

func decryptSymmetric(encBytes []byte, key []byte) (decryptedBytes []byte, err error) {
	decryptedBytes, err = xsalsa20symmetric.DecryptSymmetric(encBytes, key)
	if err != nil && err.Error() == "Ciphertext decryption failed" {
		return decryptedBytes, keyerror.NewErrWrongPassword()
	} else if err != nil {
		return decryptedBytes, err
	}
	return decryptedBytes, nil
}
