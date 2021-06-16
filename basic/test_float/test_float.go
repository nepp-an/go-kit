package main
import (
    "fmt"
)

func main() {
    var oriWidth = 1200
    var oriHeight = 800
    var width, height = 450, 336
    fmt.Printf("%d %d %d %d", oriWidth, oriHeight,width,height)
    fmt.Println()
    fmt.Printf("%f %f %d %d", float32(oriWidth)/float32(width), float32(oriHeight)/float32(height),
        int(float32(width) * float32(oriHeight/height)) ,int(float32(height) * float32(oriWidth/width)))
    if float32(oriWidth)/float32(width) > float32(oriHeight)/float32(height) {
        oriWidth = int(float32(width) * float32(oriHeight)/float32(height))
    } else {
        oriHeight = int(float32(height) * float32(oriWidth)/float32(width))
    }
    fmt.Println()
    fmt.Printf("%d %d %d %d", oriWidth, oriHeight,width,height)

    fmt.Println()

    fmt.Println(float32(250)/float32(1280) > 15)

}
