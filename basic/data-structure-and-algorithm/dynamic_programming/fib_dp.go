package main

import "fmt"

func fib(n int) int {
    var dp = make(map[int]int)
    if n == 1 || n == 2 {
        return 1
    }
    // 初始值
    dp[1] = 1
    dp[2] = 1
    for i := 3; i <= n; i++ {
        // 状态转移方程
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}


func main() {
    fmt.Println(fib(3))
    fmt.Println(fib(4))
    fmt.Println(fib(5))
}