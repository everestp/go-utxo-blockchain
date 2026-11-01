package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	privKeyLen = 64
	pubKeylen  = 32
	seedlen    = 32
	addressLen = 20
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedlen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}

}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{

		value: ed25519.Sign(p.key, msg),
	}
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, pubKeylen)
	copy(b, p.key[:32])
	return &PublicKey{
		key: b,
	}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PublicKey) AddressBytes() Address {
	return Address{
		value: p.key[len(p.key)-addressLen:],
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

type Signature struct {
	value []byte
}

func (s *Signature) Bytes() []byte {
	return s.value
}

func (s *Signature) Verify(pubkey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubkey.key, msg, s.value)
}

type Address struct {
	value []byte
}

func (a *Address) String() string {
	return hex.EncodeToString(a.value)
}
