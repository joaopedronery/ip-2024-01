package main

import (
  "fmt"
  "math"
  "strings"
)

func calculaFrequencias(s string) [5]int {
  var freq [5]int
  s = strings.ToLower(s)
  for _, ch := range s {
    switch ch {
    case 'a', 'á':
      freq[0]++
    case 'e', 'é':
      freq[1]++
    case 'i', 'í':
      freq[2]++
    case 'o', 'ó':
      freq[3]++
    case 'u', 'ú':
      freq[4]++
    }
  }
  return freq
}

func distanciaEuclidiana(f1 [5]int, f2 [5]int) float64 {
  var soma float64
  for i := 0; i < 5; i++ {
    diferenca := float64(f1[i] - f2[i])
    soma += diferenca * diferenca
  }
  return math.Sqrt(soma)
}

func main() {
  var saida string
  fmt.Scanf("%s", &saida)

  partes := strings.Split(saida, ";")

  if len(partes)!= 2 {
    fmt.Println("FORMATO INVALIDO!")
    return
  }

  string1 := partes[0]
  string2 := partes[1]

  freq1 := calculaFrequencias(string1)
  freq2 := calculaFrequencias(string2)

  distancia := distanciaEuclidiana(freq1, freq2)

  fmt.Printf("(%d,%d,%d,%d,%d)\n", freq1[0], freq1[1], freq1[2], freq1[3], freq1[4])
  fmt.Printf("(%d,%d,%d,%d,%d)\n", freq2[0], freq2[1], freq2[2], freq2[3], freq2[4])
  fmt.Printf("%.2f\n", distancia)
}
