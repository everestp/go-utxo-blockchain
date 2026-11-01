package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privKeyLen)
	pubkey := privKey.Public()
	assert.Equal(t, len(pubkey.Bytes()), pubKeylen)
}

func TestPrivateKeySign(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubkey := privkey.Public()
	msg := []byte("foo bar foo")
	sig := privkey.Sign(msg)
	assert.True(t, sig.Verify(pubkey, msg))

	//Test with  invalid message
	assert.False(t, sig.Verify(pubkey, []byte("foo")))

	//Test with invalid pubkey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubkey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubkey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubkey := privKey.Public()

}
