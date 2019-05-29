package model

import "biloba/helpers/timestamp"

type Kid struct {
	ID					string
	Name				string					`json:"name"`
	BirthDate			*timestamp.Timestamp	`json:"birthDate"`
	IsOnboardingDone 	bool					`json:"isOnboardingDone"`
	Gender				int						`json:"gender"`
	CalendarId			int						`json:"calendarId"`
	LateDate			*timestamp.Timestamp	`json:"lateDate"`
	City				City					`json:"city"`
	Size 				int						`json:"size"`
	Weight				int						`json:"weight"`
}
