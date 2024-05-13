package main

import "fmt"

func main() {
   
    var n int
   
    fmt.Scan(&n)
   
    x := make([]int, n, n)
   
    for i := 0; i < n; i++{
        fmt.Scan(&x[i])
    }

    for _, v := range x{
        fmt.Printf("%d ", v)
    }
}
