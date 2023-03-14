package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
	}
}

type DateFilterCondition struct {
	Equals string `json:"equals,omitempty"`
}

type QueryDatabasePropertyFilter struct {
	Date DateFilterCondition `json:"date,omitempty"`
}

type QueryDatabaseFilter struct {
	Property string `json:"property,omitempty"`
	QueryDatabasePropertyFilter
}

type QueryDatabaseParams struct {
	Filter QueryDatabaseFilter `json:"filter,omitempty"`
}

type QueryDatabaseResult struct {
	Id string `json:"id"`
}

type QueryDatabaseResponse struct {
	Results []QueryDatabaseResult `json:"results"`
}

func (c *Client) QueryDatabase(id string, params QueryDatabaseParams) (*QueryDatabaseResponse, error) {
	url := "https://api.notion.com/v1/databases/" + id + "/query"

	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
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

type CreatePageParent struct {
	Type       string `json:"type"`
	DatabaseId string `json:"database_id,omitempty"`
}

type TextProperty struct {
	Content string `json:"content,omitempty"`
}

type TitleProperty struct {
	Text TextProperty `json:"text,omitempty"`
}

type DateProperty struct {
	Start string `json:"start,omitempty"`
}

type CreatePageProperty struct {
	Title []TitleProperty `json:"title,omitempty"`
	Date  *DateProperty   `json:"date,omitempty"`
}

type CreatePageProperties map[string]CreatePageProperty

type CreatePageParams struct {
	Parent     CreatePageParent     `json:"parent"`
	Properties CreatePageProperties `json:"properties"`
}

func (c *Client) CreatePage(params CreatePageParams) error {
	url := "https://api.notion.com/v1/pages"

	b, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
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
