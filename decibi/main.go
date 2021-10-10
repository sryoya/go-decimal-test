package main

import (
	"fmt"
	"math"
)

func main() {
	printFloatAsBiStr(0.6, 1)
	// num := 0.625      // 変換元の10進数
	// decimalDigit := 3 // 10進数での小数の桁数

	// // 整数にして計算する
	// shift := math.Pow10(decimalDigit)
	// remain := num * shift
	// fmt.Print("0.")
	// for {
	// 	double := remain * 2

	// 	// 1桁目を取得する (その2進数での桁の値)
	// 	val := math.Floor(double / shift)
	// 	fmt.Print(val)

	// 	// 1桁目を取り除く
	// 	remain = double - val*shift

	// 	// 残りが0になったら終わり
	// 	if remain == 0 {
	// 		break
	// 	}
	// }
}

func printFloatAsBiStr(v float64, digit int) {
	// 整数にして計算する
	shift := math.Pow10(digit)
	remain := v * shift
	fmt.Print("0.")
	count := 0
	for {
		// 1. 2倍する
		double := remain * 2
		// 2. 1桁目を取得する (その2進数での桁の値)
		val := math.Floor(double / shift)
		fmt.Print(val)
		// 3. 1桁目を取り除く
		remain = double - val*shift
		// 残りが0になるまで1~3を繰り返す
		if remain == 0 {
			break
		}
		count++
		if count > 100 {
			break
		}
	}
}
