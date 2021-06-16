package main

import (
    "fmt"
    "sync"
)

type User struct {
    sync.Mutex
    name string
}

var UserArr = make(map[string]*User)

func main() {
    UserArr["a"] = &User{
        Mutex: sync.Mutex{},
        name:  "A",
    }
    test1 := UserArr["a"]
    test1.Lock()
    fmt.Println(test1.name)
    defer test1.Unlock()

    test2 := UserArr["a"]
    test2.Lock()
    fmt.Println(test2.name)

    defer test2.Unlock()
}