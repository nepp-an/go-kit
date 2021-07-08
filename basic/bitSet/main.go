package main

import "fmt"

func main() {
    var bitSet [26]int
    bitSet[2] = 1
    bitSet['a'-'a'] = 1
    fmt.Println(bitSet)
}


