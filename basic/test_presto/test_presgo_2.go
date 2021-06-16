package main

import (
	"fmt"
	"github.com/colinmarc/go-presto"
	"time"
)

func main() {

//fmt.Println(GetDAU("20200920", "us", "en"))
	//fmt.Println(GetPushUV("mini","20200715", "in", "en"))
	//
	//fmt.Println(GetDAU("20200715", "id", "id"))
	//fmt.Println(GetPushUV("mini","20210222", "ng", "en"))

	fmt.Println(GetPushPV("","20210305", "us", "en"))
}

//https://auto-presto-infra.proxy.op-mobile.opera.com
func GetDAU(date string, nation string, lang string) float64 {
	sql := `
	select
		p_date,
		sum(uv)
	from
		opera_datamarts.stat_for_active_user
	where
		p_product in ('news', 'newseu') and
		grouping_id = 7 and
		p_date in ('%s') and
		p_nation in ('%s') and p_language in ('%s')
	group by
		p_date`

	sql = fmt.Sprintf(sql, date, nation, lang)
	//fmt.Println("sql:", sql)
	query, err := presto.NewQuery("http://presto-node.proxy.op-mobile.opera.com:18180", "", "", "hive", "opera_datamarts", sql)
	if err != nil {
		fmt.Printf(err.Error())
		return -1
	}

	if row, _ := query.Next(); row != nil {
		//fmt.Println(row...)
		return row[1].(float64)
	}
	return -1
}

func GetPushPV(product string, date string, nation string, lang string) [4]float64 {
	sql := `
	select action_type, count(*) as pv
from (select distinct action_type, device_id, news_entry_id FROM opera_push.pps_events where p_product in ('news') and p_date = '%s' and p_nation = '%s' and p_language = '%s' ) t
    group by action_type`

	sql = fmt.Sprintf(sql, date, nation, lang)
	fmt.Println("sql:", sql)

	var ret [4]float64
	rows := make([][]interface{}, 0, 1)
	for i := 0; i < 2; i++ {
		time.Sleep(200 * time.Millisecond)
		query, err := presto.NewQuery("http://presto-node.proxy.op-mobile.opera.com:18180", "", "", "hive", "opera_push", sql)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if query == nil {
			continue
		}
		for {
			row, _ := query.Next()
			if row == nil {
				break
			}
			if actionType, ok := row[0].(string); ok {
				if actionType == "receive" {
					ret[0] = row[1].(float64)
				}
			}
			if actionType, ok := row[0].(string); ok {
				if actionType == "show" {
					ret[1] = row[1].(float64)
				}
			}
			if actionType, ok := row[0].(string); ok {
				if actionType == "click" {
					ret[2] = row[1].(float64)
				}
			}
			if actionType, ok := row[0].(string); ok {
				if actionType == "send_ok" {
					ret[3] = row[1].(float64)
				}
			}


			rows = append(rows, row)
		}
		fmt.Println(ret[0])
		if ret[0] > 0 {
			break
		}

	}

	fmt.Println(rows)
	//if row, _ := query.Next(); row != nil {
	//	fmt.Println(row...)
	//}
	return ret
}


func GetPushUV(product string, date string, nation string, lang string) []float64 {
	sql := `
	select action_type, count(distinct  device_id) as uv
	from opera_push.pps_events
	where p_product = '%s' and p_date = '%s' and p_nation = '%s' and p_language = '%s'
	group by action_type`

	sql = fmt.Sprintf(sql, product, date, nation, lang)
	//fmt.Println("sql:", sql)
	query, err := presto.NewQuery("http://presto-node.proxy.op-mobile.opera.com:18180", "", "", "hive", "opera_push", sql)
	if err != nil {
		fmt.Println(err.Error())
		return []float64{}
	}
	var ret []float64
	rows := make([][]interface{}, 0, 1)
	for {
		row, _ := query.Next()
		if row == nil {
			break
		}
		if row[0].(string) == "receive" {
			ret = append(ret, row[1].(float64))
		}
		if row[0].(string) == "show" {
			ret = append(ret, row[1].(float64))
		}
		if row[0].(string) == "click" {
			ret = append(ret, row[1].(float64))
		}
		if row[0].(string) == "send_ok" {
			ret = append(ret, row[1].(float64))
		}

		rows = append(rows, row)
	}
	fmt.Println(rows)
	//if row, _ := query.Next(); row != nil {
	//	fmt.Println(row...)
	//}
	return ret
}