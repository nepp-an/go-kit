//package main
//
//import "fmt"
//
//func main() {
//    var arr []string
//    arr = append(arr, "1")
//    for _, v := range arr {
//        fmt.Println(v)
//        arr = append(arr, "2")
//    }
//    n, m := 3, 2
//    //create dynamic array
//    dp := make([][]bool, n)
//    for i:=0; i<n;i++ {
//        dp[i] = make([]bool, m)
//    }
//    fmt.Println(dp)
//    var a = []int{1}
//    var b = []int{2}
//    var ret [][]int
//    var num  = a
//    ret = append(ret, num)
//    fmt.Println(ret)
//    num = b
//    fmt.Println(ret)
//}

package main

import (
    "bytes"
    "fmt"
    "runtime/debug"
)

func main() {
    nums := []string{}
    nums = append(nums, "123")
    nums = append(nums, "456")
    nums = append(nums, "7")

    test(nums)
    fmt.Println(nums)

    num := []int{}
    num = append(num, 1)
    num = append(num, 2)
    num = append(num, 3)

    fmt.Println(num)
    test1(num)
    fmt.Println(num)

    var notSupportPath = []string{
        "assets/v1/crop",
        "assets/v1/tn",
        "assets/v1/smartcrop",
    }

    for i := range notSupportPath {
        fmt.Println(notSupportPath[i])
    }
    //path := []byte("AAAA/BBBBBBBBB")
    //sepIndex := bytes.IndexByte(path,'/')
    fmt.Println(len("3aee0441210203fr_fr"))

    debug.FreeOSMemory()


    var testArr = []int{1,2,3,4}
    testEqual := testArr
    var testCopy = make([]int, len(testArr))
    copy(testCopy, testArr)
    fmt.Println(testArr, testEqual, testCopy)
    testArr[1] = 6
    fmt.Println(testArr, testEqual, testCopy)

    var strArr = "5678"
    fmt.Println(strArr[len(strArr)-4:])

    fmt.Println(getDepartIndex("100000"))
    fmt.Println(getDepartIndex("100001"))
    fmt.Println(getDepartIndex("9910110023"))
    fmt.Println(getDepartIndex("9907030170"))


    path := []byte("AAAA/BBBBBBBBB")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex:sepIndex] //Full slice expression
    fmt.Println(dir1)
}


func test1(num []int) {
    // append无法更改值。
    num = append(num, 4)
    num[2] = 4
}


func test(nums []string) {
    nums = append(nums, "8")
}


func getDepartIndex(index string) string {
    if len(index) <= 5 {
        return ""
    }
    // 9911100001 -> 110010
    // 9901010001 -> 100010
    if index[:2] == "99" {
        if index[2] == '0' {
            index = index[3:4] + "00" + index[4:6]
        } else {
            index = index[2:4] + "00" + index[4:6]
        }

    }
    return index
}