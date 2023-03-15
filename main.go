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
		log.Fatalf("Failed to load env: %v", err)
	}

	client := NewClient(env.apiToken)
	diaryRepository := NewDiaryRepository(client)
	diaryService := NewDiaryService(diaryRepository)
	diary := NewDiary(env.diaryDatabaseId)

	now := time.Now()

	exists, err := diaryService.ExistsTodaysPage(diary, now)
	if err != nil {
		log.Fatalf("Failed to ExistsTodaysPage: %v", err)
	}
	if exists {
		log.Println("Today's diary page was already created.")
		os.Exit(0)
	}

	if err := diaryService.CreateTodaysPage(diary, now); err != nil {
		log.Fatalf("Failed to createTodaysPage: %v", err)
	}

	log.Println("Today's diary page was created successfully.")
}
