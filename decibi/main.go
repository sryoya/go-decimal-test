package main

import (
	"fmt"
	"math"
)

func main() {
	num := 0.625

	remain := num
	for {
		double := remain * 2

		interger, frac := math.Modf(double)

		fmt.Print(interger)
		remain = math.Round(frac*100) / 100

		if remain == 0 {
			fmt.Println("break")
			break
		}
	}
}
