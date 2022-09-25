package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
	}
}

type QueryDatabaseParams struct {
	databaseId string
	filter     interface{}
}

type QueryDatabaseResponse struct {
	Results []struct {
		Id string `json:"id"`
	} `json:"results"`
}

func (c *Client) QueryDatabase(params QueryDatabaseParams) (*QueryDatabaseResponse, error) {
	url := "https://api.notion.com/v1/databases/" + params.databaseId + "/query"

	bodyParams := struct {
		Filter interface{} `json:"filter,omitempty"`
	}{
		Filter: params.filter,
	}
	b, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(b))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.apiToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil && err != io.EOF {
		return nil, err
	}

	var resp QueryDatabaseResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type CreatePageParams struct {
	parent     interface{}
	properties interface{}
}

func (c *Client) CreatePage(params CreatePageParams) error {
	url := "https://api.notion.com/v1/pages"

	bodyParams := struct {
		Parent     interface{} `json:"parent,omitempty"`
		Properties interface{} `json:"properties,omitempty"`
	}{
		Parent:     params.parent,
		Properties: params.properties,
	}
	b, err := json.Marshal(bodyParams)
	if err != nil {
		return err
	}

	payload := strings.NewReader(string(b))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.apiToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
