package main

import (
    "fmt"
    "hash/fnv"
    "sync"
)

func main() {
    test := make(map[string]bool)
    appe := make(map[string]bool)
    testMap := make(map[string]string)
    test["zzz"] = true
    appe["bbb"] = true
    appe["aaa"] = true

    test[""] = true

    testMap[""] = "1234"
    test = appe

    fmt.Print(testMap)
    fmt.Println()
    test = map[string]bool{}
    for i:=0;i<2;i++ {
        var tokens = map[string][]string{}
        fmt.Println("1 ",tokens)
        tokens["12"] = append(tokens["12"], "123")
        tokens["12"] = append(tokens["12"], "1234")
        fmt.Println(tokens)
    }
    safeMap := NewSafeMap()
    fmt.Println("len(safeMap.locks)")

    fmt.Println(len(safeMap.locks))


    var abMap = make(map[string]interface{})
    abMap["asdsda"] = 2
    if abMap["asdsda"] == nil {
        fmt.Println("map yes")
    } else {
        fmt.Println(abMap["asdsda"])

    }

}

type SafeMap struct {
    locks []*sync.Mutex
    store []map[string]string
}

func NewSafeMap() SafeMap {
    return SafeMap{
        locks: []*sync.Mutex{{},{}},
        store: []map[string]string{},
    }
}

func hash(k string) int {
    h := fnv.New32a()
    h.Write([]byte(k))
    return int(h.Sum32())
}

func (m *SafeMap) GetLock(k string) *sync.Mutex {
    idx := hash(k) % len(m.locks)
    return m.locks[idx]
}

func (m *SafeMap) GetStore(k string) map[string]string {
    idx := hash(k) % len(m.locks)
    return m.store[idx]
}

func (m *SafeMap) Get(k string) string {
    lock := m.GetLock(k)
    lock.Lock()
    defer lock.Unlock()

    return m.GetStore(k)[k]
}

func (m *SafeMap) Set(k, v string) {
    lock := m.GetLock(k)
    lock.Lock()
    defer lock.Unlock()

    m.GetStore(k)[k] = v
}
