package main

import "fmt"

type DiaryRepository struct {
	client *Client
}

func NewDiaryRepository(client *Client) *DiaryRepository {
	return &DiaryRepository{client}
}

func (r *DiaryRepository) FindPagesByDate(id string, date string) ([]Page, error) {
	params := QueryDatabaseParams{
		Filter: QueryDatabaseFilter{
			Property: "Date",
			QueryDatabasePropertyFilter: QueryDatabasePropertyFilter{
				Date: DateFilterCondition{
					Equals: date,
				},
			},
		},
	}

	res, err := r.client.QueryDatabase(id, params)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	var pages []Page
	for _, result := range res.Results {
		pages = append(pages, Page{id: result.Id})
	}

	return pages, nil
}

func (r *DiaryRepository) CreatePage(id string, title string, date string) error {
	params := CreatePageParams{
		Parent: CreatePageParent{
			Type:       "database_id",
			DatabaseId: id,
		},
		Properties: CreatePageProperties{
			"Name": {
				Title: []TitleProperty{
					{
						Text: TextProperty{
							Content: title,
						},
					},
				},
			},
			"Date": {
				Date: &DateProperty{
					Start: date,
				},
			},
		},
	}

	err := r.client.CreatePage(params)
	if err != nil {
		return fmt.Errorf("failed to create page: %w", err)
	}

	return nil
}
