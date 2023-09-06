# ADR 777: Cryptography v2

## Change log

* {date}: {change log}
* ????-??-??: Initial Draft

## Status

{DRAFT | PROPOSED} Not Implemented

## Abstract

TODO: Do this as the end

## Context

* Currently, there is no ADR providing a comprehensive description of the cryptographic module in the Cosmos SDK.
* There have been multiple requests for a more flexible and extensible approach to cryptography, address management, and more.
* Several open issues require significant changes for resolution.
* Similar efforts have been undertaken in the past concerning runtime modules.
* Existing signing types outside of the crypto module may pose challenges to backward compatibility while striving for a clean interface.
* Security implications must be considered during the module's redesign.

### Proposed architecture

The architecture objectives that define our design are based on the following concepts:

* **Modularity**: Users should be able to use only what they need instead of getting the whole module, keeping projects decoupled lightweight. 
* **Simplicity**: The proposal follows a modular API architecture, abstracting complex behaviour and defining a clear interaction between modules.
* **Extensibility**: Adding new features as key types, signing algorithms, etc. Has been made easier, in order to avoid forks and
promote users to come up with their own implementations of the interfaces, which should instantly work with the rest of the module without
modifications.

### **Modules**

Modules aim to encapsulate behaviours and to provide simple interface to extend and reuse.

These are the following reasons to use modules over packages:
- **Improved dependency management**: Modules have a built-in dependency management system that makes it easy to track and manage 
  dependencies.
- **Lightweight**: Users of the SDK could decide what modules to use, keeping simpler dependencies.
- **Simplified development**: Extending a module with a reduced scope allows users to handle their own implementations easier. 

```mermaid
classDiagram

Hasher <|-- CryptoProvider
CryptoCypher <|-- CryptoProvider

PubKey -- Verifier

PubKey <|-- PrivKey
PubKey <|-- Generator

PrivKey <|-- Signer
PrivKey <|-- Generator

Signature <|-- Verifier
Signature <|-- Signer

Signer <|-- CryptoProvider
Verifier <|-- CryptoProvider
```

#### Crypto provider

A Crypto provider is the middleware object that handles the interaction with different instanced modules, A provider could be seen as a controller. 

It is created through a factory / builder method which contains all the implementations of the required interfaces. It aims to encapsulate
the api of the crypto modules in one place. 

```mermaid
classDiagram
Hasher <|-- CryptoProvider
Cypher <|-- CryptoProvider
Keys <|-- CryptoProvider
Signer <|-- CryptoProvider
Verifier <|-- CryptoProvider

KeyRing --|> CryptoProvider
```

```go
type CryptoProvider interface {
 GetSigner() (signer. Signer, error)
 GetVerifier() (verifier. Verifier, error)
 GetCipher() (cypher.Cipher, error)
 GetHasher() (Hasher, error)
}
```

#### **Keyring**

Keyring serves as a middleware between ledgers and the cosmos-sdk modules. It performs operations of storing, retrieving data 
and allowing ledgers to perform operations such as signing and verifying.

Keyring will get the information related to keys contained in a **secure item** through a **secure storage** object. Trough 
keyring, Users can register their desired crypto providers to use and their respective storages implementations, It then 
matches the UUID to the registered interfaces. Returning the specific provider for interacting with a key stored in a secure
storage.

```mermaid
classDiagram

SecuredStorage <|-- KeyRing

SecuredStorage : Get(key string) (SecureItem, error)
SecuredStorage : Set(key string, item SecureItem) error
SecuredStorage : Delete(key string) error
SecuredStorage : List() ([]string, error)

SecureItem <|-- SecuredStorage
SecureItemMetadata <|-- SecureItem

Key <|-- SecureItem

Wallet --|> KeyRing
```

##### SecureItem

A secure item consists of a specific Key instance registered on keyring. It also contains the metadata for interacting with the respective
key. Such metadata could be encoding, encryption, etc.

##### SecureStorage

A secures storage represents a vault where one or more **secure items** could be stored. In order to get a secure item, the user
will have to interact with a secure storage before getting the secure item. Secure storage knows how and where to interact with
in order to get the keys.

```go
type SecureStorageSourceMetadata struct {
    Type string
    Name string
}

type SecureStorageSourceConfig struct {
    Metadata SecureStorageSourceMetadata
    Config   any // specific config for the desired backend, if necessary
}

type SecureStorageBuilder func(SecureStorageSourceConfig) (SecureStorage, error)

type SecureStorage interface {
  // Build builds the corresponding secure storage backend
  Build(SecureStorageSourceConfig) (SecureStorage, error)
  
  // Get returns the SecureItem matching the key or ErrKeyNotFound
  Get(string) (secure_item.SecureItem, error)
  // GetMetadata returns the metadata field of the SecureItem
  GetMetadata(string) (secure_item.SecureItemMetadata, error)
  // Set stores the SecureItem on the backend
  Set(string, secure_item.SecureItem) error
  // Remove removes the SecureItem matching the key
  Remove(string) error
  // Keys returns a slice of all keys stored on the backend
  Keys() ([]string, error)
}
```

#### **Wallet**

The wallet API contains the blockchain specific use cases of the crypto module. It is responsible for:
- Signing and Verifying messages.
- generating addresses out of keys

Since wallet interacts with the user keys, it contains an instance of the Keyring, it is also where the blockchain specific
logic should reside.

##### Blob

This is a wrapper for the widely used `[]byte` type that is used when handling binary data. Since crypto module handles sensitive information,
the objective is to provide some extra security capabilities around such type as:
- Zeroing values after a read operation.
- Securely handling data.

These blob structures would be passed within components of the crypto module. For example: Signature information

#### **Keys**

A key object is responsible for containing the **BLOB** key information. Keys might not be passed through functions and it is 
suggested to interact through crypto providers to limit the exposure to vulnerabilities. 

```mermaid
classDiagram
  PubKey <|-- PrivKey
  PubKey : Address() string
  PubKey : Key 

  PrivKey : PubKey() PubKey
  PrivKey : key
```

Base Key struct

```go
type KeyStruct struct {
 key Blob
}
```

Base key interface (common to private and public keys)

```go
type BaseKey interface {
 String() string
 Bytes() Blob
}
```

The generator module is responsible for generating such keys.

##### PubKey

```go
type PubKey interface {
 BaseKey
}
```

##### PrivKey

```go
type PrivKey interface {
 BaseKey
 Pubkey() PubKey //Generate a public key out of a private key
}
```

#### Signatures

A signature consists of a message/hash signed by one or multiple private keys. The main objective is to Authenticate a message signer 
trough their public key.

```go
type Signature struct {
 data Blob
}
```

##### Signer

Interface responsible for Signing a message and returning the generated Signature. It is an algorithm tied to a family of keys. 

```go
type Signer interface {
 Sign(Blob, PrivKey) (Signature, error)
}
```

##### Verifier

Verifies if given a message belongs to a public key by validating against it's respective signature.

```go
type Verifier interface {
 Verify(Blob, Signature, PubKey) (bool, error)
}
```

#### Cipher

A cipher is an api for encryption and decryption of data. Given a message it should operate through a secret.

```go
type Cipher interface {
    Encryptor
    Decryptor
}
```

##### Encryptor

Given a message and a secret, ciphers such message according to the implemented algorithm.

```go
type Encryptor interface {
    Encrypt(message Blob, secret Blob) (encryptedMessage Blob, error)
}
```

##### Decryptor

Given a Ciphered message and a secret, decrypts such message according to the implemented algorithm.

```go
type Decryptor interface {
    Decrypt(message Blob, secret Blob) (decryptedMessage Blob, error)
}
```

#### Hasher

This module contains the different hashing algorithms and conventions agreed on this matter. 

```go
type Hasher interface {
 Hash(input Blob) Blob
 CanHashIncrementally() bool
}
```

#### Hasher

This module contains the different hashing algorithms and conventions agreed on this matter.

#### Codec

This module will continue to register types and interfaces from the module according to the interface registry structure.

#### Module structure

Crypto module structure would look similar to this
- codec
- cipher
  - encryption
  - decryption
  - hashing
- docs
- keyring
  - secureItem
  - secureStorage
- keys
- provider
- signature
  - signer
  - verifier
- wallet
- types

### Overview of the whole design

```mermaid
classDiagram

SecureItem <|-- SecuredStorage
SecureItemMetadata <|-- SecureItem

SecuredStorage : Get(key string) (SecureItem, error)
SecuredStorage : Set(key string, item SecureItem) error
SecuredStorage : Delete(key string) error
SecuredStorage : List() ([]string, error)

CryptoCypher <|-- CryptoProvider
Hasher <|-- CryptoProvider
Signer <|-- CryptoProvider
Verifier <|-- CryptoProvider
CryptoProvider <|-- SecretElement
CryptoProvider <|-- SecretKeyPair
CryptoProvider <|-- SecureElement
CryptoProvider : GetCypher() (CryptoCypher, error)
CryptoProvider : GetHasher() (Hasher, error)
CryptoProvider : GetRandom() (Random, error)
CryptoProvider : GetRandomBytes(size int) ([]byte, error)

Hasher : Hash() ([]byte, error)
Hasher : HashString() (string, error)

SecureItem  <|-- CryptoCypher
CryptoCypher : Encrypt(data Blob) (Blob, error)
CryptoCypher : Decrypt(data Blob) (Blob, error)

Blob <|-- Digest
Blob <|-- SecureItem
Blob <|-- Hasher
Blob <|-- CryptoCypher
Blob : Bytes()
Blob : ReadBlob() ([]byte, error)
Blob : Wipe()


BaseKey <|-- PubKey
BaseKey <|-- PrivKey
BaseKey : String() string
BaseKey : Bytes() []byte

Signature <|-- Verifier
Signature <|-- Signer
Signature : Bytes() []byte

Signer
Signer : Sign(hash []byte, key PrivKey) (Signature, error)

Verifier
Verifier : Verify(hash []byte, sig Signature, key PubKey) (bool, error)

PubKey <|-- PrivKey
PubKey : Address() string
PubKey <|-- Generator
PubKey <|-- Verifier
PubKey <|-- SecretKeyPair

PrivKey : PubKey() PubKey
PrivKey <|-- Generator
PrivKey <|-- Signer
PrivKey <|-- SecretKeyPair

Generator
Generator : GenerateKey() (PrivKey, error)
Generator : DeriveKey(pb PubKey) (PrivKey, error)


SecretElement
SecretElement <|-- LedgerDevice
```

### Use cases

In the following scenario the USER uses an external ledger to:

1. Load stored ledger information, and using one of the keyringRecord which represents the specific LEDGER
2. Encrypt a message trough CYPHER
3. Generate an asymetric key trough GENERATOR
4. Signing a message trough the SIGNER and the Private key
5. Verifying a message trough the VERIFIER and the Public Key

**Flow overview**

***Initialization***

```mermaid
sequenceDiagram
    participant Keyring
    participant ConfigLoader
    participant SecureStorage
    
    SecureStorage->>Keyring: RegisterStorageSource()

```



```mermaid
sequenceDiagram
    participant Keyring
    participant KeyringRecord
    participant CryptoProvider

    Keyring->>Keyring: New()
    Keyring->>KeyringRecord: GetRecords()
    KeyringRecord->>KeyringRecord: Restore()
    KeyringRecord->>Keyring: KeyringRecord Instance

    Keyring->>KeyringRecord: EncryptMessage()
    KeyringRecord->>CryptoProvider: GetCipher()
    CryptoProvider->>KeyringRecord: cipher instance
    KeyringRecord->>Cipher: Encrypt(message, secret)
    Cipher->>Keyring: EncryptedMessage

    Keyring->>KeyringRecord: NewKey()
    KeyringRecord->>CryptoProvider: GetGenerator()
    CryptoProvider->>Generator: new Key
    Generator->>Keyring: Key

    Keyring->>KeyringRecord: SignMessage()
    KeyringRecord->>CryptoProvider: GetSigner()
    CryptoProvider->>KeyringRecord: signer instance
    KeyringRecord->>Signer: Sign(encryptedmessage, key)
    Signer->PrivKey: Read()
    PrivKey->Signer: Key bytes
    PrivKey->PrivKey: Wipe -- zeroing
    Signer->Signer: signMessage()
    Signer->Keyring: Signature

    Keyring->>KeyringRecord: VerifyMessage()
    KeyringRecord->>CryptoProvider: GetVerifier()
    CryptoProvider->>KeyringRecord: verifier instance
    KeyringRecord->>Verifier: Verify(encryptedmessage, signature, key)
    Verifier->PubKey: Read()
    PubKey->Verifier: Key bytes
    Verifier->Verifier: verifyMessage()
    Verifier->Keyring: true/false
```

## Alternatives

The alternatives may vary in the way of distributing the modules, putting some modules together as for example verify and signing in 
one place. This will affect the granularity of the code, thus the reusability and modularity. We aim to balance between simplicity and 
granularity.

## Decision

We will:

* Refactor module structure as the images attached.
* Define types and interfaces as the code attached.
* Refactor existing code into new structure and interfaces.
* Implement Unit Tests to ensure no backward compatibility issues.

## Consequences

### Backwards Compatibility

This refactor will involve changes on how the module is structured, providing cleaner interfaces and easier ways to use and extend. The impact should be minimal and not breaking any previous generated data.

The backward compatible sensitive elements are:

* Keys
* Signatures
* Encrypted data
* Hashes

### Positive

* Single place of truth.
* Easier to use interfaces.
* Easier to extend.
* Maintainability.
* Incentivize addition of implementations instead of forks.
* Decoupling.
* Sanitization of code.

### Negative

* It will involve an effort to adapt existing code.
* It will require attention to detail and audition.

### Neutral

* It will involve extensive testing.

## Test Cases

- The code will be unit tested to ensure a high code coverage
- There should be integration tests around Wallet, keyring and crypto providers.
- There should be benchmark tests for hashing, keyring, encryption, decryption, signing and verifying functions.

## Further Discussions

> While an ADR is in the DRAFT or PROPOSED stage, this section should contain a
> summary of issues to be solved in future iterations (usually referencing comments
> from a pull-request discussion).
>
> Later, this section can optionally list ideas or improvements the author or
> reviewers found during the analysis of this ADR.


## References

* {reference link}
