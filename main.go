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
	diary := NewDiary(env.diaryDatabaseId, client)

	now := time.Now()

	b, err := diary.hasTodaysPage(now)
	if err != nil {
		log.Fatalf("failed to hasTodaysPage: %v", err)
	}
	if b {
		fmt.Println("Today's diary page was already created.")
		return
	}

	if err := diary.createTodaysPage(now); err != nil {
		log.Fatalf("failed to createTodaysPage: %v", err)
	}

	fmt.Println("Today's diary page was created successfully.")
}
