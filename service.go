package main

import (
	"fmt"
	"time"
)

type DiaryService struct {
	diaryRepository *DiaryRepository
}

func NewDiaryService(diaryRepository *DiaryRepository) *DiaryService {
	return &DiaryService{diaryRepository}
}

func (s *DiaryService) ExistsTodaysPage(diary Diary, now time.Time) (bool, error) {
	date := now.Format("2006-01-02")

	pages, err := s.diaryRepository.FindPagesByDate(diary.id, date)
	if err != nil {
		return false, fmt.Errorf("failed to find pages by date: %w", err)
	}

	if len(pages) > 0 {
		return true, nil
	}

	return false, nil
}

func (s *DiaryService) CreateTodaysPage(diary Diary, now time.Time) error {
	dayOfWeek := GetJapaneseDayOfWeek(now)
	title := now.Format("2006/01/02") + "(" + dayOfWeek + ")"
	date := now.Format("2006-01-02")

	return s.diaryRepository.CreatePage(diary.id, title, date)
}
