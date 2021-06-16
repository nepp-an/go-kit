package main

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"math/rand"
	"strings"
	"time"
)

var EsClient *elastic.Client

const (
	ZEUS       = "zeus"
	ARES       = "ares"
	AGGREGATOR1 = "aggregator"
)

var ESSettingFlag = map[string]bool{
	ZEUS:       false,
	ARES:       false,
	AGGREGATOR1: false,
}

func main() {
	EsClient, err := elastic.NewClient(elastic.SetErrorLog(log.StandardLogger()),
		elastic.SetURL(strings.Split("http://localhost:9200/", ",")...),
		elastic.SetMaxRetries(5))
	if err != nil {
		panic(err)
	}
	defer EsClient.Stop()
	index := "debug_aggregator-20191129"
	exists, err:=EsClient.IndexExists(index).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists{
		_, err = EsClient.CreateIndex(index).Body(AGGREGATOR).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

	//time.Sleep(time.Second*5)
	//_, err = EsClient.DeleteIndex(index).Do(context.Background())
	//if err != nil {
	//	log.Errorf("DeleteESIndex failed, err is ", err.Error())
	//}
	for i:=0 ;i<10 ;i++  {
		_, err = EsClient.Index().
			Index(index).
			Type("aggregator").
			Id(GetRandomString(10)).
			BodyString(content).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
	}


	//_, err = EsClient.IndexPutSettings(index).BodyString(Setting).Do(context.Background())
	//if err != nil {
	//	//panic(err)
	//}
	//fmt.Println(ESSettingFlag[ZEUS])
	//ESSettingFlag[ZEUS] = true
	//fmt.Println(ESSettingFlag[ZEUS])
	//fmt.Println(ESSettingFlag[AGGREGATOR])
	//fmt.Println(ESSettingFlag[ARES])
	//ticker:=time.NewTicker(time.Second*5)
	//go func() {
	//	for _ =range ticker.C {
	//		println("test")
	//	}
	//}()
	//time.Sleep(time.Minute)

	resp, err := EsClient.
		Search(index).
		Query(elastic.NewQueryStringQuery("user_id:9f9c7e0610ea744d731b75e02d0e101d8485d51")).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d", resp.Hits.TotalHits)

}

//生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


const Setting = `
{
    "settings":{
        "index.mapping.total_fields.limit":99999,
        "number_of_replicas":0,
        "refresh_interval":"1m",
        "index.translog.durability":"async",
        "index.translog.flush_threshold_size":"1024mb"
	}
}
`
const content = `

  {

	"request_id":"SUBSCRIBE_SUGGESTION_7c537d23-ee3d-409a-9476-8ced3c3c48e3",

	"user_id":"9f9c7e0610ea744d731b75e02d0e101d8485d51",

	"country":"gh",

	"language":"en",
 
    "product":"./img/5.jpg",
	"session_info":""
 
  }`
const content_1 = `
{
    "request_id":"SUBSCRIBE_SUGGESTION_7c537d23-ee3d-409a-9476-8ced3c3c48e3",
    "user_id":"9f9c7e0610ea744d731b75e02d0e101d8485d51",
    "module":"aggregator",
    "ts":1574868149,
    "country":"gh",
    "language":"en",
    "product":"news",
    "abtest":"[{"module":"video_infra","eid":"4192","vid":"10862","parameters":{"skip_peer_review":0}},{"module":"infra","eid":"3870","vid":"10092","parameters":{}},{"module":"zeus","eid":"4339","vid":"11362","parameters":{"random_ratio":0.2}},{"module":"ares","eid":"4294","vid":"11191","parameters":{}},{"module":"apollo","eid":"2639","vid":"7279","parameters":{"category_v2_score_filter_threshold":0.9,"use_category_filter":false,"use_category_v2_score_filter":true,"use_related_model_for_others":true}},{"module":"infra","eid":"4244","vid":"11047","parameters":{"related_tag_mtd":"combined"}},{"module":"adx","eid":"4279","vid":"11151","parameters":{"prepare_use_shuffle_merge":1}},{"module":"ares","eid":"2892","vid":"7800","parameters":{"broadcast":{"long_ratio":0.4}}},{"module":"aggregator","eid":"4379","vid":"11511","parameters":{"abnormal_threshold":0.5,"combine_model_name":"control","greater_than_registration_day":1,"news_begin":3,"news_diff":2,"news_video_rank_enable":true,"random_range":3}},{"module":"push","eid":"4377","vid":"11502","parameters":{"lr_model_name":"push_v2"}},{"module":"apollo","eid":"4167","vid":"10794","parameters":{}},{"module":"apollo","eid":"3679","vid":"9639","parameters":{"click_coef_combine_stay":0.5,"stay_coef_combine_stay":0.5,"stay_model_name":"stay_classify_v4","turn_on_lr_stay":true}},{"module":"zeus","eid":"4303","vid":"11223","parameters":{"hot_topic_name":{"turn_off":true}}},{"module":"video_rank","eid":"4336","vid":"11349","parameters":{"model_name":"normpt_pairwise_v2"}},{"module":"video_recaller","eid":"4055","vid":"10513","parameters":{"item_cf_withnews":1}},{"module":"zeus","eid":"","vid":"","parameters":{"tracking_turn_on":true}}]",
    "request_uri":"/gh/en/v1/news/subscribe/suggestion?product=news&features=33554351&ac=3g&lang=en&category_id=&uid=9f9c7e0610ea744d731b75e02d0e101d8485d5a1&exclude=summary",
    "news_raw_rsp":"{"experiments":{"":"","2639":"7279","2892":"7800","3679":"9639","3870":"10092","4055":"10513","4167":"10794","4192":"10862","4244":"11047","4279":"11151","4294":"11191","4303":"11223","4336":"11349","4339":"11362","4377":"11502","4379":"11511"},"request_id":"SUBSCRIBE_SUGGESTION_7c537d23-ee3d-409a-9476-8ced3c3c48e3","dedup_prefix":"","articles":[{"articles":[{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Nana Addo Dankwa Akufo-Addo is a Ghanaian politician currently the President of Ghana.","lastupdate_time":1568734466,"posts":288,"publisher_id":"news_dlsub_00888","publisher_logo":"http://res.feednews.com/assets/v2/aaa13b94f73c2d08a9ce583679cfb571?width=255\u0026height=255\u0026quality=hq","publisher_name":"Nana Akufo-Addo","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":147477,"likes":0,"group_id":"news_cate_00013","group_name":"Politics","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Shatta Michy","lastupdate_time":1568734466,"posts":11,"publisher_id":"news_subent_07591","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Shatta Michy","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":71875,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Wendy Shay is a Ghanaian Afropop and Afrobeat singer.","lastupdate_time":1568734466,"posts":66,"publisher_id":"news_subent_05390","publisher_logo":"http://res.feednews.com/assets/v2/eb0f7812bac6d7e383a751ae215a1998?width=255\u0026height=255\u0026quality=hq","publisher_name":"Wendy Shay","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":121291,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Ronald Kwaku Dei Appiah better known a Bisa Kdei is a Ghanaian native solo singer and record producer.","lastupdate_time":1568734466,"posts":7,"publisher_id":"news_dlsub_00896","publisher_logo":"http://res.feednews.com/assets/v2/8d62777d06899d98ab9737879bc30a40?width=255\u0026height=255\u0026quality=hq","publisher_name":"Bisa Kdei","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":63337,"likes":0,"group_id":"news_dlcate_00020","group_name":"Singer","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"John Dramani Mahama is a Ghanaian politician who served as President of Ghana from 24 July 2012 to 7 January 2017.","lastupdate_time":1568734466,"posts":160,"publisher_id":"news_dlsub_00815","publisher_logo":"http://res.feednews.com/assets/v2/50512e0f84eac6251f7a93bf620796b7?width=255\u0026height=255\u0026quality=hq","publisher_name":"John Mahama","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":154879,"likes":0,"group_id":"news_cate_00013","group_name":"Politics","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Emmanuel Botway, known by his stage name Kwaw Kese, also popularly called Abodam, is a Ghanaian hiplife artist.","lastupdate_time":1568734466,"posts":25,"publisher_id":"news_dlsub_00953","publisher_logo":"http://res.feednews.com/assets/v2/e6d6fec188411bc4affc699508b62b64?width=255\u0026height=255\u0026quality=hq","publisher_name":"Kwaw Kese","reason":"Recommendation for You","suggest_reason":"Recommendation for You","subscribers":68985,"likes":0,"group_id":"news_cate_00002","group_name":"Music","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Gladstorm Kwabena Akwaboah Jnr., known by his stage name Akwaboah Jnr, is a Ghanaian singer-songwriter and producer from Mampong Beposo.","lastupdate_time":1568734466,"posts":37,"publisher_id":"news_dlsub_00934","publisher_logo":"http://res.feednews.com/assets/v2/4ece58640ab634deeeea19983a38bdb5?width=255\u0026height=255\u0026quality=hq","publisher_name":"Akwaboah Jnr","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":61055,"likes":0,"group_id":"news_dlcate_00020","group_name":"Singer","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about King Promise","lastupdate_time":1568734466,"posts":8,"publisher_id":"news_subent_07724","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"King Promise","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":80906,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Kwesi Arthur","lastupdate_time":1568734466,"posts":43,"publisher_id":"news_subent_07681","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Kwesi Arthur","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":101823,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Fella Makafui is an actress, known for Kada River.","lastupdate_time":1568734466,"posts":33,"publisher_id":"news_subent_09151","publisher_logo":"http://res.feednews.com/assets/v2/3535b553bdd80601bbf165ce60cc794e?width=255\u0026height=255\u0026quality=hq","publisher_name":"Fella Makafui","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":107240,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Kelvyn Boy","lastupdate_time":1568734466,"posts":76,"publisher_id":"news_subent_23232","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Kelvyn Boy","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":62202,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Asamoah Gyan is a Ghanaian professional footballer who plays for Kayserispor and captains the Ghanaian national team.","lastupdate_time":1568734466,"posts":65,"publisher_id":"news_dlsub_00893","publisher_logo":"http://res.feednews.com/assets/v2/7f038829f1ddaaa78739864d8f5141dc?width=255\u0026height=255\u0026quality=hq","publisher_name":"Asamoah Gyan","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":137182,"likes":0,"group_id":"news_dlcate_00001","group_name":"Football Star","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Nana Ama McBrown","lastupdate_time":1568734466,"posts":27,"publisher_id":"news_subent_13462","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Nana Ama McBrown","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":91362,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Efia Odo","lastupdate_time":1568734466,"posts":113,"publisher_id":"news_subent_04240","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Efia Odo","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":80202,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Chinyere Yvonne Okoro is a Ghanaian actress. She received Ghana Movie Awards Best Actress Award in 2010.","lastupdate_time":1568734466,"posts":25,"publisher_id":"news_dlsub_00812","publisher_logo":"http://res.feednews.com/assets/v2/6ea268218932e4abafb8eb6ce32b5e65?width=255\u0026height=255\u0026quality=hq","publisher_name":"Yvonne Okoro","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":102146,"likes":0,"group_id":"news_dlcate_00013","group_name":"Actress","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about DJ Switch","lastupdate_time":1568734466,"posts":59,"publisher_id":"news_subent_04694","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"DJ Switch","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":52284,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Okyeame Kwame","lastupdate_time":1568734466,"posts":39,"publisher_id":"news_subent_18402","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Okyeame Kwame","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":53186,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Kofi Kinaata","lastupdate_time":1568734466,"posts":19,"publisher_id":"news_subent_20407","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Kofi Kinaata","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":73872,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Follow for all news about Yaa Jackson","lastupdate_time":1568734466,"posts":9,"publisher_id":"news_subent_06721","publisher_logo":"http://res.feednews.com/assets/v2/cbfa619430cbcb589cd921eaadb58b3d?width=255\u0026height=255\u0026quality=hq","publisher_name":"Yaa Jackson","reason":"Recommendation for You","suggest_reason":"Recommendation for You","subscribers":71935,"likes":0,"group_id":"news_subcate_00001","group_name":"Others","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"},{"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"","news_entry_id":"","news_id":"","original_url":"","publisher_info":{"description":"Nigerian singer, songwriter, performer and actress. Follow for her updates.","lastupdate_time":1568734466,"posts":176,"publisher_id":"news_dlsub_00042","publisher_logo":"http://res.feednews.com/assets/v2/db5f3cbdb9a83b7a96c47dc508b52eb0?width=255\u0026height=255\u0026quality=hq","publisher_name":"Tiwa Savage","reason":"Most People Follow Today","suggest_reason":"Most People Follow Today","subscribers":146141,"likes":0,"group_id":"news_cate_00002","group_name":"Music","type":1,"publisher_type":0,"league_table_url":"","subscribable":true,"follow_flag":0,"artifact_id":"","enable_show_in_v2":false},"title":"","type":"publisher"}],"category_meta":{"name":"","id":"","icon":""},"feedback_reason_v2":null,"images":null,"logo":"http://res.feednews.com/assets/v2/565aee71a02f38ecbc88961a24828227","news_entry_id":"","news_id":"","original_url":"","title":"Follow Topics You May Like","type":"publishers"}]}"
}
`

const aggre = `
{
  "mappings": {
    "aggregator": {
      "properties": {
        "abtest": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "country": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "language": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "module": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "news_raw_rsp": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "product": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "request_id": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "request_uri": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "ts": {
          "type": "date",
          "format": "epoch_second"
        },
        "user_id": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        }
      }
    }
  }
}
`

const AGGREGATOR = `
{
  "mappings": {
    "aggregator": {
      "properties": {
        "abtest": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "country": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "language": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "module": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "news_raw_rsp": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "product": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "request_id": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "request_uri": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "ts": {
          "type": "date",
          "format": "epoch_second"
        },
        "user_id": {
          "type": "keyword"
        }
      }
    }
  }
}
`