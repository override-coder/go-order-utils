package utils

import (
	"crypto/rand"
	"math"
	"math/big"
)

func GenerateRandomSalt() int64 {
	maxInt := math.Pow(2, 32)
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(maxInt)))
	return nBig.Int64()
}

func GenerateRandomSaltBigInt() *big.Int {
	maxBigInt := new(big.Int).Lsh(big.NewInt(1), 256)
	maxBigInt.Sub(maxBigInt, big.NewInt(1))

	nBig, err := rand.Int(rand.Reader, maxBigInt)
	if err != nil {
		panic(err)
	}

	return nBig.Add(nBig, big.NewInt(1))
}
