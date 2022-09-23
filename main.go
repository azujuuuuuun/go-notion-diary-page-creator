package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	apiToken        string
	diaryDatabaseId string
}

func GetEnv() Env {
	return Env{
		apiToken:        os.Getenv("NOTION_API_TOKEN"),
		diaryDatabaseId: os.Getenv("NOTION_DIARY_DATABASE_ID"),
	}
}

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
	filter     map[string]interface{}
}

type QueryDatabaseResponse struct {
	Results []struct {
		Id string `json:"id"`
	} `json:"results"`
}

func (c *Client) QueryDatabase(params QueryDatabaseParams) (*QueryDatabaseResponse, error) {
	url := "https://api.notion.com/v1/databases/" + params.databaseId + "/query"

	// TODO: Handle nil
	b, err := json.Marshal(params.filter)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader("{\"filter\":" + string(b) + "}")

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

func main() {
	fmt.Println("Creating diary page started.")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := GetEnv()
	client := NewClient(env.apiToken)

	now := time.Now()

	params := QueryDatabaseParams{
		databaseId: env.diaryDatabaseId,
		filter: map[string]interface{}{
			"property": "Date",
			"date": map[string]interface{}{
				"equals": now.Format("2006-01-02"),
			},
		},
	}
	res, err := client.QueryDatabase(params)
	if err != nil {
		log.Fatalf("failed to queryDatabase: %v", err)
	}

	if len(res.Results) > 0 {
		fmt.Println("Today's diary page was already created.")
		return
	}

	// TODO: create page
	// fmt.Println("Today's diary page was created successfully.")
}
