package main

import (
	"testing"
	"time"
)

func TestGetJapaneseDayOfWeek(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want string
	}{{
		name: "Sunday",
		now:  time.Date(2023, time.March, 12, 9, 0, 0, 0, time.UTC),
		want: "日",
	}, {
		name: "Monday",
		now:  time.Date(2023, time.March, 13, 9, 0, 0, 0, time.UTC),
		want: "月",
	}, {
		name: "Tuesday",
		now:  time.Date(2023, time.March, 14, 9, 0, 0, 0, time.UTC),
		want: "火",
	}, {
		name: "Wednesday",
		now:  time.Date(2023, time.March, 15, 9, 0, 0, 0, time.UTC),
		want: "水",
	}, {
		name: "Thursday",
		now:  time.Date(2023, time.March, 16, 9, 0, 0, 0, time.UTC),
		want: "木",
	}, {
		name: "Friday",
		now:  time.Date(2023, time.March, 17, 9, 0, 0, 0, time.UTC),
		want: "金",
	}, {
		name: "Saturday",
		now:  time.Date(2023, time.March, 18, 9, 0, 0, 0, time.UTC),
		want: "土",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if GetJapaneseDayOfWeek(tt.now) != tt.want {
				t.Errorf("now = %v, want = %v", tt.now, tt.want)
			}
		})
	}
}
