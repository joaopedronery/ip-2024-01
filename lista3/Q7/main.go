package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var entrada string
	var N int
	for {
		n, _ := fmt.Scanln(&N)
		if n == 0 || N == 0 {
			break
		}
		var x []int
		for i := 0; i < N; i++ {
			var num int
			fmt.Scan(&num)
			x = append(x, num)
		}
		M := x[0]
		for _, num := range x {
			if num > M {
				M = num
			}
		}
		sort.Ints(x)
		freq := make([]int, M+1)
		for _, num := range x {
			freq[num]++
		}
		contar := 0
		for i := 0; i <= M; i++ {
			contar += freq[i]
			entrada += fmt.Sprintf("(%d) %d\n", i, contar)
		}
	}
	fmt.Print(strings.TrimSpace(entrada))
}