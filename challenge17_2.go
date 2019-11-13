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
		if len(ones) == 1 && ones[0] == i {
			n.Count += v
		}

		if len(ones) == 2 && ones[1] == i {
			n.Count += v
		}

		if len(ones) == 3 && ones[2] == i {
			n.Count += v
		}
	}
	return n.Count
}

func (n *Numbers)CountTens(tens []int) int {
	for i, v := range n.Tens {
		if len(tens) == 2 && tens[1] == i { 
			n.Count += v
		}

		if len(tens) == 3 && tens[2] == i {
			n.Count += v
		}
	}
	return n.Count
}

func (n *Numbers)CountTenths(tenths []int) int {
	for i, v := range n.Tenths {
		if len(tenths) == 2 && tenths[0] == i {
			n.Count += v
		}

		if len(tenths) == 2 && tenths[1] != 0 {
			n.Count += v
		}

		if len(tenths) == 3 && tenths[1] == i {
			n.Count += v
		}
	}
	return n.Count
}

func (n *Numbers)CountHundreds(hundreds []int) int {
	// Handle first digit //
	for i, v := range n.All {
		if hundreds[0] == i {
			n.Count += v
		}
	}

	// Handle "hundred and" //
	n.Count += 10
	return n.Count
}
 
func main() {
	nums := Numbers{
		// 0 - 9 //
		All: 		[]int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4},

		// 11 - 19 //
		Tens: 	[]int{0, 6, 6, 8, 8, 7, 7, 9, 8, 8},

		// 10, 20, 30, 40, 50, 60, 70, 80, 90 //
		Tenths:	[]int{0, 3, 6, 5, 5, 5, 5, 7, 6, 6},

		Count: 	0,
	}

	for j := 1; j <= 21; j++ {
		CurrentNum := IntToStringToInt(j)

		if len(CurrentNum) == 1 {
			nums.CountOnes(CurrentNum)
			fmt.Println(nums.Count)
		} 

		if len(CurrentNum) == 2 && CurrentNum[1] == 0 {
			nums.CountTenths(CurrentNum)
			fmt.Println(nums.Count)
		}

		if len(CurrentNum) == 2 {
			nums.CountTens(CurrentNum)
			fmt.Println(nums.Count)
		}

		// Handle one hundred //
		if len(CurrentNum) == 3 && CurrentNum[1] == 0 && CurrentNum [2] == 0 {
			nums.Count += 10
			fmt.Println(nums.Count) 
		}

		// Handle one hundred and tenths //
		if len(CurrentNum) == 3 && CurrentNum[2] == 0 {
			nums.CountHundreds(CurrentNum)
			nums.CountTenths(CurrentNum)
			fmt.Println(nums.Count)
		}

		if len(CurrentNum) == 3 {
			nums.CountHundreds(CurrentNum)
			fmt.Println(nums.Count) 
		}

		fmt.Println(CurrentNum)
		CurrentNum = nil
	}
}