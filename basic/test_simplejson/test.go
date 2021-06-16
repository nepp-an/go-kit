package main

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "math/rand"
)

var json_str string = `{
    "articles":[
        {
            "articles":[
                {
                    "category":"News_Politics",
                    "comment_num":17,
                    "feedback_reason_v2":{
                        "not_interested":[
                            {
                                "content":"arrestations leaders",
                                "description":"",
                                "id":"keyword",
                                "text":"",
                                "value":"arrestations leaders"
                            },
                            {
                                "content":"jean-louis billon",
                                "description":"",
                                "id":"keyword",
                                "text":"",
                                "value":"jean-louis billon"
                            },
                            {
                                "content":"yeclo.com",
                                "description":"",
                                "id":"domain",
                                "text":"",
                                "value":"yeclo.com"
                            },
                            {
                                "content":"yeclo.com",
                                "description":"",
                                "id":"Others",
                                "text":"",
                                "value":"yeclo.com"
                            }
                        ],
                        "report":null
                    },
                    "images":null,
                    "logo":"https://res.6chcdn.feednews.com/assets/v1/logo/xxl/yeclo_com",
                    "news_entry_id":"7fd986aa201109fr_ci",
                    "news_id":"4f56e6dd3d5a8c77c9145669cd202ea9",
                    "open_type":"transcoded",
                    "original_url":"https://www.yeclo.com/arrestations-leaders-de-lopposition-billon-sort-du-silence/",
                    "publish_timestamp":1604920761,
                    "share_url":"http://opr.news/7fd986aa201109fr_ci?client=news",
                    "social_info":{
                        "shared_count":0,
                        "shared_people":[
                            "http://res.6chcdn.feednews.com/v0/avatar/A7/IMG_1115132156_bigger.jpg",
                            "http://res.6chcdn.feednews.com/v0/avatar/A7/lA2WoEnT_bigger.jpg"
                        ],
                        "shared_people_names":[

                        ],
                        "total_dislike_count":1,
                        "total_like_count":39,
                        "total_shared_count":37
                    },
                    "source":"yeclo.com",
                    "source_name":"yeclo.com",
                    "thumbnail":[
                        "http://res.6chcdn.feednews.com/assets/v2/6f188813a22a26939f7f14e92f53c34b?width=450&height=336&quality=hq&category=CI_News_Politics&format=webp"
                    ],
                    "timestamp":1604920761,
                    "title":"Arrestations leaders de l'Opposition : Billon sort du silence",
                    "transcoded_url":"https://stories.6chcdn.feednews.com/news/detail/4f56e6dd3d5a8c77c9145669cd202ea9?app_version=7.26.5&country=ci&language=fr",
                    "type":"normal"
                }
            ],
            "category_meta":{
                "icon":"",
                "id":"",
                "name":""
            },
            "feedback_reason_v2":null,
            "images":null,
            "logo":"",
            "more_id":"2508",
            "news_entry_id":"",
            "news_id":"",
            "original_url":"",
            "title":"Dernières actualités",
            "type":"cms_topic_tab_news"
        },
        {
            "articles":[
                {
                    "category":"News_Politics",
                    "feedback_reason_v2":{
                        "not_interested":[
                            {
                                "content":"laurent gbagbo",
                                "description":"",
                                "id":"keyword",
                                "text":"",
                                "value":"laurent gbagbo"
                            },
                            {
                                "content":"cojep de charles blé",
                                "description":"",
                                "id":"keyword",
                                "text":"",
                                "value":"cojep de charles blé"
                            },
                            {
                                "content":"linfodrome.com",
                                "description":"",
                                "id":"domain",
                                "text":"",
                                "value":"linfodrome.com"
                            },
                            {
                                "content":"linfodrome.com",
                                "description":"",
                                "id":"Others",
                                "text":"",
                                "value":"linfodrome.com"
                            }
                        ],
                        "report":null
                    },
                    "flags":1,
                    "images":null,
                    "logo":"https://res.6chcdn.feednews.com/assets/v1/logo/xxl/linfodrome_com",
                    "news_entry_id":"4bafe817201109fr_ci",
                    "news_id":"126eda89c305997dd9d45ee70c4c9606",
                    "open_type":"original",
                    "original_url":"https://www.linfodrome.com/vie-politique/62725-video-des-premiers-signes-de-vie-du-president-affi-n-guessan-le-cojep-de-charles-ble-goude-denonce-un-montage-grossier-et-degradant",
                    "publish_timestamp":1604917486,
                    "share_url":"http://opr.news/4bafe817201109fr_ci?client=news",
                    "social_info":{
                        "shared_count":0,
                        "shared_people":[
                            "http://res.6chcdn.feednews.com/v0/avatar/A8/Ni_eRiNX_bigger.jpg",
                            "http://res.6chcdn.feednews.com/v0/avatar/A3/406229_10100240174605264_30300188_42868708_1864886765_n_bigger.jpg"
                        ],
                        "shared_people_names":[

                        ],
                        "total_like_count":7,
                        "total_shared_count":3
                    },
                    "source":"linfodrome.com",
                    "source_name":"linfodrome.com",
                    "thumbnail":[
                        "http://res.6chcdn.feednews.com/assets/v2/126eda89c305997dd9d45ee70c4c9606?width=450&height=336&quality=hq&category=CI_News_Politics&format=webp"
                    ],
                    "timestamp":1604917486,
                    "title":"Vidéo des premiers signes de vie du président Affi N’guessan: le Cojep de Charles Blé Goudé dénonce un montage grossier et dégradant",
                    "transcoded_url":"https://stories.6chcdn.feednews.com/news/detail/126eda89c305997dd9d45ee70c4c9606?app_version=7.26.5&country=ci&language=fr",
                    "type":"normal"
                }
            ],
            "category_meta":{
                "icon":"",
                "id":"",
                "name":""
            },
            "feedback_reason_v2":null,
            "images":null,
            "logo":"",
            "more_id":"2509",
            "news_entry_id":"",
            "news_id":"",
            "original_url":"",
            "title":"Réactions",
            "type":"cms_topic_tab_news"
        }
    ],
    "dedup_prefix":"EDITOR_CURATED_NEWS_",
    "experiments":{
        "4978":"13185",
        "5445":"14220",
        "5812":"15002",
        "6493":"16474",
        "6554":"16611",
        "6655":"16823",
        "6659":"16832",
        "6664":"16844",
        "6683":"16883",
        "6740":"16997",
        "6753":"17026",
        "6758":"17035",
        "6770":"17060"
    },
    "request_id":"EDITOR_CURATED_NEWS_8858e90c-2ddf-4c94-a48a-71a38fe6db59",
    "response_ts":1605235375
}`

var video_str = `{
    "result": {
        "message": "success",
        "code": 0
    },
    "request_id": "45a967031ba95d368266b7b9d479ed8a",
    "articles": [
        {
            "news_entry_id": "",
            "news_id": "",
            "type": "cms_topic_tab_news",
            "enable_more_button": false,
            "more_id": "2509",
            "title": "Prolongation du confinement",
            "articles": [
                {
                    "news_entry_id": "P-eVE3696K6",
                    "news_id": "",
                    "type": "insta_clip",
                    "infra_feedback": "{}",
                    "enable_more_button": false,
                    "insta_obj": {
                        "id": "P-eVE3696K6",
                        "publisher": {
                            "id": "07e426150785480fba88d90e60cceaa9",
                            "name": "BFMTV",
                            "description": "Bienvenue sur mon profil",
                            "portrait": "http://res.feednews.com/assets/v2/9d6af31853576332ebdbe207e1e84fba?height=80&width=80",
                            "points": 18,
                            "video_points": 18,
                            "new_post_count": 0,
                            "follower_count": 0,
                            "follow_flag": 0,
                            "influencer": true
                        },
                        "type": "clip",
                        "description": "Est-ce que je pourrai rouvrir mon salon de coiffure le 1er décembre ? - BFMTV répond à vos questions",
                        "thumbnail": {
                            "width": 1280,
                            "height": 720,
                            "url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_Xgb1VxU7m.jpg_640x360",
                            "format": "jpg",
                            "preview_url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_Xgb1VxU7m.jpg_200x112"
                        },
                        "board": {
                            "id": "B-grG9kGmqd",
                            "name": "InstaNews",
                            "remark": "Nowto",
                            "description": "InstaNews",
                            "type": "board",
                            "cover": {
                                "width": 360,
                                "height": 110,
                                "url": "http://6chcdn-res.feednews.com/instaclip/img/B-B0QDDEwdz-B.jpg",
                                "format": "jpg"
                            },
                            "thumbnail": {
                                "width": 56,
                                "height": 56,
                                "url": "http://6chcdn-res.feednews.com/instaclip/img/B-B0QDDEwdz-A.png",
                                "format": "png"
                            },
                            "post_count": 53316
                        },
                        "tags": [
                            {
                                "id": "T-Q5mBO6Jgb",
                                "name": "Gros titres",
                                "type": "tag",
                                "cover": {
                                    "width": 360,
                                    "height": 110,
                                    "url": "http://6chcdn-res.feednews.com/instaclip/img/B-o56NN7Q1a-B.jpg",
                                    "format": "jpg"
                                },
                                "thumbnail": {
                                    "width": 56,
                                    "height": 56,
                                    "url": "http://6chcdn-res.feednews.com/instaclip/img/B-o56NN7Q1a-A.png",
                                    "format": "png"
                                },
                                "post_count": 15178,
                                "follow_flag": 0
                            }
                        ],
                        "like_count": 0,
                        "dislike_count": 0,
                        "comment_count": 0,
                        "share_count": 0,
                        "liked": false,
                        "disliked": false,
                        "favored": false,
                        "video": {
                            "id": "Q06_Xgb1VxU7m",
                            "thumbnail": {
                                "width": 1280,
                                "height": 720,
                                "url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_Xgb1VxU7m.jpg_640x360",
                                "format": "jpg",
                                "preview_url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_Xgb1VxU7m.jpg_200x112"
                            },
                            "duration": 87,
                            "size": 2227200,
                            "format": "mp4",
                            "width": 854,
                            "height": 480,
                            "url": "http://video.6chcdn.com/instaclip06/video/Q06_Xgb1VxU7m.mp4",
                            "download_url": "http://video.6chcdn.com/instaclip06/video/Q06_Xgb1VxU7m.mp4",
                            "has_audio": true,
                            "source_type": "normal",
                            "rates": {
                                "p144": 1670400,
                                "p240": 2227200,
                                "p360": 3340800,
                                "p480": 4454400,
                                "p540": 5568000,
                                "p720": 8908800
                            }
                        },
                        "share_url": "https://instaclips.op-mobile.opera.com/post/P-eVE3696K6?share=1&country=fr&language=fr",
                        "create_time": 1605224299000,
                        "state": 100,
                        "algo_data": "{\"feed\": \"https://www.youtube.com/user/BFMTV\", \"tags\": [], \"treetags\": []}",
                        "video_update_lock": false,
                        "infra_feedback": "{\"comment_count\":0,\"loc\":\"-\",\"infra_bit_rate\":\"playlist\",\"like_count\":0,\"gender\":\"unknown\",\"create_time\":1605224299000,\"follower_count\":0,\"share_count\":0,\"width\":854,\"video_full_duration\":87,\"dislike_count\":0,\"follow_flag\":0,\"height\":480}",
                        "entrance_type": "CRAWLER"
                    }
                },
                {
                    "news_entry_id": "P-xXO5NL82L",
                    "news_id": "",
                    "type": "insta_clip",
                    "infra_feedback": "{}",
                    "enable_more_button": false,
                    "insta_obj": {
                        "id": "P-xXO5NL82L",
                        "publisher": {
                            "id": "07e426150785480fba88d90e60cceaa9",
                            "name": "BFMTV",
                            "description": "Bienvenue sur mon profil",
                            "portrait": "http://res.feednews.com/assets/v2/9d6af31853576332ebdbe207e1e84fba?height=80&width=80",
                            "points": 18,
                            "video_points": 18,
                            "new_post_count": 0,
                            "follower_count": 0,
                            "follow_flag": 0,
                            "influencer": true
                        },
                        "type": "clip",
                        "description": "Puis-je aller chercher mes filles chez leur mère ce week-end ? - BFMTV répond à vos questions",
                        "thumbnail": {
                            "width": 1280,
                            "height": 720,
                            "url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_CBSsvH77r.jpg_640x360",
                            "format": "jpg",
                            "preview_url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_CBSsvH77r.jpg_200x112"
                        },
                        "board": {
                            "id": "B-grG9kGmqd",
                            "name": "InstaNews",
                            "remark": "Nowto",
                            "description": "InstaNews",
                            "type": "board",
                            "cover": {
                                "width": 360,
                                "height": 110,
                                "url": "http://6chcdn-res.feednews.com/instaclip/img/B-B0QDDEwdz-B.jpg",
                                "format": "jpg"
                            },
                            "thumbnail": {
                                "width": 56,
                                "height": 56,
                                "url": "http://6chcdn-res.feednews.com/instaclip/img/B-B0QDDEwdz-A.png",
                                "format": "png"
                            },
                            "post_count": 53316
                        },
                        "tags": [
                            {
                                "id": "T-Q5mBO6Jgb",
                                "name": "Gros titres",
                                "type": "tag",
                                "cover": {
                                    "width": 360,
                                    "height": 110,
                                    "url": "http://6chcdn-res.feednews.com/instaclip/img/B-o56NN7Q1a-B.jpg",
                                    "format": "jpg"
                                },
                                "thumbnail": {
                                    "width": 56,
                                    "height": 56,
                                    "url": "http://6chcdn-res.feednews.com/instaclip/img/B-o56NN7Q1a-A.png",
                                    "format": "png"
                                },
                                "post_count": 15178,
                                "follow_flag": 0
                            }
                        ],
                        "like_count": 0,
                        "dislike_count": 0,
                        "comment_count": 0,
                        "share_count": 0,
                        "liked": false,
                        "disliked": false,
                        "favored": false,
                        "video": {
                            "id": "Q06_CBSsvH77r",
                            "thumbnail": {
                                "width": 1280,
                                "height": 720,
                                "url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_CBSsvH77r.jpg_640x360",
                                "format": "jpg",
                                "preview_url": "http://6chcdn-res.feednews.com/instaclip06/img/Q06_CBSsvH77r.jpg_200x112"
                            },
                            "duration": 41,
                            "size": 1049600,
                            "format": "mp4",
                            "width": 854,
                            "height": 480,
                            "url": "http://video.6chcdn.com/instaclip06/video/Q06_CBSsvH77r.mp4",
                            "download_url": "http://video.6chcdn.com/instaclip06/video/Q06_CBSsvH77r.mp4",
                            "has_audio": true,
                            "source_type": "normal",
                            "rates": {
                                "p144": 787200,
                                "p240": 1049600,
                                "p360": 1574400,
                                "p480": 2099200,
                                "p540": 2624000,
                                "p720": 4198400
                            }
                        },
                        "share_url": "https://instaclips.op-mobile.opera.com/post/P-xXO5NL82L?share=1&country=fr&language=fr",
                        "create_time": 1605224303000,
                        "state": 100,
                        "algo_data": "{\"feed\": \"https://www.youtube.com/user/BFMTV\", \"tags\": [], \"treetags\": []}",
                        "video_update_lock": false,
                        "infra_feedback": "{\"comment_count\":0,\"loc\":\"-\",\"infra_bit_rate\":\"playlist\",\"like_count\":0,\"gender\":\"unknown\",\"create_time\":1605224303000,\"follower_count\":0,\"share_count\":0,\"width\":854,\"video_full_duration\":41,\"dislike_count\":0,\"follow_flag\":0,\"height\":480}",
                        "entrance_type": "CRAWLER"
                    }
                }
            ]
        },
        {
            "news_entry_id": "",
            "news_id": "",
            "type": "cms_topic_tab_news",
            "enable_more_button": false,
            "more_id": "2509",
            "title": "Le confinement, objet de débat",
            "articles": []
        },
        {
            "news_entry_id": "",
            "news_id": "",
            "type": "cms_topic_tab_news",
            "enable_more_button": false,
            "more_id": "2528",
            "title": "Covid, quels sont les chiffres ?",
            "articles": []
        }
    ]
}`


var payload = `{"flags":148,"group_id":"news-4","id":"ED#11094089","popup_auto_cancel":false,"push_title":"Top News 89","text":"","timeout":86400,"title":"89 Mbare Bus terminus tender scandal rocks Harare City Council "}`

func main() {

   res, err := simplejson.NewJson([]byte(payload))
   if err != nil {
       fmt.Println(err.Error())
   }
    var groupIds = []string{
        "news-1", "news-3", "news-4",
    }
    res.Set("group_id", groupIds[rand.Intn(len(groupIds))])
    extraResult, _ := res.Encode()
    fmt.Println(string(extraResult))



    //res, err := simplejson.NewJson([]byte(json_str))
    //
    //if err != nil {
    //    fmt.Printf("%v\n", err)
    //    return
    //}
    //news, err := simplejson.NewJson([]byte(video_str))
    //
    //if err != nil {
    //    fmt.Printf("%v\n", err)
    //    return
    //}
    //
    //newsArr, _ := news.Get("articles").Array()
    //
    ////Get("articles").GetIndex(j).Get("articles").Array()
    //resArr, _ := res.Get("articles").Array()
    //for i := range resArr {
    //    moreID, _ := res.Get(common.ARTICLES).GetIndex(i).Get("more_id").String()
    //    for j := range newsArr {
    //        moreIDNew, _ := news.Get(common.ARTICLES).GetIndex(j).Get("more_id").String()
    //        if moreIDNew == moreID {
    //            resTmp, _ := res.Get(common.ARTICLES).GetIndex(i).Get(common.ARTICLES).Array()
    //            newsTmp, _ := news.Get(common.ARTICLES).GetIndex(j).Get(common.ARTICLES).Array()
    //            for k := range newsTmp {
    //                resTmp, _ = AggrCard(resTmp, newsTmp[k], 1)
    //            }
    //
    //            res.Get(common.ARTICLES).GetIndex(i).Set(common.ARTICLES, resTmp)
    //        }
    //    }
    //}
    //
    //ret , _ := res.Encode()
    //fmt.Println(string(ret))

}

func AggrCard(base []interface{}, candi interface{}, pos int) (aggr []interface{}, offset int) {
    if pos == 0 {
        aggr = append(aggr, candi)
        return append(aggr, base...), pos
    }
    if len(base) <= pos {
        return append(base, candi), len(base)
    }
    aggr = append(aggr, base[:pos]...)
    aggr = append(aggr, candi)
    aggr = append(aggr, base[pos:]...)
    return aggr, pos
}