package main

import (
    "fmt"
    "strconv"
)

func main() {
    inPut := 1098924424973
    fmt.Println((inPut&(1<<39))==0)
    in := "1098924424974"
    features, err := strconv.ParseUint(in, 10, 64)
    fmt.Println("features = ", features)
    if err != nil || (features&(1<<39) == 0) {
        fmt.Println("false")
    }

}
