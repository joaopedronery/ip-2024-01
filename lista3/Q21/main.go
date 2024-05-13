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
            if v%2 == 0{
                fmt.Printf("%d\n", v)
            }
        }
       
        sort.Sort(sort.Reverse(sort.IntSlice(s)))
       
        for _, v := range s{
            if v%2 != 0{
                fmt.Printf("%d\n", v)
            }
        } 
}
