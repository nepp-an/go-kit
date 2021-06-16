package main

import (
    "fmt"
    "strconv"
)

func main()  {
    fmt.Println(1<<42)
    fmt.Println(features("17591598841613"))
}

func rangeBitwiseAnd(m int, n int) int {
    var shift uint
    for m < n {
        m, n = m >> 1, n >> 1
        shift++
    }
    //m is the same part
    //shift is the position of the same prefix

    return m << shift
}

func rangeBitwiseAnd2(m int, n int) int {
    for m < n {
        n &= n - 1
    }
    return n
}

func features(input string) bool {
    features, err := strconv.ParseUint(input, 10, 64)
    if err != nil || (features&(1<<42) == 0) {
        // Do not support instaclips
        return false
    }
    return true
}
