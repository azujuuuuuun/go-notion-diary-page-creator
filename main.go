package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Creating diary page started.")

	env := GetEnv()
	client := NewClient(env.apiToken)

	now := time.Now()

	params := QueryDatabaseParams{
		filter: map[string]interface{}{
			"property": "Date",
			"date": map[string]interface{}{
				"equals": now.Format("2006-01-02"),
			},
		},
	}
	res, err := client.QueryDatabase(env.diaryDatabaseId, params)
	if err != nil {
		log.Fatalf("failed to queryDatabase: %v", err)
	}

	if len(res.Results) > 0 {
		fmt.Println("Today's diary page was already created.")
		return
	}

	jaWeekdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	cpParams := CreatePageParams{
		parent: map[string]interface{}{
			"type":        "database_id",
			"database_id": env.diaryDatabaseId,
		},
		properties: map[string]interface{}{
			"Name": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"text": map[string]interface{}{
							"content": now.Format("2006/01/02") + "(" + jaWeekdays[now.Weekday()] + ")",
						},
					},
				},
			},
			"Date": map[string]interface{}{
				"date": map[string]interface{}{
					"start": now.Format("2006-01-02"),
				},
			},
		},
	}
	if err := client.CreatePage(cpParams); err != nil {
		log.Fatalf("failed to createPage: %v", err)
	}

	fmt.Println("Today's diary page was created successfully.")
}
