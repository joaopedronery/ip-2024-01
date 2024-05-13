package main

import (
	"fmt"
)

func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := getMax(arr)
	vCount := make([]int, max + 1)
	vOrd := make([]int, len(arr))

	for _, num := range arr {
		vCount[num]++
	}

	for i := 1; i <= max; i++ {
		vCount[i] += vCount[i - 1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		vOrd[vCount[arr[i]] - 1] = arr[i]
		vCount[arr[i]]--
	}

	return vOrd
}

func main() {
	var N int
	for {
		fmt.Scan(&N)

		if N == 0 {
			break
		}

		arr := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&arr[i])
		}

		orderedArr := countingSort(arr)

		for i, num := range orderedArr {
			if i == len(orderedArr) - 1 {
				fmt.Print(num)
			} else {
				fmt.Print(num, " ")
			}
		}
		fmt.Println()
	}
}
