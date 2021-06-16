package main

import (
    "context"
    "fmt"
    "time"
)
func main() {
    ctx := context.WithValue(context.Background(), "default", "open")
    ctx1 := context.WithValue(ctx, "category", "sports")

    if value, ok := ctx1.Value("default").(string); ok {
        fmt.Println(value)
    }
    if value, ok := ctx1.Value("category").(string); ok {
        fmt.Println(value)
    }
    ctx2 := context.WithValue(ctx1, "category", "ok")
    if value, ok := ctx2.Value("default").(string); ok {
        fmt.Println(value)
    }
    if value, ok := ctx2.Value("category").(string); ok {
        fmt.Println(value)
    }

    ctx, cancel := context.WithTimeout(ctx2, 1 * time.Second)
    defer cancel()
    select {
        case <- ctx.Done():
            fmt.Println("timeout")
        case <- time.After(3 * time.Second):
            fmt.Println("later")
    }
    return
}