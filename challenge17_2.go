package main 

import (
	"fmt"
	"strings"
	"strconv"
)

type Numbers struct {
	All  				[]int
	Tens				[]int
	Tenths			[]int
	Count 			int
}

func IntToStringToInt(num int) []int {
	CurrentNum := make([]int, 0)

	stringNum := strconv.Itoa(num)
	s := strings.Split(stringNum, "")

	for _, v := range s {
		j, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		CurrentNum = append(CurrentNum, j)
	}
	return CurrentNum  	
}

func (n *Numbers)CountOnes(ones []int) int {
	for i, v := range n.All {
		if ones[0] == i {
			n.Count += v
		}
	}
	return n.Count
}

func (n *Numbers)CountTens(tens []int) int {
	for i, v := range n.Tenths {
		if tens[0] == i {
			n.Count += v
		}
	}

	for i, v := range n.Tens {
		if tens[1] == i && tens[1] != 0 {
			n.Count += v
		}
	}

	return n.Count
}

func main() {
	nums := Numbers{
		// 0 - 9 //
		All: 		[]int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4},

		// 11 - 19 //
		Tens: 	[]int{0, 6, 6, 8, 8, 7, 7, 9, 8, 8},

		// 20, 30, 40, 50, 60, 70, 80, 90 //
		Tenths:	[]int{0, 3, 6, 5, 5, 5, 5, 7, 6, 6},
		Count: 	0,
	}

	for j := 1; j <= 20; j++ {
		CurrentNum := IntToStringToInt(j)

		if len(CurrentNum) == 1 {
			nums.CountOnes(CurrentNum)
			fmt.Println(nums.Count)
		} 

		if len(CurrentNum) == 2 {
			nums.CountTens(CurrentNum)
			fmt.Println(nums.Count)
		}
		fmt.Println(CurrentNum)
		CurrentNum = nil
	}
}