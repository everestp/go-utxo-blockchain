package types

import (
	"crypto/sha256"

	"github.com/everestp/blocker/crypto"
	"github.com/everestp/blocker/proto"
	pb "github.com/golang/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	// Clone the transaction and clear all input signatures so hashing
	// is consistent and doesn't mutate the original transaction.
	txCopy := pb.Clone(tx).(*proto.Transaction)
	for _, input := range txCopy.Inputes {
		input.Signature = nil
	}

	b, err := pb.Marshal(txCopy)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {
	// Compute the hash of the transaction without any signatures
	txHash := HashTransaction(tx)

	for _, input := range tx.Inputes {
		if len(input.Signature) == 0 {
			return false
		}

		sig := crypto.SignatureFromBytes(input.Signature)
		pubKey := crypto.PublicKeyFromBytes(input.PublicKey)

		if !sig.Verify(pubKey, txHash) {
			return false
		}
	}
	return true
}
