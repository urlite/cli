package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UrlResult struct {
	ShortUrl  string `json:"shortUrl"`
	TotalUrls int    `json:"totalUrls"`
}

func GetShortUrl(url string) (*UrlResult, error) {
	jsonStr := []byte(fmt.Sprintf(`{"url":"%s"}`, url))
	req, err := http.NewRequest("POST", "https://urlite.cc", bytes.NewBuffer(jsonStr))
	req.Header.Set("x-api-key", "")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, nil
	}
	var objectResult UrlResult
	err = json.Unmarshal(body, &objectResult)
	if err != nil {
		return nil, err
	}
	return &objectResult, nil
}
