package main

import (
	"fmt"
)

func IsPrime(value int) bool {

	for i := 2; i < value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func FindPrimes(quantity int) []int {
	allPrimes := make([]int, 0)

	for i := 2; i <= quantity; i++ {
		if IsPrime(i) == true {
			allPrimes = append(allPrimes, i)
		}
	}
	return allPrimes
}

// func BatchAllPrimes(all []int) [][]int {
// 	BatchSize := 20
// 	var Batches [][]int

// 	for BatchSize < len(all) {
// 		all, Batches = all[BatchSize:], append(Batches, all[0:BatchSize:BatchSize])
// 	}
// 	Batches = append(Batches, all)
// 	fmt.Println(Batches)
// 	return Batches
// }

func ReducePrimes(allPrimes []int) int {
	sum := 0

	for _, v := range allPrimes {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(ReducePrimes(FindPrimes(2000000)))
	// BatchAllPrimes(FindPrimes(2000000))
}
