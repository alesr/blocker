package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	// AddressSize is the size of an address in bytes.
	AddressSize int = 20
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKey() (*PrivateKey, error) {
	seed := make([]byte, ed25519.SeedSize)

	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		return nil, fmt.Errorf("could not instantiate new private key: %w", err)
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}, nil
}

func NewPrivateKeyFromSeed(seed []byte) (*PrivateKey, error) {
	if len(seed) != ed25519.SeedSize {
		return nil, fmt.Errorf("invalid seed size: %d", len(seed))
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}, nil
}

func NewPrivateKeyFromString(seed string) (*PrivateKey, error) {
	hexStr, err := hex.DecodeString(seed)
	if err != nil {
		return nil, fmt.Errorf("could not decode hex string: %w", err)
	}
	return NewPrivateKeyFromSeed(hexStr)
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{data: ed25519.Sign(p.key, msg)}
}

func (p *PrivateKey) PublicKey() *PublicKey {
	b := make([]byte, ed25519.PublicKeySize)
	copy(b, p.key[ed25519.PublicKeySize:])
	return &PublicKey{key: b}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PublicKey) Address() *Address {
	return &Address{value: p.key[len(p.key)-AddressSize:]}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func (p *PublicKey) Verify(msg, sig []byte) bool {
	return ed25519.Verify(p.key, msg, sig)
}

type Signature struct {
	data []byte
}

func (s *Signature) Bytes() []byte {
	return s.data
}

func (s *Signature) Verify(msg []byte, pub *PublicKey) bool {
	return pub.Verify(msg, s.data)
}

type Address struct {
	value []byte
}

func (a *Address) Bytes() []byte {
	return a.value
}

func (a *Address) String() string {
	return fmt.Sprintf("%x", a.value)
}
