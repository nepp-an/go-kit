package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "strings"
    "time"

    log "github.com/Sirupsen/logrus"
)

func main() {
    fmt.Println(time.Now().Unix())

    testIndex := "9907000000"
    if testIndex[:2] == "99" {
        if testIndex[2] == '0' {
            testIndex = testIndex[3:4] + "0000"
        } else {
            testIndex = testIndex[2:4] + "0000"
        }
    }
    fmt.Println(testIndex)

    url := "/us/en/v1/news/category/city_us_en5494"
    index := strings.Index(url, "city_")

    fmt.Println(index)
    fmt.Println(strings.ToUpper(url[index:]))
    testStr := `"Fauci": Coronavirus pandemic showed 'undeniable effects of racism in our society'`
    if strings.Contains(testStr, `"`) {
        log.Infof(testStr)
        log.Infof(strings.ReplaceAll(testStr, `"`, `\\"`))
    }

    data := `{"Name":"\\"Happy\\""}`
    response := &Response{
        Code: 1,
        Msg:  testStr,
        Data: data,
    }
    _, err := json.Marshal(&response)
    if err != nil {
        fmt.Println("err", err)
    }
    testStr = "\"Je perverser, desto 'and' besser\""
    testStr = strings.ReplaceAll(testStr, `"`, `"`)
    testStr = strings.ReplaceAll(testStr, `'`, `\'`)
    testMap := map[string]string{
        "title":testStr,
    }
    byteBuf := bytes.NewBuffer([]byte{})
    encoder := json.NewEncoder(byteBuf)
    encoder.SetEscapeHTML(false)
    encoder.Encode(testMap)
    fmt.Println(byteBuf.String())

    res := byteBuf.String()
    fmt.Println(res)
    res = strings.ReplaceAll(res, `\"`, `\\"`)
    res = strings.ReplaceAll(res, `'`, `\'`)


    fmt.Println(string(res))
    a := fmt.Sprintf("{'ab_li': '%s', 'uri': '%s','extra': '%s'}", res,res,res)
    ret, _ := json.Marshal(a)
    fmt.Println(string(ret))



}

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}
type People struct {
    Name string
}