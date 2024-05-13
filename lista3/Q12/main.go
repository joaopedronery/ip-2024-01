package main

import (
    "fmt"
    "sort"
)

func main() {
       
        var n int
        fmt.Scan(&n)
        s := make([]int, n, n)
        for i := 0; i < n; i++{
        fmt.Scan(&s[i])
        }
       
        fmt.Println("SAIDA")
        sort.Ints(s)
        for _, v := range s{
            fmt.Printf("%d\n", v)
        }
       
}