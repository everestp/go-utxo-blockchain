package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privKeyLen)
	pubkey := privKey.Public()
	assert.Equal(t, len(pubkey.Bytes()), pubKeylen)
}

func TestNewPrivateKeyFromSeed(t *testing.T) {
	seed := make([]byte, seedlen)
	_, err := io.ReadFull(rand.Reader, seed)
	assert.NoError(t, err)

	privKey1 := NewPrivateKeyFromSeed(seed)
	assert.Equal(t, len(privKey1.Bytes()), privKeyLen)

	// Creating a private key from the same seed should yield identical keys
	privKey2 := NewPrivateKeyFromSeed(seed)
	assert.Equal(t, privKey1.Bytes(), privKey2.Bytes())
	assert.Equal(t, privKey1.Public().Bytes(), privKey2.Public().Bytes())
}
func TestNewPrivateKeyFromString(t *testing.T) {
	// Generate a random seed
	seed := make([]byte, seedlen)
	_, err := io.ReadFull(rand.Reader, seed)
	assert.NoError(t, err)

	// Encode the seed to a hex string
	hexStr := hex.EncodeToString(seed)

	// Create private key from string
	privKey := NewPrivateKeyFromString(hexStr)
	assert.NotNil(t, privKey)
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	// Verify it matches a private key created directly from the seed
	expectedPrivKey := NewPrivateKeyFromSeed(seed)
	assert.Equal(t, expectedPrivKey.Bytes(), privKey.Bytes())
}
func TestPrivateKeySign(t *testing.T) {
	privkey := GeneratePrivateKey()
	pubkey := privkey.Public()
	msg := []byte("foo bar foo")
	sig := privkey.Sign(msg)
	assert.True(t, sig.Verify(pubkey, msg))

	//Test with invalid message
	assert.False(t, sig.Verify(pubkey, []byte("foo")))

	//Test with invalid pubkey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubkey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubkey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubkey := privKey.Public()
	address := pubkey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
