package client

import (
	"bytes"
	"encoding/json"
	"io"
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

	// read in body for JSON unmarshalling (using
	// unmarshal instead of decode because we receive
	// a complete JSON object, not a JSON stream)
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, resp.Body); err != nil {
		return
	}

	answer = new(model.Response)
	err = json.Unmarshal(buf.Bytes(), &answer)

	return
}
