package main

import (
    "fmt"
    "net/url"
)

func main() {
    var a = `www.operanewsapp.com%2Fmz%2Fpt%2Fshare%2Fdetail%3Fnews_id%3D970a95517d7ff510ad70118ce54b17f5_mz%26news_entry_id%3D734447bd200727pt_mz%26open_type%3Dtranscoded%26request_id%3DPUSH_44023829-8ec7-4ba9-9cee-aa488711c702%26from%3Dmini_push`
    ret, _ := url.QueryUnescape(a)
    fmt.Println(ret)

    var b = `https://img-src.feednews.com/assets/v2/image%2Fdefault_images%2Fothers%2Fadult-2178656_1920.jpg?`
    ret = url.QueryEscape(b)

    ret1, _ := url.QueryUnescape(ret)
    fmt.Println(ret1)

    b = `https://img-src.feednews.com/assets/v2/image%2Fdefault_images%2Fothers%2Fadult-2178656_1920.jpg?`
    ret1, _ = url.QueryUnescape(b)
    fmt.Println(ret1)
}
