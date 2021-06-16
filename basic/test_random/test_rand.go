package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		//time.Sleep(time.Duration(rand.Int63n(10)+5)*time.Second)
		//fmt.Println(rand.Int63n(10)+5)

		var groupIds = []string{
			"news-1", "news-2", "news-3", "news-4",
		}
		fmt.Println(groupIds[rand.Intn(len(groupIds))])
			//extra["group_name"] = "news"
		fmt.Println(rand.Intn(120))

	}



	//for ; ; {
	//	//fmt.Println("first is ",rand.Intn(100))
	//	rand.Seed(time.Now().Unix())
	//	//fmt.Println(rand.Intn(100))
	//	if rand.Intn(100) == 99 {
	//		fmt.Println("99")
	//
	//	}
	//	time.Sleep(10*time.Millisecond)
	//}
	return
}