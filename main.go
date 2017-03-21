package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/codingconcepts/ud/model"
)

const (
	root = "http://api.urbandictionary.com/v0/define"
)

func main() {
	args := os.Args[1:]
	resp, err := request(strings.Join(args, " "))
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.List) < 1 {
		return
	}

	example := resp.List[0]
	fmt.Printf("Definition: %s\n", example.Definition)
	fmt.Printf("Example: %s\n", example.Example)
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
