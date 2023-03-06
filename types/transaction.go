package types

import (
	"crypto/sha256"
	"fmt"

	"github.com/alesr/blocker/crypto"
	"github.com/alesr/blocker/proto"
	pb "github.com/golang/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) (*crypto.Signature, error) {
	hash, err := HashTransaction(tx)
	if err != nil {
		return nil, fmt.Errorf("could not hash transaction: %w", err)
	}
	return pk.Sign(hash), nil
}

func HashTransaction(tx *proto.Transaction) ([]byte, error) {
	b, err := pb.Marshal(tx)
	if err != nil {
		return nil, fmt.Errorf("could not marshal proto transaction: %w", err)
	}

	hash := sha256.Sum256(b)
	return hash[:], nil
}

func VerifyTransaction(tx *proto.Transaction) (bool, error) {
	for _, input := range tx.Inputs {
		sig, err := crypto.SignatueFromBytes(input.Signature)
		if err != nil {
			return false, fmt.Errorf("could not decode signature: %w", err)
		}

		pubKey, err := crypto.PublicKeyFromBytes(input.PublicKey)
		if err != nil {
			return false, fmt.Errorf("could not decode public key: %w", err)
		}

		// By the time we create the transaction we don't yet have it signed.
		// TODO: Find a better way to do this, as we should not modify the transaction.
		input.Signature = nil

		hashTx, err := HashTransaction(tx)
		if err != nil {
			return false, fmt.Errorf("could not hash transaction: %w", err)
		}

		if !sig.Verify(hashTx, pubKey) {
			return false, nil
		}
	}
	return true, nil
}
