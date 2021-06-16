package test_channel

import (
    "fmt"
    "runtime"
    "strconv"
    "testing"
    "time"
)

func timeout(f func(chan bool)) error {
    done := make(chan bool)
    go f(done)
    select {
    //receive 超时控制
    case <-done:
        fmt.Println("done")
        return nil
    case <-time.After(15*time.Millisecond):
        return fmt.Errorf("timeout")
    }
}

func test(t *testing.T, f func(chan bool)) {
    t.Helper()
    for i := 0; i < 100; i++ {
        timeout(f)
    }
    time.Sleep(time.Second * 2)
    t.Log("number "+strconv.Itoa(runtime.NumGoroutine()))
}

func doGoodthing(done chan bool) {
    time.Sleep(10*time.Millisecond)
    //尝试发送
    select {
    case done <- true:
        fmt.Println("send ok")
    default:
        fmt.Println("send failed")
        return
    }
}

func TestGoodTimeout(t *testing.T) {
    test(t, doGoodthing)
}
