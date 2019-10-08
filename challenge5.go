package main

import (
	"fmt"
)

func main() {

	for i := 1; ; i++{
		if i % 1 == 0 && i % 2 == 0 && i % 3 == 0 && i % 5 == 0 && i % 7 == 0 && i % 11 == 0 && i % 13 == 0 && i % 16 == 0 && i % 17 == 0 && i % 18 == 0 && i % 19 == 0 && i % 20 == 0 {
			fmt.Println(i)
			break
		} 
	}
}