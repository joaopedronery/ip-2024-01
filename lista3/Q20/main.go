package main

import (
	"fmt"
	"math"
)

type ponto struct {
	x, y, z float64
}

func main() {
	var n int
	fmt.Scan(&n)

	pontos := make([]ponto, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&pontos[i].x, &pontos[i].y, &pontos[i].z)
	}

	for i := 1; i < n; i++ {
		vx := vetor(pontos[i-1], pontos[i])
		cmax := maiorabs(vx)
		fmt.Printf("%.2f\n", math.Abs(cmax))
	}
}

func vetor(p1, p2 ponto) ponto {
	return ponto{p2.x - p1.x, p2.y - p1.y, p2.z - p1.z}
}

func maiorabs(v ponto) float64 {
	max := math.Abs(v.x)
	if math.Abs(v.y) > max {
		max = math.Abs(v.y)
	}
	if math.Abs(v.z) > max {
		max = math.Abs(v.z)
	}
	return max
}