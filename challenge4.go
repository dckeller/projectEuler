package main

import (
	"fmt"
)

// func CheckLength(num int) bool {
// 	count := 0
// 	for num != 0 {
// 		num /= 10
// 		count += 1
// 	}
// 	if count%2 == 0 {
// 		return true
// 	}
// 	return false
// }

func CheckPalandrome(all int) bool {
	Check := make([]int, 6)

	Check[0] = (all / 100000) % 10
	Check[1] = (all / 10000) % 10
	Check[2] = (all / 1000) % 10
	Check[3] = (all / 100) % 10
	Check[4] = (all / 10) % 10
	Check[5] = all % 10

	if Check[0] == Check[5] && Check[1] == Check[4] && Check[2] == Check[3] {
		return true
	}
	return false
}

func makeNumbers() int {
	Base := make([]int, 0)
	currentMax := 0

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			fmt.Println(i)
			fmt.Println(j)
			fmt.Println("--------")
			product := i * j
			if currentMax > product {
				break
			}
			// if CheckLength(product) == true && CheckPalandrome(product) == true {
			if CheckPalandrome(product) == true {
				currentMax = product
				Base = append(Base, currentMax)
			}
		}
	}
	return Base[len(Base)-1]
}

func main() {
	fmt.Println(makeNumbers())
}
