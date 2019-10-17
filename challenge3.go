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

func FindPrimes(highestPrime int) []int {
	allPrimes := make([]int, 0)

	for i := 2; i <= highestPrime; i++ {
		if IsPrime(i) == true {
			allPrimes = append(allPrimes, i)
		}
	}
	return allPrimes
}

func FindFactor(all []int, num int) int {
	Highest := make([]int, 0)

	for _, factors := range all {
		if num % factors == 0 {
			Highest = append(Highest, num)
		}
	}
	return Highest[len(Highest)-1]
}

func main() {
	fmt.Println(FindFactor(FindPrimes(1000), 1000))
}
