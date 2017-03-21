package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
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

	// if no hits, bum out
	if len(resp.List) < 1 {
		fmt.Println("No results")
		return
	}

	// if multiple results, sort them by rank descending
	// to ensure the best result is the one we display
	if len(resp.List) > 1 {
		sort.Slice(resp.List, func(i int, j int) bool {
			return resp.List[i].ThumbsUp > resp.List[j].ThumbsUp
		})
	}

	example := resp.List[0]
	if example.Definition != "" {
		fmt.Printf("Definition: %s\n", example.Definition)
	}
	if example.Example != "" {
		fmt.Printf("Example: %s\n", example.Example)
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

	answer = new(model.Response)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&answer)

	return
}
