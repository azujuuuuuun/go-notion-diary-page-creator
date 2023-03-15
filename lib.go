package main

import "time"

func GetJapaneseDayOfWeek(now time.Time) string {
	jaWeekdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	return jaWeekdays[now.Weekday()]
}
