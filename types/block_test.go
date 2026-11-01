package types

import (
	"bytes"
	"testing"

	"github.com/everestp/blocker/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashBlock(t *testing.T) {
	block := utils.RandomBlock()

	hash1 := HashBlock(block)
	hash2 := HashBlock(block)

	require.Len(t, hash1, 32)
	require.Equal(t, hash1, hash2)

	block.Header.Height++

	hash3 := HashBlock(block)
	require.False(t, bytes.Equal(hash1, hash3))
}

func TestSignBlock(t *testing.T) {

	block := utils.RandomBlock()
	hash := HashBlock(block)

	assert.Equal(t, 32, len(hash))
}
