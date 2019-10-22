package main 

import (
	"fmt"
)

func Squares(oneHundred int) int {
	Nums := make([]int, oneHundred)
	AllNums := make([]int, oneHundred)
	SumSquares := 0
	SquaredSum := 0

	for i := range Nums {
		Nums[i] = 1 + i
		SquaredSum += Nums[i]
		fmt.Println(SquaredSum)
	}
	for i, v := range Nums {
		AllNums[i] = v * v
		SumSquares += AllNums[i]
		fmt.Println(SumSquares)
	}
	return (SquaredSum * SquaredSum) - SumSquares
}

func main() {
	fmt.Println(Squares(100))
}