package entity

import "time"

type ScheduleOrder struct {
	Id string
	CustomerEmail string
	IsActive bool
	StartDate time.Time //End date will not be required. This will always be a Week order
}