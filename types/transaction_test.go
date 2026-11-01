package types

import (
	"testing"

	"github.com/everestp/blocker/crypto"
	"github.com/everestp/blocker/proto"
	"github.com/everestp/blocker/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Bytes()

	toPrivkey := crypto.GeneratePrivateKey()
	toAddress := toPrivkey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   utils.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}
	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}
	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress, // change back to sender for change
	}
	tx := &proto.Transaction{
		Version: 1,
		Inputes: []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()

	// 1. Test successful transaction verification
	assert.True(t, VerifyTransaction(tx))

	// 2. Test failure when transaction outputs are tampered with
	tx.Outputs[0].Amount = 10 // changed amount from 5 to 10
	assert.False(t, VerifyTransaction(tx))
}
