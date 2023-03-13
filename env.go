package main

import (
	"fmt"
	"os"
	"strings"
)

type Env struct {
	apiToken        string
	diaryDatabaseId string
}

func GetEnv() (Env, error) {
	var env Env
	var missing []string

	for k, v := range map[string]*string{
		"NOTION_API_TOKEN":         &env.apiToken,
		"NOTION_DIARY_DATABASE_ID": &env.diaryDatabaseId,
	} {
		*v = os.Getenv(k)

		if *v == "" {
			missing = append(missing, k)
		}
	}

	if len(missing) > 0 {
		return env, fmt.Errorf("missing env(s): " + strings.Join(missing, ", "))
	}

	return env, nil
}
