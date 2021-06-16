package main

import (
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func download(url string) ([]byte, error){
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("user-agent", "Mozilla/5.0")
    req.Header.Add("cookie", cookie)
    client := &http.Client{}
    response, err := client.Do(req)
    //response, err := http.Get(url)
    if err != nil {
        return []byte{}, err
    }
    return ioutil.ReadAll(response.Body)
}

func main() {
    var url string
    url = "https://www.usnews.com/dims4/USNEWS/ad03f32/2147483647/thumbnail/970x647/quality/85/?url=http%3A%2F%2Fmedia.beam.usnews.com%2Ff3%2Ff8%2F7a1caf3542289f5c4dd09e51f569%2F201222newsvariant-editorial.jpg"
    url = "https://www.thenewstribune.com/latest-news/x8qegj/picture249892288/alternates/FREE_1140/IMG_0352.jpg"
    imgByte, err := download(url)
    if err != nil {
        panic(err)
    }
    //add format verified
    imgFormat := http.DetectContentType(imgByte)
    if strings.Contains(imgFormat, "text/html") {
        fmt.Println("failed")
    }
    fmt.Println(imgFormat)
    fmt.Println(base64.StdEncoding.EncodeToString(imgByte))
}


var cookie = `bm_mi=19FDD56323E7683098299BBF1D73D3C0~JrSLO6fkRRU2dyDcSh9bCU5wWOWdt8sc9RCeaV6VyIQtHNuVMqj+GikghLYZ7Yxqk0yn1Ju5v+vdFqtsRdMGjXwVRvVImYJFE1Xfie8VK3qeuSiUxyX/Mme9vW+KlG7wHZwiTXqtFA4qd8NEpMvcWNF8J/4L9WyuuAwOEWxqi6lAgBdOnukHAN+kWAet9n5KDGxpBsGXIdP9MEbq9ellxPezwO+EIs1nkN4eclxiSwM4mRHek3woSBh9fKa8EmN0e2z4dfqwTrI3Euridy3NUvoxQm1GJDFSOAfPZ9NXMRQy5mFHM5TGlSOVSjzZG9YpTd1cCPBMA5Osd8KUeSpQOg==; Domain=.usnews.com; Path=/; Max-Age=7191; HttpOnly`