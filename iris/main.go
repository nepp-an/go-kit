package main

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/pprof"

    "github.com/kataras/iris/v12/middleware/recover"
)

func main() {
    // Creates an iris application without any middleware by default
    app := iris.New()

    // Global middleware using `UseRouter`.
    //
    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    app.UseRouter(recover.New())

    // Per route middleware, you can add as many as you desire.
    app.HandleMany("GET", "/debug/pprof /debug/pprof/{action:path}", pprof.New())
    // Listen and serve on 0.0.0.0:8080
    app.Listen(":8080")
}
