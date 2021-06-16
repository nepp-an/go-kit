package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)


type shortUrl struct {
	Domain string `json:"short_domain"`
}

func GetShortUrl(url string) string {
	commentUrl := `http://nginx-web.feednews.com:8880/form_api`
	postData := make(map[string]string)
	postData["link"] = url
	postData["content"] = "test1"
	postData["group_id"] = strconv.Itoa(3)
	postData["ttl"] = strconv.Itoa(3110400000)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range postData {
		w.WriteField(k, v)
	}
	w.Close()
	req, _ := http.NewRequest("POST", commentUrl, body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Add("X-Email", "puqianga@opera.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err is", err)
		return ""
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	var retUrl shortUrl
	if err = json.Unmarshal(data, &retUrl); err != nil {
		fmt.Println("unmarshal failed err is", err)
		return ""
	}
	ret := "http://s.opr.news" + "/" + retUrl.Domain

	return ret
}

func main() {
	var input = "123#90923"
	if index := strings.Index(input, "#"); index != -1 {
		typePush := input[:index]
		fmt.Println(typePush)
	}
	fmt.Println(time.Now().Unix())

	testIndex := "9917000000"
	if testIndex[:2] == "99" {
		if testIndex[2] == '0' {
			testIndex = testIndex[3:4] + "0000"
		} else {
			testIndex = testIndex[2:4] + "0000"
		}
	}
	fmt.Println(testIndex)



	test := `"Concieved in the state of Lagos on 11 August 1982. Anambra State Queen Nwokoye Hails, Nigeria. She's a creator and actress from Nollywood.
Queen's Air Force Education had her elementary school and her highschool in the state of Enugu. Queen College.
Queen Nwokoye has a degree in Sociology and Anthropology from the University of Nnamdi, Azikiwe.
She began her acting career in 2004, when she was cast in a film called \"Nna Guys,\" she became a part of Limelight.
She starred in more than 180 Nollywood Movies such as Last Kobo, Expanding the Universe, My Girlfriend, Her Beauty, Protection Risk, The Girl is Mine, Save the baby, Lady of Faith, Last Supper, End of Mirror, Jealous Buddy and the Gentleman's League.
Queen collaborated with several celebrity parties, such as Samuel Oliva, Ify Afuba, Ifeanyi Ike, Frank Odeke, Henry Isaac etc. Queen works with Samouel Bellyn, Ifeanyi Ezeonu and Bellar.
Queen Nwokoye has received Received Several Prizes, Including best actress in a leading role in the City Citizens Entertainment Prizes, the best indigenous female actress and Face of Nollywood.
Happily married with one lovely daughter is Queen Nwokoye.
Their marriage
Her adorable Daughter
"`
	fmt.Println(test[:300])

	test="20196"
	test = test[:(len(test)-4)] + "0000"
	fmt.Println("cityINdex is "+ test)


	splitStr := `<p>In my article today, I have selected some funny memes that I want to share with you. But before I proceed with the memes, I have made known to you that, my contents either educate or entertain, or both. Please can we learn something small concerning laughter?</p>
<p>Let's look at the word \"laugh\".</p>
<p>Laugh as a verb: to make the spontaneous sounds and movements of the face and body that are the instinctive expressions of lively amusement and sometimes also of derision.</p>
<p>Synonyms: guffaw, chuckle, chortle, cackle, howl, roar, ha-ha, fall about, roar/hoot with laughter, shake with laughter, be convulsed with laughter, dissolve into laughter, split one's sides, be doubled up, be in stitches, die laughing, be rolling in the aisles, laugh like a drain, bust a gut, break up, be creased up, crease up, crack up, giggle, titter, snigger, snicker, tee-hee, burst out laughing, hold one's sides, ridicule, mock, deride, scoff at, jeer at, sneer at, jibe at, make fun of, poke fun at, make jokes about, heap scorn on, scorn, pooh-pooh, lampoon, satirize, caricature, parody, taunt, tease, torment, send up, take the mickey out of, poke mullock at, make sport of, take the piss out of.</p>
<p>Now let's get today's 30 funny pictures which I have posted them here.</p>
<p><br></p>
<p>1.<img src=\"http://res.feednews.com/assets/v2/eff6c0da9e47ccc04ba60fc3cff4ddce?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>2.<img src=\"http://res.feednews.com/assets/v2/1faec5e601ad5cd550bab7f5df4c9672?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>3.<img src=\"http://res.feednews.com/assets/v2/911dff01a0634c74d24efc2253cbe57d?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>4.<img src=\"http://res.feednews.com/assets/v2/60229167aeb3b542bb1d3e719493f073?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>5.<img src=\"http://res.feednews.com/assets/v2/354342bf7a0de4325b4a4c561c851800?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>6.<img src=\"http://res.feednews.com/assets/v2/0eccefdda45bb43e5d1b797bb60f62ca?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<h2>7. Wicked.</h2>
<p><br></p>
<p><img src=\"http://res.feednews.com/assets/v2/456b60b3ae20bae756390aeb0c995d6b?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<h2>8. Adam?</h2>
<p><br></p>
<p><img src=\"http://res.feednews.com/assets/v2/15acd605078ff7d4f60ff7d7a7cec22e?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>9.<img src=\"http://res.feednews.com/assets/v2/9756894e87ffa7617c732f0582d1f7ad?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>10.</p>
<p><img src=\"http://res.feednews.com/assets/v2/ebc510706cd67ab6051ef4be83ee1621?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>11.</p>
<p><img src=\"http://res.feednews.com/assets/v2/61f73ca994426412155c8eacfad2c01f?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>12.</p>
<p><img src=\"http://res.feednews.com/assets/v2/9e17cf6c54b74c4d2715105655560e1d?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>13.</p>
<p><img src=\"http://res.feednews.com/assets/v2/ed32c88c50b9502cf85fbf20941630fd?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>14.</p>
<p><img src=\"http://res.feednews.com/assets/v2/5c7eb2e514e2dad1f904da64071dd4d9?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>15.</p>
<p><img src=\"http://res.feednews.com/assets/v2/d439a69cbd14cd2bd985b52d941ba5f8?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>16.</p>
<p><img src=\"http://res.feednews.com/assets/v2/afa17252fbb33f53e386774164462e93?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<h2>17. One word for the man</h2>
<p><br></p>
<p><img src=\"http://res.feednews.com/assets/v2/12c6f88c8b66911c0e7d0e30f368186c?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>18.</p>
<p><img src=\"http://res.feednews.com/assets/v2/be2fc41652702f62d64b178565ce1057?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>19.</p>
<p><img src=\"http://res.feednews.com/assets/v2/effb61f54264d73047462ee31652586a?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<h2>20. No.</h2>
<p><br></p>
<p><img src=\"http://res.feednews.com/assets/v2/04c0402d979b26ca300e4880891b72dd?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>21.</p>
<p><img src=\"http://res.feednews.com/assets/v2/6a4377d0a29b51e400e34164e790c491?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>22.</p>
<p><img src=\"http://res.feednews.com/assets/v2/a957fd244167098c635cfced65b72660?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>23.</p>
<p><img src=\"http://res.feednews.com/assets/v2/2b4d0f2e98a5bdbea3c676432a39e324?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>24.</p>
<p><img src=\"http://res.feednews.com/assets/v2/07bf4b9cc81b31c30ba8b78a7c5d1923?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>25.</p>
<p><img src=\"http://res.feednews.com/assets/v2/b2a379ddb3ba4f51f38f0764dd86ec4b?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>26.</p>
<p><img src=\"http://res.feednews.com/assets/v2/2e44848c592e36d37f9e7f935a7ac74c?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<h2>27. I and my brother are one.</h2>
<p><br></p>
<p><img src=\"http://res.feednews.com/assets/v2/98ad21666623bf9feeb4f6ea6c5e0aef?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>28.</p>
<p><img src=\"http://res.feednews.com/assets/v2/bbffce2be465f3e70811b6d0a358523f?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>29.</p>
<p><img src=\"http://res.feednews.com/assets/v2/76daecc59099749fb7ac6186560eb735?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<h2>30. School girls.</h2>
<p><br></p>
<p><img src=\"http://res.feednews.com/assets/v2/1374e961b9533d42ad903530ba99dd55?quality={opera_quality}&amp;format={opera_format}\"></p>
<p><br></p>
<p>Leave your comments below.</p>
<p>What exactly made your day?</p>
<p><br></p>
<p>Download Opera News app, log in and follow me for more memes.</p>
<p>Don't forget to like and share.</p>`
	regTopic := regexp.MustCompile("http://res.feednews.com/assets/v2/[a-zA-Z0-9]+/?")
	tmp:=regTopic.FindAllString(splitStr, -1)
	fmt.Println(tmp[0])
	var ret string = tmp[0]
	fmt.Println(ret[strings.Index(tmp[0], "v2/")+3:])

	str := "99f169e69086868c9bb533f2d1afa790?quality=uhq&width=200&height=150"
	fmt.Println("last")
	fmt.Println(str[:strings.Index(str, "?")])


	testUrl := "http://res.feednews.com/assets/v2/9115e8ec99beb0a7f4bb0b58d751dab2"
	fmt.Println(strings.Replace(testUrl, "http://", "https://", 1))

	fmt.Println(10/1000)

	fmt.Println(strings.Split("123", "||||"))


	fmt.Println("//////////////////////////////////")
	fmt.Println(strings.Split(`106|160|240|320|400|480|504|720|960`, "|"))
	var EntryID = "s12345"
	fmt.Println(EntryID[0:1]=="s")

	// key: image/default_images/IN_Crime/IN_Crime 00014.jpg
	keys := []string{
		`image/default_images/IN_Crime/IN_Crime 00014.jpg`,
		`image/default_images/IN_Crime/IN_Crime 00015.jpg`,
		`image/default_images/IN_Crime/IN_Crime 00016.jpg`,
		`image/default_images/IN_Crime/IN_Crime 00017.jpg`,
		`image/default_images/IN_Crime/IN_Crime 00018.jpg`,
		`image/default_images/IN_Crime_1/IN_Crime 00019.jpg`,
		`image/default_images/IN_Crime_1/IN_Crime 00010.jpg`,
		`image/default_images/IN_Crime_1/IN_Crime 00013.jpg`,
		`image/default_images/IN_Crime_1/IN_Crime 00012.jpg`,
		`image/default_images/IN_Crime_1/IN_Crime 00011.jpg`,
	}
	LoadDefaultImages(keys)
	fmt.Println(defaultImageMap)

	testMap := make(map[string]string)
	testMap["a"] = "a"
	fmt.Println(len(testMap["b"]))

	str = "/assets/d1/d2/d3"
	tmp = strings.Split(str, "/")
	//index := strings.LastIndex(str, "/")
	fmt.Println(str[:0])

	a1 := `image/default_images/IN_Crime/IN_Crime 000999.jpg`
	a2 := strings.ReplaceAll(a1, "/", "%2F")
	fmt.Println(a2)
	a3 := "/assets/v2/9ab6ea354002294e828874db070fc4bc?width=225&height=168&quality=hq&category=NG_News_Entertainment&default=true&format=webp"
	a3 = strings.ReplaceAll(a3, "&default=true", "")
	fmt.Println(a3)
	fmt.Println(getRealImageID("1234.webp"))
	fmt.Println(getRealImageID("1234.."))
	fmt.Println(getRealImageID("1234"))

	fmt.Println(removeIDFromPath(`/assets/v1/crop/xxl/c/lq/s/602f1242388855ac54129ec56a67bd1a.webp`))
	fmt.Println(removeIDFromPath(`/assets/v1/crop/xl/c/mq/s/31919fca50cd3f46837d752372c7fbd9-1`))

	fmt.Println(versionCompare("8.3.1.1"))
	fmt.Println(versionCompare("8.4.1.1"))
	fmt.Println(versionCompare("7.9.1.1"))

	str = "tom and \"jerry\""
	fmt.Println(str)
	format := "//"
	if index := strings.Index(format, "/"); index >=0  {
		format = format[index+1:]
	}
	fmt.Println(format)
	var cl = "eg_ar"
	var language string
	if tmp := strings.Split(cl, "_"); len(tmp) == 2{
		language = tmp[1]
	}
	fmt.Println(language)

	splitStr = "_"
	fmt.Println(strings.Split(splitStr, "_"))
}


func versionCompare(version string) bool {
	if version > "8.3" {
		return true
	}
	return false
}

func removeIDFromPath(path string) string{
	if index := strings.LastIndex(path, "/"); index != -1 {
		return path[:(index+1)]
	}
	return path
}
func getRealImageID(id string) string {
	if index := strings.Index(id, "."); index != -1 {
		return id[:index]
	}
	return id
}

var defaultImageMap = make(map[string][]string)
func LoadDefaultImages( keys []string) {
	//list from Upload-Images bucket

	//save in a map
	// key: image/default_images/IN_Crime/IN_Crime 00014.jpg
	//依赖：key必须是顺序遍历的
	//依赖：key格式固定
	var imageArr []string
	var category string
	for _, key := range keys {
		tmps := strings.Split(key, "/")
		fmt.Println(tmps)
		if len(tmps) != 4 {
			continue
		}
		if category == "" {
			category = tmps[2]
			imageArr = []string{tmps[3]}
		} else if category == tmps[2] {
			imageArr = append(imageArr, tmps[3])
		} else {
			defaultImageMap[category] = imageArr
			imageArr = []string{tmps[3]}
			category = tmps[2]
		}
	}
	defaultImageMap[category] = imageArr

}

var langMap = map[string]bool{
	"en": true,
	"fr": true,
	"sw": true,
	"ar": true,
	"pt": true,
	"am": true,
	"ha": true,
	"zu": true,
	"ru": true,
}