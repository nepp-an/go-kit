package main

import (
    "fmt"
    "sort"
)

func main() {
    input := []int{10,1,2,7,6,1,5}
    target := 8
    fmt.Println(combinationSum2(input, target))

}

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var res [][]int

    var backtrack func(int, int, []int)
    backtrack = func(idx int, sum int, nums []int) {

        if sum == target {
            t := make([]int, len(nums))
            copy(t, nums)
            var flag = false
            for j:=0; j<len(res);j++ {
                if testEq(res[j], t) {
                    flag = true
                }
            }
            if !flag{
                res = append(res, t)
            }
        }
        if sum > target {
            return
        }
        for i := idx; i < len(candidates); i++ {
            backtrack(i+1, sum + candidates[i], append(nums,candidates[i]))
        }
    }
    backtrack(0, 0, []int{})
    return res
}

func testEq(a, b []int) bool {
    // If one is nil, the other must also be nil.
    if (a == nil) != (b == nil) {
        return false;
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}