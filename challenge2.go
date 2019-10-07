package main

import (
	"fmt"
)

func FibSequence(nums int) []int {
	Seq := make([]int, 0)
	a := 0
	b := 1
	temp := 0
	for i := 0; i < nums; i++ {
	  if nums == 0 {
	  	Seq = append(Seq, 0)
	  } else if nums == 1 || nums == 2 {
	  	Seq = append(Seq, 1)
	  }	else {
		  temp = a
			a = b
			b = (temp + b)
			Seq = append(Seq, b)
			fmt.Println(Seq)
	  }
	}
	return Seq
}

func ReduceEvens(fib []int) int {
	sum := 0

	for _, v := range fib {
		if v % 2 == 0 && v < 4000000 {
			sum = sum + v
			fmt.Println(sum)
		}
	}
	return sum
}

func main() {
	fmt.Println(ReduceEvens(FibSequence(50)))
}