
## Notes

Objectives:

Reference: https://en.wikipedia.org/wiki/PKCS_11

Plugin structure / Hashicorp plugins (requires mTLS)
- Need to ensure correct TLS practices
- Maintain backwards compatibility through a deprecated interface

OBJECTIVE

- Generic support for internal+external signers
    - Generic Hardware Wallets (Ledger + others?)
    - Other Remote Signers plugins (as long as they implement the correct interface)
    - Add TPM 2.0 support       https://github.com/google/go-tpm
    - Initial basic PKCS#11
      https://docs.oasis-open.org/pkcs11/pkcs11-base/v3.0/os/pkcs11-base-v3.0-os.pdf
      https://docs.aws.amazon.com/cloudhsm/latest/userguide/pkcs11-library.html

## Developer Experience

### Development

### Open problems

#### Patches

There are cases where different approaches / upgrades have been incorporated without doing the proper changes through the whole
code base, resulting on multiple libraries that do the same thing slightly different, but if intendeed to fix this, it would involve
a major effort.

Example:
- Curently there are 2 versions of secp256k1 library at the same time
    - dcred: Which should be deprecated, still being used to calculate some checks, like knowing if a point is in the curve
    - secp256k1-voi: A new library that is constant time and is the one that should be used from now on.

#### Code duplication

There are pieces of code that do the same thing replicated over the whole sdk. It will make more difficult to migrate
since some of it is being used differently. Unifying this code enables developers to have one place of truth and also updating
it safe.

Examples:
- Key generation: Seckp256k1 key generation is repeated in different parts instead of being used trough one place
    - **crypto/hd/algo.go** contains a function that generates a private key with the received bytes. This function should be key agnostic, since multiple parts of the code depends on it.
    - **crypto/hd/hdpath,go** it parses the key directly from dcrd library instead of passing through the cosmos package
    - **ledger_mock.go**

#### Colateral damage

Working with Keys and cryptography in general is sensitive, some changes that might seem small impact the project in some unpredictable ways.
i.e: Seems like some parts of the project are using ed25519 keys which shouldn't be used yet, still, changing how the address are generated, break validators and the network
Some of the affected areas are:
- Network
    - Validators
    - Blocks
    - Tx
- Wallets
- Keys

---------------------------------------

Keyring vs Wallet management (same or split?)

- Wallet management
    - Åddress Encoding
        - bech32 / bech32m
        - HRP
    - Address to Keyring reference+links
    - Vanity address?
    - Check address is valid?

........................

- Keyring
    - Keep REFERENCES to signer? entities
        - Record (pub/priv keys)  --> reference  URL ledger://    hsm://....      priv://ddd
    - A SignerRecord?? instance should knows how to store itself
    - Armoring
        - Not always possible (OpenPGP support)
    - Manage mTLS keys???
      https://github.com/hashicorp/go-plugin/blob/main/mtls.go

......................

- Signer  (persistence)
    - From a keypair object
    - From some external reference (remote, etc.)
    - Retrieve instance from keyring
    - Example: Ledger may keep a pubkey reference that is checked. Locally it can be imported

- Generate/Derive
    - From hardware
    - From pure entropy (KDF)
    - From previous key material + metadata (BIP44)
    - Retrieve instance from keyring

- Verifier
    - Verify Signature + Digest
    - Validate pubkey
        - is on curve https://solanacookbook.com/references/keypairs-and-wallets.html#how-to-check-if-a-public-key-has-an-associated-private-key
          secp256k1

- Digest
    - F(Blob, hashFunction)


TX -> BLOB -> Digest -> F() -> signature

Tools/Helpers
- BIP39
- BIP32/BIP44

...............

Primitives (local)
- Key derivation functions (KDF)
- Signature Schemes
- Encryption
- Hashing
    - .... at least https://pkg.go.dev/crypto#Hash
- Password hashing (https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
    - argon2id > scrypt > bcrypt
    - ....


Open questions
- How to provide backwards compatibility?
    - Old keyring vs new Keyring?
    - Migration tool? RISKY
- How multisig should work?
- Support for threshold signature schemes?
- signature aggregation (schnoor + BLS)
- other reasons why this was forked?
- PQC
    - CRYSTALS-DILITHIUM
    - FALCON
    - SPHINCS+
    - https://csrc.nist.gov/Projects/post-quantum-cryptography/post-quantum-cryptography-standardization/round-3-submissions


I'm a bit worried that some decisions (like signing types (DIRECT, TEXTUAL, AUX...) ) that cosmos have outside the crypto module, will make it difficult to keep backwards compatibility while having a clean interface. But from a security perspective i thinkthe structure you suggests make sense.


More key types
https://docs.aws.amazon.com/cloudhsm/latest/userguide/pkcs11-key-types.html
- secp224r1, secp256r1, secp256k1, secp384r1
- rsa

---------------------

External references
https://solanacookbook.com/references/keypairs-and-wallets.html#how-to-generate-a-new-keypair

----------------------

codec
we cannot move/change this

hd
key derivation


mnemonic / seed
derivation (is optional)
priv key
pub key
