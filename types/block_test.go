package types

import (
	"testing"

	"github.com/alesr/blocker/crypto"
	"github.com/alesr/blocker/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignBlock(t *testing.T) {
	block, err := util.RandomBlock()
	require.NoError(t, err)

	pk, err := crypto.NewPrivateKey()
	require.NoError(t, err)

	sig, err := SignBlock(pk, block)
	require.NoError(t, err)

	assert.Equal(t, 64, len(sig.Bytes()))

	hash, err := HashBlock(block)
	require.NoError(t, err)

	assert.True(t, sig.Verify(hash, pk.PublicKey()))
}

func TestHashBlock(t *testing.T) {
	t.Parallel()

	block, err := util.RandomBlock()
	require.NoError(t, err)

	hash, err := HashBlock(block)
	require.NoError(t, err)

	assert.Equal(t, 32, len(hash))
}
