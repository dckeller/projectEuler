package main 

import (
	"fmt"
	"math"
	"strconv"
)

func findPower(power float64) int {
	pow := math.Pow(2, power)
	fmt.Println(pow)
	return int(pow)
}

func main() {
	num := findPower(float64(1000))
	stringNum := strconv.Itoa(num)
	sum := 0
	digits := make([]int, 0)
	
	for _, runeValue := range stringNum {
		i, _ := strconv.ParseInt(string(runeValue), 10, 0)
		digits = append(digits, int(i))	
	}

	for _, nums := range digits {
		sum += nums
	}
	// fmt.Println(sum)
}