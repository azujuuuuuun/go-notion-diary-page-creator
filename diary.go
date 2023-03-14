package main

import (
	"fmt"
	"time"
)

type Diary struct {
	databaseId string
	client     *Client
}

func NewDiary(databaseId string, client *Client) *Diary {
	return &Diary{databaseId: databaseId, client: client}
}

func (d *Diary) hasTodaysPage(now time.Time) (bool, error) {
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

	res, err := d.client.QueryDatabase(d.databaseId, params)
	if err != nil {
		return false, fmt.Errorf("failed to QueryDatabase: %v", err)
	}

	if len(res.Results) > 0 {
		return true, nil
	}

	return false, nil
}

func (d *Diary) createTodaysPage(now time.Time) error {
	jaWeekdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	params := CreatePageParams{
		Parent: CreatePageParent{
			Type:       "database_id",
			DatabaseId: d.databaseId,
		},
		Properties: CreatePageProperties{
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

	err := d.client.CreatePage(params)
	if err != nil {
		return fmt.Errorf("failed to CreatePage: %v", err)
	}

	return nil
}
