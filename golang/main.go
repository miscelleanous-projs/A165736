package main

import (
	"fmt"
	"math/big"
)

// OEIS-A165736
// https://oeis.org/A165736

func a165736(n *big.Int) *big.Int {
	if new(big.Int).Mod(n, big.NewInt(10)).Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	result := new(big.Int).Set(n)

	for i := 1; i <= 10; i++ {
		exp := big.NewInt(int64(i))
		base := new(big.Int).Set(n)
		result.Exp(base, result, new(big.Int).Exp(big.NewInt(10), exp, nil))
	}

	return result
}

func main() {
	fmt.Print("[")
	for i := big.NewInt(1); i.Cmp(big.NewInt(24+1)) < 0; i.Add(i, big.NewInt(1)) {
		if i.Cmp(big.NewInt(1)) == 1 {
			fmt.Print(", ")
		}
		fmt.Print(a165736(i))
	}
	fmt.Print("]")
}
