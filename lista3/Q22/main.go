package main

import "fmt"

func main() {
  var n, d int
  var numero string
  var resultados []string

  for {
    fmt.Scanf("%d %d", &n, &d)
    if n == 0 && d == 0 {
      break
    }
    fmt.Scan(&numero)

    aRemover := n - d
    digitos := make([]byte, 0, n)

    for i := 0; i < n; i++ {
      digitoAtual := numero[i]
 
      for len(digitos) > 0 && digitos[len(digitos)-1] < digitoAtual && aRemover > 0 {
        digitos = digitos[:len(digitos)-1]
        aRemover--
      }

      digitos = append(digitos, digitoAtual)
      
    }
      
      resultados = append(resultados, string(digitos[:d])) 
    }

    for _, resultado := range resultados{
      fmt.Println(resultado)
    }
}
