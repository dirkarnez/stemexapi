package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Post[K, T any](httpClient *http.Client, header http.Header, url string, payload *K) (*T, error) {
	jsonValue, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header = header

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var a T
	err = json.Unmarshal(bytes, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
