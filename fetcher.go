package fetcher

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var Fetcher fetcher

type fetcher struct {
	Endpoint string
	Token    string
}

func CreateFetcher(endpoint string) *fetcher {
	return &fetcher{Endpoint: endpoint}
}
func (f *fetcher) Request(method, path string, body map[string]string) (*http.Response, error) {
	var data io.Reader = nil
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		data = bytes.NewBuffer(jsonData)
	}
	req, err := http.NewRequest(method, f.Endpoint+path, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "applicaton/json")
	if f.Token != "" {
		req.Header.Add("Authorization", "Bearer "+f.Token)
	}
	client := &http.Client{}
	return client.Do(req)
}
