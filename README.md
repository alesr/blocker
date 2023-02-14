# blocker
[![.github/workflows/ci.yaml](https://github.com/alesr/blocker/actions/workflows/ci.yaml/badge.svg)](https://github.com/alesr/blocker/actions/workflows/ci.yaml)

This project is an ongoing exercise to learn about blockchain and cryptocurrencies, and the work is in progress.

## Crypto package

The crypto package contains code for generating, signing, and verifying cryptographic keys, as well as signing and verifying signatures. Here are some important details about the package:

- It is written in Go and includes functions for generating public/private key pairs using the Ed25519 algorithm.
- It provides functions for signing messages and verifying signatures using the Ed25519 algorithm.
- It also includes functions for generating addresses from public keys.
- The PrivateKey struct represents a private key, while the PublicKey struct represents a public key.
- The Signature struct represents a signature, and the Address struct represents an address derived from a public key.

## Types package

The types package contains code for defining the data types used in a blockchain, as well as some utility functions. Here are some important details about the package:

- It is written in Go and includes functions for hashing blocks and signing them.
- It defines the Block and Header messages using the Protocol Buffers format, which is used to serialize and deserialize data - structures.
- The Block message contains a header and a list of transactions, while the Header message contains the metadata for a block.
- The SignBlock function is used to sign a block with a private key, while the HashBlock function is used to hash a block's header.

## Util package

The util package contains some utility functions for generating random data. Here are some important details about the package:

- It is written in Go and includes functions for generating random blocks and hashes.
- The RandomBlock function generates a random block with a random hash.
- The RandomHash function generates a random hash of 32 bytes.

## Protocol Buffers file

The `block.proto` file defines the messages used in the blockchain. Here are some important details about the file:

- It is written in the Protocol Buffers format, which is used to serialize and deserialize data structures.
- It defines the Block and Header messages, which are used to represent a block and its metadata.
- The Transaction message is empty, as it is not used in this implementation.

