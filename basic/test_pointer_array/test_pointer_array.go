package main

import "fmt"

func main() {
	type FetcherKey struct {
		Key       string
		IndexType string
		Country   string
		Language  string
		Count     int32
	}

	type indexTerm struct {
		indexName string
		indexType string
		key       string
		count     int
		begin     int
		end       int
		tsMin     int
		tsMax     int
		fetchKey  FetcherKey
	}
	key1 := FetcherKey{
		Key:       "economy",
		IndexType: "",
		Country:   "",
		Language:  "",
		Count:     0,
	}
	key2 := FetcherKey{
		Key:       "policy",
		IndexType: "",
		Country:   "",
		Language:  "",
		Count:     0,
	}
	var shortTermIndex []indexTerm
	for i:=0; i<2;i++ {
		var temp indexTerm
		if i == 0 {
			temp = indexTerm{
				indexName: "",
				indexType: "",
				key:       "",
				count:     0,
				begin:     0,
				end:       0,
				tsMin:     0,
				tsMax:     0,
				fetchKey:  key1,
			}
		} else {
			temp = indexTerm{
				indexName: "",
				indexType: "",
				key:       "",
				count:     0,
				begin:     0,
				end:       0,
				tsMin:     0,
				tsMax:     0,
				fetchKey:  key2,
			}
		}

		shortTermIndex = append(shortTermIndex, temp)
	}

	var fetchKey  []*FetcherKey
	//指针数组不能用range
	for _, v := range shortTermIndex {
		fmt.Printf("v is %d \n", &v)
		fetchKey = append(fetchKey, &(v.fetchKey))
		//fmt.Printf("fetchKey is %#v \n", fetchKey[0])
		//fmt.Printf("&(v.fetchKey) is %#v \n", &(v.fetchKey))
	}
	for _, v := range fetchKey {
		fmt.Printf("fetchKey.Key is %#v \n", v.Key)
	}
}
