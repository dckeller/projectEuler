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

func (n *Numbers)CountOnes(ones int) int {
	for i, v := range n.All {
		if ones == i {
			n.Count += v
		}
	}	
	return n.Count
}

func (n *Numbers)CountTens(tens int) int {
	for i, v := range n.Tens {
		if tens == i { 
			n.Count += v
		}
	}	
	return n.Count
}

func (n *Numbers)CountTenths(tenths int) int {
	for i, v := range n.Tenths {
		if tenths == i {
			n.Count += v
		}
	}
	return n.Count
}

func main() {
	nums := Numbers{
		// 0 - 9 //
		All: 		[]int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4},

		// 10 - 19 //
		Tens: 	[]int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8},

		// 20, 30, 40, 50, 60, 70, 80, 90 //
		Tenths:	[]int{0, 0, 6, 6, 5, 5, 5, 7, 6, 6},

		Count: 	0,
	}

	for j := 1; j <= 1000; j++ {
		CurrentNum := IntToStringToInt(j)

		// Handle 1 - 9 //
		if len(CurrentNum) == 1 {
			nums.CountOnes(CurrentNum[0])
			fmt.Println(nums.Count)
		}

		// Handle 10 - 19 //
		if len(CurrentNum) == 2 && CurrentNum[0] == 1 {
			nums.CountTens(CurrentNum[1])
			fmt.Println(nums.Count)
		}

		// Handle Tenths //
		if len(CurrentNum) == 2 && CurrentNum[0] > 1 && CurrentNum[1] == 0 {
			nums.CountTenths(CurrentNum[0])
			fmt.Println(nums.Count)
		}

		// Handle 21 - 99 //
		if len(CurrentNum) == 2 && CurrentNum[0] >= 2 && CurrentNum[1] >= 1 {
			nums.CountTenths(CurrentNum[0])
			nums.CountOnes(CurrentNum[1])
			fmt.Println(nums.Count)
		}

		// Handle 100 // 
		if len(CurrentNum) == 3 && CurrentNum[1] == 0 && CurrentNum[2] == 0 {
			nums.CountOnes(CurrentNum[0])
			nums.Count += 7
			fmt.Println(nums.Count)
		}

		// Handle 101 - 109 //
		if len(CurrentNum) == 3 && CurrentNum[1] == 0 && CurrentNum[2] >= 1 {
			nums.CountOnes(CurrentNum[0])
			nums.Count += 10
			nums.CountOnes(CurrentNum[2])
			fmt.Println(nums.Count)
		}

		// Handle 110 - 119 // 
		if len(CurrentNum) == 3 && CurrentNum[1] == 1 {
			nums.CountOnes(CurrentNum[0])
			nums.Count += 10 
			nums.CountTens(CurrentNum[2])
			fmt.Println(nums.Count)
		}

		// Handle 120ths //
		if len(CurrentNum) == 3 && CurrentNum[2] == 0 && CurrentNum[1] >= 2 && CurrentNum[0] >= 1 {
			nums.CountOnes(CurrentNum[0])
			nums.Count += 10 
			nums.CountTenths(CurrentNum[1])
			fmt.Println(nums.Count)
		}

		// Handle 121 - 999 //
		if len(CurrentNum) == 3 && CurrentNum[2] >= 1 && CurrentNum[1] >= 2 && CurrentNum[0] >= 1 {
			nums.CountOnes(CurrentNum[0])
			nums.Count += 10
			nums.CountTenths(CurrentNum[1])
			nums.CountOnes(CurrentNum[2])
			fmt.Println(nums.Count)
		}

		// Handle 1000 //
		if len(CurrentNum) == 4 {
			nums.CountOnes(CurrentNum[0])
			nums.Count += 8
			fmt.Println(nums.Count)
		}

		fmt.Println(CurrentNum)
		CurrentNum = nil
	}
}			