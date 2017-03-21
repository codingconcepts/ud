package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/codingconcepts/ud/model"
)

const (
	root = "http://api.urbandictionary.com/v0/define"
)

func main() {
	resp, err := request("blah")
	if err != nil {
		log.Fatal(err)
	}

	for _, example := range resp.List {
		fmt.Println(example)
	}
}

func request(term string) (answer *model.Response, err error) {
	var request *http.Request
	if request, err = http.NewRequest(http.MethodGet, root, nil); err != nil {
		return
	}

	query := request.URL.Query()
	query.Add("term", term)
	request.URL.RawQuery = query.Encode()

	client := http.Client{
		Timeout: time.Second,
	}

	var resp *http.Response
	if resp, err = client.Do(request); err != nil {
		return
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	resp.Body.Close()

	answer = new(model.Response)
	err = json.Unmarshal(body, &answer)

	return
}
