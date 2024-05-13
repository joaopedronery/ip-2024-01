package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	resultados := make([]string, T)

	for t := 0; t < T; t++ {
		var n int
		fmt.Scan(&n)

		sequencia := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&sequencia[i])
		}

		menordist := 2001
		comparacao := 0

		for i := 1; i < n; i++ {
			for j := 0; j < i; j++ {
				dist := abs(sequencia[i] - sequencia[j])
				comparacao++
				if dist < menordist {
					menordist = dist
				}
			}
		}

		resultados[t] = fmt.Sprintf("%d %d", menordist, comparacao)
	}

	for _, resultado := range resultados {
		fmt.Println(resultado)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}