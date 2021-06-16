package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    sjson "github.com/bitly/go-simplejson"
    "io/ioutil"
    "net/http"
    "strconv"
)

func main() {
    fmt.Println(GetPush("opera", "20210312", "us", "en"))
}

type kylinBody struct {
    Sql string `json:"sql"`
    OffSet int `json:"offset"`
    Limit  int `json:"limit"`
    AcceptPartial bool `json:"acceptPartial"`
    Project string `json:"project"`
}

func GetPush(product, date, nation, lang string) [9]int {

    sql := `select
    p_date,
    count(distinct click_device_id) as click_uv,
    count(distinct show_device_id) as show_uv,
    count(distinct send_ok_device_id) as send_ok_uv,
    count(distinct receive_device_id) as receive_uv,
    sum(send_ok_pv) as send_ok_pv,
    sum(receive_pv) as receive_pv,
    sum(show_pv) as show_pv,
    sum(click_pv) as click_pv
from
    opera_push.pps_stat_view
where
    p_date in ('%s') and p_product in ('%s') and p_nation = '%s' and p_language = '%s'
group by
    p_date`
    query := fmt.Sprintf(sql, date, product, nation, lang)
    postBody := kylinBody{
        Sql: query,
        OffSet: 0,
        Limit: 5000,
        AcceptPartial:false,
        Project: "opera_stat",
    }
    body, _ := json.Marshal(postBody)
    url := "https://auto-kylin.proxy.op-mobile.opera.com/kylin/api/query"
    request, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
    request.Header.Add("Authorization", "Basic QURNSU46S1lMSU4=")
    request.Header.Add("Content-Type", "application/json")

    client := http.DefaultClient
    resp, err := client.Do(request)
    if err != nil {
        return [9]int{}
    }
    var ret [9]int
    if resp.Body != nil {
        bodyRaw, _ := ioutil.ReadAll(resp.Body)
        j, _ := sjson.NewJson(bodyRaw)
        resultArr, _ := j.Get("results").Array()
        tmp := resultArr[0].([]interface{})
        fmt.Println(tmp)
        for i := range tmp {

           v := tmp[i].(string)
           fmt.Println(v)
           a, _ := strconv.Atoi(v)

           ret[i] = a
        }
        return ret
    }
    return [9]int{}

}

func GetKylinNewsPush(date, nation, lang string, ios bool) [9]int {

    sql := `select
    p_date,
    count(distinct click_device_id) as click_uv,
    count(distinct show_device_id) as show_uv,
    count(distinct send_ok_device_id) as send_ok_uv,
    count(distinct receive_device_id) as receive_uv,
    sum(send_ok_pv) as send_ok_pv,
    sum(receive_pv) as receive_pv,
    sum(show_pv) as show_pv,
    sum(click_pv) as click_pv
from
    opera_push.pps_stat_view
where
    p_date in ('%s') and p_product in ('news', 'newseu') and p_nation = '%s' and p_language = '%s'
group by
    p_date`
    if ios {
        sql = `select
    p_date,
    count(distinct click_device_id) as click_uv,
    count(distinct show_device_id) as show_uv,
    count(distinct send_ok_device_id) as send_ok_uv,
    count(distinct receive_device_id) as receive_uv,
    sum(send_ok_pv) as send_ok_pv,
    sum(receive_pv) as receive_pv,
    sum(show_pv) as show_pv,
    sum(click_pv) as click_pv
from
    opera_push.pps_stat_view
where
    p_date in ('%s') and p_product in ('iosnews', 'iosnewseu') and p_nation = '%s' and p_language = '%s'
group by
    p_date`
    }
    query := fmt.Sprintf(sql, date, nation, lang)
    postBody := kylinBody{
        Sql: query,
        OffSet: 0,
        Limit: 5000,
        AcceptPartial:false,
        Project: "opera_stat",
    }
    body, _ := json.Marshal(postBody)
    url := "https://auto-kylin.proxy.op-mobile.opera.com/kylin/api/query"
    request, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
    request.Header.Add("Authorization", "Basic QURNSU46S1lMSU4=")
    request.Header.Add("Content-Type", "application/json")

    client := http.DefaultClient
    resp, err := client.Do(request)
    if err != nil {
        return [9]int{}
    }
    var ret [9]int
    if resp.Body != nil {
        bodyRaw, _ := ioutil.ReadAll(resp.Body)
        j, _ := sjson.NewJson(bodyRaw)
        resultArr, _ := j.Get("results").Array()
        tmp := resultArr[0].([]interface{})
        fmt.Println(tmp)
        for i := range tmp {

            v := tmp[i].(string)
            fmt.Println(v)
            a, _ := strconv.Atoi(v)

            ret[i] = a
        }
        return ret
    }
    return [9]int{}

}