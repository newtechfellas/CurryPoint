package entity

import "time"

var ScheduleDayItem struct {
	Id         string
	ScheduleId string
	Date       time.Time
	MenuItemId string
	Count      int
}
