package main

import (
	"fmt"
	"math/big"
)

func main() {
	// a := new(big.Rat).SetInt64(1)
	// b := big.NewRat(1, 2)
	// diff := a.Sub(a, b) // aも更新される
	// fmt.Printf("%s - %s = %v\n", a.FloatString(3), b.FloatString(3), diff.FloatString(3))

	total := new(big.Rat).SetInt64(1)
	a := big.NewRat(1, 3)
	b := big.NewRat(1, 3)
	c := big.NewRat(1, 3)
	remain := total.Sub(total, a).Sub(total, b).Sub(total, c) // 1 - 1/3 - 1/3 - 1/3 = 0
	fmt.Printf("1 - %s - %s -%s = %s", a.FloatString(3), b.FloatString(3), c.FloatString(3), remain.FloatString(3))
}
