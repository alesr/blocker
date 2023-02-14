package types

import (
	"crypto/sha256"
	"fmt"

	"github.com/alesr/blocker/crypto"
	"github.com/alesr/blocker/proto"
	pb "github.com/golang/protobuf/proto"
)

func SignBlock(pk *crypto.PrivateKey, block *proto.Block) (*crypto.Signature, error) {
	hash, err := HashBlock(block)
	if err != nil {
		return nil, fmt.Errorf("could not hash block: %w", err)
	}
	return pk.Sign(hash), nil
}

// HashBlock creates a SHA256 hash of the header.
func HashBlock(block *proto.Block) ([]byte, error) {
	b, err := pb.Marshal(block)
	if err != nil {
		return nil, fmt.Errorf("could not marshal proto block: %w", err)
	}

	hash := sha256.Sum256(b)
	return hash[:], nil
}
