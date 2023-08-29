# Design Nodes

## Main layers

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

...............

Test 

## Open questions
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
