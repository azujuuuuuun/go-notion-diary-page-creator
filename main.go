package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Creating diary page started.")

	env, err := GetEnv()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}
	client := NewClient(env.apiToken)
	diary := NewDiary(env.diaryDatabaseId, client)

	now := time.Now()

	b, err := diary.hasTodaysPage(now)
	if err != nil {
		log.Fatalf("failed to hasTodaysPage: %v", err)
	}
	if b {
		log.Println("Today's diary page was already created.")
		os.Exit(0)
	}

	if err := diary.createTodaysPage(now); err != nil {
		log.Fatalf("failed to createTodaysPage: %v", err)
	}

	log.Println("Today's diary page was created successfully.")
}
