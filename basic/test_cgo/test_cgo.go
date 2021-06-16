// hello.go
package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

type CChar C.char

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
