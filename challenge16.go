package main 

import (
	"fmt"
	"math/big"
	"strconv"
)

// func findPower(power float64) int {
// 	pow := math.Pow(2, power)
// 	fmt.Println(pow)
// 	return int(pow)
// }

// func main() {
// 	num := findPower(float64(1000))
// 	stringNum := strconv.Itoa(num)
// 	sum := 0
// 	digits := make([]int, 0)

// 	for _, runeValue := range stringNum {
// 		i, _ := strconv.ParseInt(string(runeValue), 10, 0)
// 		digits = append(digits, int(i))	
// 	}

// 	for _, nums := range digits {
// 		sum += nums
// 	}
// }


// big/math package func(x *Int) Exp //
func powBig(num, power int) *big.Int {
	bigNum := big.NewInt(int64(num))

	// The number to be added to //
	res := big.NewInt(1)



	for power > 0 {
		fmt.Println(res)
		fmt.Println(bigNum)
		fmt.Println(power)
		fmt.Println("------------")

		// The number to be changed and then equal to res //
		temp := new(big.Int)

		// If the power is an odd number, multiply 1 and num //
		if power % 2 == 1 {
			temp.Mul(res, bigNum)
			res = temp
		}

		// If the power is even multiply num and num // 
		temp = new(big.Int)
		temp.Mul(bigNum, bigNum)

		// If even or odd bigNum equals multiplied temp //
		bigNum = temp

		// Continue loop power reaches 0 //
		power /= 2
	}

	return res
}

func HandleBigInt(big *big.Int) []int64 {
	stringNum := big.String()
	digits := make([]int64, 0)

	// Runes are the unicode (number) for a string //
	for _, runeValue := range stringNum {

		// (Pass the string (convert rune to string), base, integer type) //
		i, _ := strconv.ParseInt(string(runeValue), 10, 0)

		digits = append(digits, i)
	}
	return digits
}

func ReducePower(bigInt []int64) int64 {
	sum := int64(0)

	for _, v := range bigInt {
		sum += v
	}
	return sum
}


func main() {
	// fmt.Println(powBig(2, 1000))
	fmt.Println(ReducePower(HandleBigInt(powBig(2, 1000))))
}
