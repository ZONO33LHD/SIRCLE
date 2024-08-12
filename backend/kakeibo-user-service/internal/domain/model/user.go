package model

import "time"

type User struct {
	ID                      string
	Name                    string
	Email                   string
	Birthday                time.Time
	Budgets                 []Budget
	Goals                   []Goal
	NotificationPreferences NotificationPreferences
}

type Budget struct {
	ID                 string
	UserID             string
	Name               string
	Amount             int64
	Category           string
	StartDate          time.Time
	EndDate            time.Time
	ReminderFrequency  ReminderFrequency
}

type Goal struct {
	ID                 string
	UserID             string
	Name               string
	TargetAmount       float64
	CurrentAmount      float64
	TargetDate         time.Time
	ReminderFrequency  ReminderFrequency
}

type NotificationPreferences struct {
	BudgetNotifications bool
	GoalNotifications   bool
}

type ReminderFrequency int

const (
	ReminderDaily ReminderFrequency = iota
	ReminderWeekly
	ReminderMonthly
)
