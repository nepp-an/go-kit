package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type ABTestExperiment struct {
	Model        string                 `json:"module"`
	ExperimentId string                 `json:"eid,omitempty"`
	VersionId    string                 `json:"vid,omitempty"`
	Parameters   map[string]interface{} `json:"parameters"`
}
type ABTest struct {
	parameter string
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("host is ", r.Host)
	ABTestParament := r.Form.Get("ABTestParament")
	var abTest []ABTestExperiment
	err := json.Unmarshal([]byte(ABTestParament), &abTest)
	if err != nil {
		fmt.Println("error=", err)
	}
	t, err := template.ParseFiles("./test_html/abtest.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, map[string]interface{}{"parameter": ABTestParament}); err != nil {
		w.Write([]byte(fmt.Sprintf("There is an error, error is %s", err.Error())))
		return
	}
	fmt.Println(string(buffer.Bytes()))
	w.Write(buffer.Bytes())

}

func main() {
	//http.HandleFunc("/html/pics/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, r.URL.Path[1:])
	//})
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9997", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
