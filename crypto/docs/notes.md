# Notes

* Cosmos SDK codebase - There are many TODOs everywhere in the code base. We can get a lot of tasks by fixing those. We should collect with grep, create issues and triage

* ADR does not take into account higher level components in the design.


----

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
- Developer Experience
- Development
- Open problems

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
