package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	a, _ := decimal.NewFromString("0.1")
	b := decimal.NewFromInt(3)
	c, _ := decimal.NewFromString("0.3")

	fmt.Printf("a * b =%v\n", a.Mul(b))
	fmt.Printf("c:% v\n", c)
	fmt.Printf("a * b == c: %t\n", a.Mul(b).Equal(c))
}
