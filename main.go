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
		Filter: QueryDatabaseFilter{
			Property: "Date",
			QueryDatabasePropertyFilter: QueryDatabasePropertyFilter{
				Date: DateFilterCondition{
					Equals: now.Format("2006-01-02"),
				},
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
		Parent: CreatePageParent{
			Type:       "database_id",
			DatabaseId: env.diaryDatabaseId,
		},
		Properties: map[string]CreatePageProperty{
			"Name": {
				Title: []TitleProperty{
					{
						Text: TextProperty{
							Content: now.Format("2006/01/02") + "(" + jaWeekdays[now.Weekday()] + ")",
						},
					},
				},
			},
			"Date": {
				Date: &DateProperty{
					Start: now.Format("2006-01-02"),
				},
			},
		},
	}
	if err := client.CreatePage(cpParams); err != nil {
		log.Fatalf("failed to createPage: %v", err)
	}

	fmt.Println("Today's diary page was created successfully.")
}
