package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const (
		updateIntervalHour = 10
		updateBeginMinute  = 5
	)
	loc, _:= time.LoadLocation("Asia/Chongqing")
	fmt.Println(time.Now().In(loc))
	fmt.Println("format: ", time.Now().Format("2006-01-01"))
	now := time.Now().UTC()
	fmt.Println(now)
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1.UTC().Unix())

	next := time.Now().In(loc)
	fmt.Println("next", next)
	fmt.Println("now", time.Now().UTC())

	fmt.Println(next.Sub(time.Now().UTC()))

	if strings.Contains(next.Sub(time.Now().UTC()).String(), "-") {
		fmt.Println("true")
	}
	now = time.Now().UTC()
	d, _ = time.ParseDuration("-24h")
	d1 = now.Add(d)
	fmt.Println(d1.Format("2006-01-02"))

	fmt.Println("time.Now().UTC().Hour()", time.Now().UTC().Hour())


	fmt.Println(time.Unix(1594161141729/1000, 0))

	var timeString = "2020-07-09T13:01:08Z"
	var timeLayOut = "2006-01-02T15:04:05Z"
	ts, _  := time.Parse(timeLayOut, timeString)
	fmt.Println(ts)

	//20200720/07:15:43
	var timeString1 = "20200720/07:15:43"
	var timeLayOut1 = "20060102/15:04:05"
	ts1, _  := time.Parse(timeLayOut1, timeString1)
	fmt.Println(ts1)

	fmt.Println(time.Unix(	1595264618,1000000000))

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("20060102T15"))
}
