package main

import (
	"fmt"
)

func sequenceMaker(start int64) []int64 {
	sequence := make([]int64, 0)
	starter := start

	for i := int64(0); ; i++ {
		if start == 1 {
			sequence = append(sequence, starter)
			break
		} else {
			if start % 2 == 0 {
				start = start / 2
				sequence = append(sequence, start)
			} else {
				start = (start * 3) + 1
				sequence = append(sequence, start)
			}
		}
	}
	return sequence
}

func main() {
	finalSequence := make([]int64, 0)

	for i := int64(999999); i > 1; i-- {
		tempLength := sequenceMaker(i)

		if len(tempLength) > len(finalSequence) {
			finalSequence = nil
			for i := int64(0); i < int64(len(tempLength)); i ++ {
				finalSequence = append(finalSequence, tempLength[i])
			}
		}
		fmt.Println(tempLength)
		fmt.Println(finalSequence)
		fmt.Println("---------------")
	}
}
