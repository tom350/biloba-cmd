package model

import "biloba-cmd/helpers/timestamp"

type Kid struct {
	ID					string
	Name				string
	BirthDate			*timestamp.Timestamp
	IsOnboardingDone 	bool
	Gender				int
	CalendarId			int
	LateDate			*timestamp.Timestamp
	City				City
	Size 				int
	Weight				int
}

type City struct {
	Code	string
	Name	string
	ZipCode	string
	ZipName	string
}

type Kids []Kid
