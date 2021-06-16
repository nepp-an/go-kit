package main

//func main() {
//    input := []int{2,3, 5}
//    target := 8
//    fmt.Println(combinationSum(input, target))
//
//}

func combinationSum(candidates []int, target int) [][]int {
   var res [][]int

   var backtrack func(int, int, []int)
   backtrack = func(idx int, sum int, nums []int) {

       if sum == target {
           res = append(res, nums) //nums 不是值传递吗
       }
       if sum >= target {
           return
       }
       for i := idx; i < len(candidates); i++ {
           backtrack(i, sum + candidates[i], append(nums,candidates[i]))
       }
   }
   backtrack(0, 0, []int{})
   return res
}

//func combinationSum(candidates []int, target int) [][]int {
//    var res [][]int
//
//    var backtrack func(int, int, []int)
//    backtrack = func(idx int, sum int, nums []int) {
//        fmt.Println("nums: ",nums)
//        fmt.Println("sum: ",sum)
//        fmt.Println("res:",res)
//        fmt.Println("---------")
//
//        if sum == target {
//            t := make([]int, len(nums))
//            copy(t, nums)
//            //t := nums
//            //res = append(res, t)
//            res = append(res, nums)
//
//            fmt.Println("nums: ",nums)
//            fmt.Println("sum: ",sum)
//            fmt.Println("res:",res)
//            fmt.Println("=================")
//        }
//        if sum >= target {
//            return
//        }
//        for i := idx; i < len(candidates); i++ {
//            backtrack(i, sum + candidates[i], append(nums,candidates[i]))
//        }
//    }
//    backtrack(0, 0, []int{})
//    return res
//}
