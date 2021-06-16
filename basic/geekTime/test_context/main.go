package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type Result struct {
   Hit string
   Err error
}

//// Search runs query on a backend and returns the result.
//type Search func(query string) Result

//
//// First runs query on replicas and returns the first result.
//func First(query string, replicas ...Search) Result {
//    c := make(chan Result, len(replicas))
//    search := func(replica Search) { c <- replica(query) }
//    for _, replica := range replicas {
//        go search(replica)
//    }
//    return <-c
//}

// Search runs query on a backend and returns the result.
type Search func(ctx context.Context, query string) Result

var getResult Search = func(ctx context.Context, query string) Result {
    time.Sleep(100*time.Millisecond)
    return Result{
        Hit: query,
        Err: nil,
    }
}


func main() {
    result := First(context.Background(), "begin", getResult, getResult, getResult)
    fmt.Println(result.Hit)
    c := make(chan os.Signal, 1)

    signal.Notify(c, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    input := <-c
    fmt.Printf("%#v", input)
}

// First runs query on replicas and returns the first result.
func First(ctx context.Context, query string, replicas ...Search) Result {
    c := make(chan Result, len(replicas))
    ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
    defer cancel()
    search := func(replica Search) {
        ret := replica(ctx, query)
        if ret.Err == nil {
            c <- ret
        }

        //c <- replica(ctx, query)
    }
    for _, replica := range replicas {
        go search(replica)
    }
    select {
    case <-ctx.Done():
        return Result{Err: ctx.Err()}
    case r := <-c:
        return r
    }
}

