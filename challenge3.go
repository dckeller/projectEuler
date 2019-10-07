package main 

import (
	"fmt"
	// "math"
)

func IsPrime(value int) []int {
	Prime := make([]int, 0)

  for i := 2; i < value; i++ {
      if value%i == 0 {
          Prime = append(Prime, 0)
      } else {
      	Prime = append(Prime, i)
      	fmt.Println(Prime)
      }
  }
  return Prime
}

func main() {
	IsPrime(100)
}