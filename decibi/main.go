package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

// FloatToBinaryStr converts a float64 to a binary string format.
func FloatToBinaryStr(f float64, originalDigit, binaryMaxDigit int) string {
	// get the integer part and the decimal part
	intf, f := math.Modf(f)
	i := int(intf)

	var b strings.Builder

	// convert the integer part to binary
	for i > 0 {
		b.WriteString(strconv.FormatInt(int64(i%2), 10))
		i = i / 2
	}

	// covert the decimal part to binary
	shift := math.Pow10(originalDigit)
	remain := f * shift
	count := 0
	for {
		double := remain * 2
		val := math.Floor(double / shift)
		b.WriteString(strconv.FormatInt(int64(val), 10))
		remain = double - val*shift
		if remain == 0 {
			break
		}
		count++
		if count > binaryMaxDigit {
			break
		}
	}
	return b.String()
}
