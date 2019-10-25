package main 

import (
	"fmt"
	"math"
)

func PythagoreanFinder(a, b, c float64) bool {
	if (a * a) + (b * b) == (c * c) {
		return true
	}
	return false  
}

func FindTriplet(a, b, c, quantity float64) bool {
	if a + b + c == quantity {
		return true
	}
	return false
}

func PythagMaker(stop float64) float64 {
	var a, b, c float64
	var Answer float64 

			b = math.Sqrt(stop / 2)
			a = math.Sqrt(b / 2)
			c = math.Sqrt(stop - (b + a))

			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println("--------------------")
			
			if PythagoreanFinder(a, b, c) == true {
				fmt.Println("----")
			}

			if FindTriplet(a, b, c, stop) == true  {
				Answer = (a + b + c)
			}
	return Answer
}

func main() {
	// fmt.Println(PythagoreanFinder(3, 4, 5))
	// fmt.Println(FindTriplet(3, 4, 5, 12))
	fmt.Println(PythagMaker(float64(1000)))
}