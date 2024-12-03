package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    s2 := append(s[:0], s[1:]...)
    s2 = append(s2, 3)
    fmt.Println(s)
    fmt.Println(s2)
    fmt.Println(&s[0])
    fmt.Println(&s2[0])
}
