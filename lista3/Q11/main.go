package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	x := make([]int, n, 1000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
		fmt.Printf("%d ", x[i])
	}

	fmt.Print("\n")

	for i := n - 1; i >= 0; i--{
		fmt.Printf("%d ", x[i])
	}

	fmt.Print("\n")

	m := -1

	for i := 0; i < n; i++{
		if x[i] > m{
			m = x[i]
		}
	}
	fmt.Println(m)

	mn := 99999

	for i := 0; i < n; i++{
		if x[i] < mn{
			mn = x[i]
		}
	}
	fmt.Print(mn)
}