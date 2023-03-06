package types

import (
	"testing"

	"github.com/alesr/blocker/crypto"
	"github.com/alesr/blocker/proto"
	"github.com/alesr/blocker/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	fromPK, err := crypto.NewPrivateKey()
	require.NoError(t, err)

	fromAddr := fromPK.PublicKey().Bytes()

	toPK, err := crypto.NewPrivateKey()
	require.NoError(t, err)

	toAddr := toPK.PublicKey().Bytes()

	hash, err := util.RandomHash()
	require.NoError(t, err)

	input := &proto.TxInput{
		PreviousOutHash:  hash,
		PreviousOutIndex: 0,
		PublicKey:        fromAddr,
	}

	output1 := &proto.TxOutput{
		Amount:  uint64(5),
		Address: toAddr,
	}

	output2 := &proto.TxOutput{
		Amount:  uint64(95),
		Address: fromAddr,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig, err := SignTransaction(fromPK, tx)
	require.NoError(t, err)

	input.Signature = sig.Bytes()

	valid, err := VerifyTransaction(tx)
	require.NoError(t, err)

	assert.True(t, valid)
}
