package main

import (
    "encoding/json"
    "fmt"
    "github.com/micro/go-micro/util/log"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
    "time"
)

type UserCityRsp struct {
    ID   string `json:"uid"`
    City string `json:"city"`
}

func main() {
    //var uids = []string{
    //    "4a1ee0a1c678d641d75802e994b8b52101285498",
    //    "4a1ee0a1c678d641d75802e994b8b52101285498",
    //    "2cbda5fde2d2ab4b865a40c32424b92101075498",
    //}
    //getCount()
    fmt.Println(getDocCount("fcm_token_za_en"))
    fmt.Println(getDate())
}

func getCitysFromDevices(deviceIds []string) map[string]string {
    var ret = make(map[string]string)
    client := &http.Client{}
    url := "https://db.ams.op-mobile.opera.com/user_city?uids=" + strings.Join(deviceIds, ",")
    resp, err := client.Get(url)
    if err != nil {
        fmt.Errorf("getCitysFromDevices client.Get failed err is %s", err.Error())
        return ret
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Errorf("getCitysFromDevices ioutil.ReadAll failed err is %s", err.Error())
        return ret
    }
    var cityArr []UserCityRsp
    err = json.Unmarshal(data, &cityArr)
    if err != nil {
        fmt.Errorf("getCitysFromDevices json.Unmarshal failed err is %s", err.Error())
        return ret
    }
    //fmt.Println(cityArr)
    for i := range cityArr {

        ret[cityArr[i].ID] = cityArr[i].City

    }

    return ret
}

func getDate() int {
    dateStr := time.Now().Format("20060102")
    date, _ := strconv.Atoi(dateStr)
    return date
}

func getCount() {

    client := &http.Client{}
    url := "https://sg-profile-push-1-9201.singa.op-mobile.opera.com/_cat/count/fcm_token_za_en"
    resp, err := client.Get(url)
    if err != nil {
        fmt.Errorf("getCitysFromDevices client.Get failed err is %s", err.Error())
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Errorf("getCitysFromDevices ioutil.ReadAll failed err is %s", err.Error())
    }
    if len(strings.Split(string(data), " ")) != 3 {
        fmt.Println(string(data))
    }
    ret := strings.Replace(string(data), "\n", "", -1)
    docCount := strings.Split(ret, " ")[2]
    fmt.Println(strconv.Atoi(docCount))

    //fmt.Println(cityArr)

}


func getDocCount(indexName string) int {
    var docCount int
    client := &http.Client{}
    url := "https://sg-profile-push-1-9201.singa.op-mobile.opera.com/_cat/count/" + indexName
    resp, err := client.Get(url)
    if err != nil {
        log.Errorf("getDocCount client.Get failed err is %s", err.Error())
        return docCount
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Errorf("getCitysFromDevices ioutil.ReadAll failed err is %s", err.Error())
        return docCount
    }
    if len(strings.Split(string(data), " ")) != 3 {
        log.Infof(string(data))
        return docCount
    }
    retStr := strings.Replace(string(data), "\n", "", -1)
    docCountStr := strings.Split(retStr, " ")[2]

    docCount, _ = strconv.Atoi(docCountStr)

    return docCount
}