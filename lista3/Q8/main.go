package main

import "fmt"

func main() {
       
        var n int
        fmt.Scan(&n)
        y := make([]int, n, n)
        for i := 0; i < n; i++{
        fmt.Scan(&y[i])
        }
        for _, v := range y{
           fmt.Printf("%b\n", v)
       }
}