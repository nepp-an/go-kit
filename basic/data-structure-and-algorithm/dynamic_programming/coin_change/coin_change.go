package main

import (
    "fmt"
    "math"
)

func main() {
    var coins = []int{1, 2, 5}
    num := coinChange(coins, 11)
    fmt.Println(num)
}

// 凑零钱问题
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    return helper(coins, amount)
}

func helper(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    if amount < 0 {
        return -1
    }
    res := math.MaxInt32
    for _, coin := range coins {
        tmp := helper(coins, amount-coin)
        if tmp == -1 {
            continue
        }
        res = min(res, tmp+1)
    }
    if res == math.MaxInt32 {
        return -1
    }
    return res
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}