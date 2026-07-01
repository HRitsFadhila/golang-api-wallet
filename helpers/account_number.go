package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateAccountNumber() string{
	prefix := "26"
	max := big.NewInt(100000000)
	randomNumber, _ := rand.Int(rand.Reader, max)
	formattedRandom := fmt.Sprintf("%08d", randomNumber.Int64())

	return prefix + formattedRandom
}