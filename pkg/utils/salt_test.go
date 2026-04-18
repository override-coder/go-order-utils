package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomSalt(t *testing.T) {
	r1 := GenerateRandomSalt()
	r2 := GenerateRandomSalt()
	r3 := GenerateRandomSalt()

	assert.Positive(t, r1)
	assert.Positive(t, r2)
	assert.Positive(t, r3)

	assert.NotEqual(t, r1, r2)
	assert.NotEqual(t, r2, r3)
}

func TestGenerateRandomSaltBigInt(t *testing.T) {
	r1 := GenerateRandomSaltBigInt()
	r2 := GenerateRandomSaltBigInt()
	r3 := GenerateRandomSaltBigInt()

	assert.Positive(t, r1.Sign())
	assert.Positive(t, r2.Sign())
	assert.Positive(t, r3.Sign())

	assert.LessOrEqual(t, r1.BitLen(), 256)
	assert.LessOrEqual(t, r2.BitLen(), 256)
	assert.LessOrEqual(t, r3.BitLen(), 256)

	assert.NotEqual(t, 0, r1.Cmp(r2))
	assert.NotEqual(t, 0, r2.Cmp(r3))
}
