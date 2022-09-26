package main

import (
	"fmt"
	"log"
	"time"
)

func buildQueryDatabaseParams(now time.Time) QueryDatabaseParams {
	return QueryDatabaseParams{
		Filter: QueryDatabaseFilter{
			Property: "Date",
			QueryDatabasePropertyFilter: QueryDatabasePropertyFilter{
				Date: DateFilterCondition{
					Equals: now.Format("2006-01-02"),
				},
			},
		},
	}
}

func buildCreatePageParams(databaseId string, now time.Time) CreatePageParams {
	jaWeekdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	return CreatePageParams{
		Parent: CreatePageParent{
			Type:       "database_id",
			DatabaseId: databaseId,
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
}

func main() {
	fmt.Println("Creating diary page started.")

	env := GetEnv()
	client := NewClient(env.apiToken)

	now := time.Now()

	qdParams := buildQueryDatabaseParams(now)
	res, err := client.QueryDatabase(env.diaryDatabaseId, qdParams)
	if err != nil {
		log.Fatalf("failed to queryDatabase: %v", err)
	}

	if len(res.Results) > 0 {
		fmt.Println("Today's diary page was already created.")
		return
	}

	cpParams := buildCreatePageParams(env.diaryDatabaseId, now)
	if err := client.CreatePage(cpParams); err != nil {
		log.Fatalf("failed to createPage: %v", err)
	}

	fmt.Println("Today's diary page was created successfully.")
}
