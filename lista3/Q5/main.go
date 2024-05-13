package main

import (
	"fmt"
)

func main() {
	var entrada [][]int

	for {
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		x := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&x[i])
		}

		entrada = append(entrada, x)
	}

	var saida []string
	for _, x := range entrada {
		maior := -1
		indice := -1

		for i, v := range x {
			if v > maior {
				maior = v
				indice = i
			}
		}

		saida = append(saida, fmt.Sprintf("%d %d", indice, maior))
	}

	for _, vs := range saida {
		fmt.Println(vs)
	}
}