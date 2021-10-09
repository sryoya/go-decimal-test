package main

import (
	"fmt"
	"math"
)

func main() {
	num := 0.625

	remain := num
	var res []bool
	for {
		double := remain * 2
		res = append(res, double >= 1)

		_, frac := math.Modf(double)
		remain = frac

		fmt.Println(math.Floor(double))
		fmt.Println(double)
		if math.Floor(double) == double {
			break
		}
		if len(res) > 5 {
			break
		}
	}

	printBinary(res)
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
