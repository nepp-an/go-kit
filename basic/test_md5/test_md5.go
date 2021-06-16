package main

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "fmt"
)

var secret = "5826af512974ea8e42510486ba3941bb"


func main() {
    fmt.Println(MD5("deleteFile"))
    fmt.Println(ComputeHmacSha256("DownloadAccessKey", secret))
    fmt.Println(ComputeHmacSha256("UploadAccessKey", secret))

}

func MD5Bytes(s []byte) string {
    ret := md5.Sum(s)
    return hex.EncodeToString(ret[:])
}

//计算字符串MD5值
func MD5(s string) string {
    return MD5Bytes([]byte(s))
}

func ComputeHmacSha256(message string, secret string) string {
    key := []byte(secret)
    h := hmac.New(sha256.New, key)
    h.Write([]byte(message))
    //	fmt.Println(h.Sum(nil))
    sha := hex.EncodeToString(h.Sum(nil))
    //	fmt.Println(sha)

    //	hex.EncodeToString(h.Sum(nil))
    return base64.StdEncoding.EncodeToString([]byte(sha))
}
