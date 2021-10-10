package main

import (
	"fmt"
	"math"
)

func main() {
	num := 0.625
	decimalDegit := 3
	shift := math.Pow10(decimalDegit)

	remain := num * shift
	// var res []bool
	for {
		double := remain * 2

		// 切り捨て
		val := math.Floor(double / shift)

		fmt.Print(val)
		remain = double - val*shift

		// interger, frac := math.Modf(double)

		// fmt.Println(interger)
		// remain = math.Round(frac*100) / 100

		if remain == 0 {
			fmt.Println("break")
			break
		}
	}
}

func printBinary(bs []bool) {
	for _, b := range bs {
		if b {
			fmt.Print(1)
			continue
		}

		fmt.Print(0)
	}
}
