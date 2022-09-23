package main

import "os"

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
