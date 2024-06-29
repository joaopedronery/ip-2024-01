package main

import (
	"fmt"
	"sort"
)

func main() {
	integersList := []int{10, 12, 2, 6, 22, 21, 9, 5}
	sortList(integersList)
	fmt.Print(integersList)
}

func sortList(vetor []int) {
	sort.Ints(vetor)
}