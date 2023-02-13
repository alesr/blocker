package crypto

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPrivateKey(t *testing.T) {
	t.Parallel()

	privKey, err := NewPrivateKey()
	require.NoError(t, err)

	assert.Equal(t, ed25519.PrivateKeySize, len(privKey.Bytes()))
	assert.Equal(t, ed25519.PublicKeySize, len(privKey.PublicKey().Bytes()))
}

func TestNewPrivateKeyFromString(t *testing.T) {
	t.Parallel()

	seed := "28ca1d010779e7c9d6a8f61201d987369f21357daa4eac1d807fb8a7a079d00f"
	privKey, err := NewPrivateKeyFromString(seed)
	require.NoError(t, err)

	assert.Equal(t, ed25519.PrivateKeySize, len(privKey.Bytes()))
	assert.Equal(t, ed25519.PublicKeySize, len(privKey.PublicKey().Bytes()))
}

func TestPrivateKey_Sign(t *testing.T) {
	t.Parallel()

	privKey, err := NewPrivateKey()
	require.NoError(t, err)

	msg := []byte("xablablau")
	sig := privKey.Sign(msg)

	t.Run("verify signature", func(t *testing.T) {
		assert.True(t, sig.Verify(msg, privKey.PublicKey()))
	})

	t.Run("verify signature with wrong message", func(t *testing.T) {
		assert.False(t, sig.Verify([]byte("something else"), privKey.PublicKey()))
	})

	t.Run("verify signature with wrong public key", func(t *testing.T) {
		privKey, err := NewPrivateKey()
		require.NoError(t, err)

		assert.False(t, sig.Verify(msg, privKey.PublicKey()))
	})

	t.Run("verify signature with wrong signature", func(t *testing.T) {
		privKey, err := NewPrivateKey()
		require.NoError(t, err)

		assert.False(t, sig.Verify(msg, privKey.PublicKey()))
	})
}

func TestPublicKey_Address(t *testing.T) {
	t.Parallel()
	givenSeed := "28ca1d010779e7c9d6a8f61201d987369f21357daa4eac1d807fb8a7a079d00f"
	expectedAddr := "c45d2ed56f5756fa9d1e2058b1a2d8503194bfc2"

	privKey, err := NewPrivateKeyFromString(givenSeed)
	require.NoError(t, err)

	addr := privKey.PublicKey().Address()

	assert.Equal(t, AddressSize, len(addr.Bytes()))
	assert.Equal(t, expectedAddr, addr.String())
}
