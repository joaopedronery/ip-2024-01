package main

import(
	"fmt"
)

func main() {

	var n int

	fmt.Scan(&n)

	x := make([]int, n, 1000000)

	for i := 0; i < n; i++{
		fmt.Scan(&x[i])
	}

	contar := make(map[int]int)

	for _, num := range x {
        contar[num]++
    }

	var mnum, mcontar int

	for i, v := range contar {
        if v > mcontar {
            mnum = i
            mcontar = v
        }
    }
	fmt.Printf("%d\n%d", mnum, mcontar)
}