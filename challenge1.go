package main 

import (
	"fmt"
)

func multiples(input int) []int {
	var Nums = make([]int, 0)

	for i := 0; i < input; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			Nums = append(Nums, i)
			fmt.Println(Nums)
		}
	}
	return Nums
}

func reduce(nums []int) int {
	sum := 0

	for _, v := range nums {
		sum = sum + v
	}
	return sum
}

func main() {
	fmt.Println(reduce(multiples(1000)))
}