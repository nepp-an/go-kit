package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "hash/fnv"
)

func main() {
    fmt.Println(getIndex("test"))
    fmt.Println(getIndex("mddddasssda"))
    for i :=0; i<5; i++ {
        fmt.Println(Hash("1131135b54f5fdbaba14be9417f557e6")%100)
        fmt.Println(Hash("b05e4297dd7f45fd8a747fa8f15b6b44")%16)
        fmt.Println(Hash("adce36b5677c72f4a74700f8dc021c10-3")%16)
    }
}

func getIndex(key string) int {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    h := fnv.New32a()
    err := enc.Encode(key)
    if err != nil {
        return 0
    }
    h.Write(buf.Bytes())
    return (int)(h.Sum32()) % 5
}

func Hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}