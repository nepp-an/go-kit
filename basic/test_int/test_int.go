package main

import "fmt"

func main() {
    fmt.Println(0x00010001)
    fmt.Println(0x00020002)
    fmt.Println(0x00040004)

    fmt.Println((0x0000FFFF&0x00000000) > 0)

    fmt.Println(fmt.Sprintf("%.2f%%", (float32(1.1)/float32(2.2))*float32(100)))
}