package main


import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	useTera := flag.Bool("tera", false, "tera")
	fmt.Println(*useTera)
	a := make(map[string]interface{})
	fmt.Println(a)
	var testArraynil []*int
	fmt.Println(testArraynil==nil)
	go func() {

	}()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("1%10=%d",1%10 )
	var out = 1.00
	for i:=1;i<10;i++{
		out = out*1.03
	}
	fmt.Printf("out = %f", out)
	time.Sleep(5 * time.Second)
}
