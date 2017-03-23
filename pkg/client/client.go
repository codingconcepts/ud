package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/codingconcepts/ud/pkg/model"
)

const (
	root = "http://api.urbandictionary.com/v0/define"
)

// Request performs the Urban Dictionary request and returns
// a Response object with the search hits.
func Request(term string) (answer *model.Response, err error) {
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
