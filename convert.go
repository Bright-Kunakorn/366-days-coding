package main

import (
	"math/big"
)

func multiply(num1 string, num2 string) string {
	bigNum1, _ := new(big.Int).SetString(num1, 10)
	bigNum2, _ := new(big.Int).SetString(num2, 10)
	result := new(big.Int).Mul(bigNum1, bigNum2)

	return result.String()
}
