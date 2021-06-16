package main

import (
    "fmt"
    "time"
)

func main() {
    cl := []string{"de_de","us_en", "gb_en", "fr_fr"}
    NewAllCLTokenCache(cl)
    token1 := PushToken{
        id:   "qwer",
        time: time.Now().Unix(),
    }
    token2 := PushToken{
        id:   "asdf",
        time: time.Now().Unix()+100,
    }
    token3 := PushToken{
        id:   "zxcv",
        time: time.Now().Unix()+200,
    }
    token4 := PushToken{
        id:   "rtyu",
        time: time.Now().Unix()+300,
    }

    token5 := PushToken{
        id:   "uiop",
        time: time.Now().Unix()+400,
    }

    AddToken("us_en", "qwer", &token1)
    AddToken("us_en", "asdf", &token2)
    rangeTokenCache("us_en")
    fmt.Println(getLastModified("us_en"))
    rangeTokenCache("fr_fr")
    fmt.Println()

    AddToken("us_en", "zxcv", &token3)
    AddToken("us_en", "tyui", &token4)
    rangeTokenCache("us_en")
    fmt.Println(getLastModified("us_en"))

    AddToken("us_en", "uiop", &token5)
    rangeTokenCache("us_en")
    fmt.Println(getLastModified("us_en"))
}