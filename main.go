package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	fmt.Println(diagonalDifference(arr))
}

func diagonalDifference(arr [][]int32) int32 {
	var acc int32
	var acc2 int32
	for i, line := range arr {
		acc = acc + line[i]
		acc2 = acc2 + line[len(line)-1-i]
		continue
	}
	fixedAmount := acc - acc2
	intergerResult := math.Abs(float64(fixedAmount))
	return int32(intergerResult)
}
