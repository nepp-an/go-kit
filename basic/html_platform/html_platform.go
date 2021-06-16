package main

import (
	"fmt"
	"html/template"
	"net"
	"os"
)

type Person struct {
	Name string
	age  string
}

func main() {
	t, err := template.New("recallListPage").Parse(HTMLPage)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
	fmt.Println(GetLocalIP())
}
const HTMLPage = `
<!DOCTYPE html>
<html>

<head>
    <title>
    </title>
</head>

<body>
    <p>
        <!--p代表当前对象-->
        hello,{{.Name}}
        {{.}}
    </p>
</body>
</html>`

func GetLocalIP() string {
	AddressArray, _ := net.InterfaceAddrs()
	for _, v := range AddressArray {
		if ip, ok := v.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}
	return ""
}