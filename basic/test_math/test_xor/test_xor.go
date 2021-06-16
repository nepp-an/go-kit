package main

import (
    "fmt"
    "hash/fnv"
)

func main() {
    a := "key"
    b := "secret"
    fmt.Println(a,b)
}

func Hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}