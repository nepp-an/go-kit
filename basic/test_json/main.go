package main

import (
    "encoding/json"
    "fmt"
)

type Resp struct {
    SuccessNum int `json:"success"`
    Failure int `json:"failure"`
}

func main() {
    msg := "{\"success\":1,\"failure\":1,\"illegal_tokens\":[\"123IQAAAACy0hj3AAA4bnhnxBVmXI0xDUDmM29mrRoGjZa-cooZMaxERaXbNfBZS9KWx_jgvhFYq9N3-R-L3mh2nWFVhQIGVATdC0_osrhPScVAhhSSbg\"]}"
    var resp Resp
    err := json.Unmarshal([]byte(msg), &resp)
    if err == nil {
        fmt.Printf("%#v", resp)
    }

}
