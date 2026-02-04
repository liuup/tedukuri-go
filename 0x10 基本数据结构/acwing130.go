package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	binomial := new(big.Int).Binomial(int64(2*n), int64(n))
	denom := big.NewInt(int64(n + 1))
	catalan := new(big.Int).Div(binomial, denom)
	fmt.Println(catalan.String())
}

