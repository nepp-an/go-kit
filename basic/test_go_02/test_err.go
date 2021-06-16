package main

import (
    "fmt"
    "github.com/pkg/errors"
    "os"
)

func main() {
    _, err := os.Open("")
    err = errors.Wrap(err, "failed")
    err = errors.Wrap(err, "failed 2")

    fmt.Printf("err is %+v", err)

    fmt.Println("------")

}
