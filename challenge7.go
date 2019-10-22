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

	for i := 2; ; i++ {
		if IsPrime(i) == true {
			allPrimes = append(allPrimes, i)
		} else if len(allPrimes) == highestPrime {
			break
		}
	}
	highest := allPrimes[len(allPrimes)-1]
	fmt.Println("The highest prime is %d", highest)
	return allPrimes
}

func main() {
	FindPrimes(10001)
}