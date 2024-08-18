package model

import "time"

type User struct {
	ID                      int                      `json:"id"`
	Name                    string                   `json:"name"`
	Email                   string                   `json:"email"`
	PasswordHash            []byte                   `json:"password_hash"`
	NotificationPreferences *NotificationPreferences `json:"notification_preferences"`
	CreatedAt               time.Time                `json:"created_at"`
	UpdatedAt               time.Time                `json:"updated_at"`
}

type Budget struct {
	ID                string
	UserID            string
	Name              string
	Amount            int64
	Category          string
	StartDate         time.Time
	EndDate           time.Time
	ReminderFrequency ReminderFrequency
}

type Goal struct {
	ID                string
	UserID            string
	Name              string
	TargetAmount      float64
	CurrentAmount     float64
	TargetDate        time.Time
	ReminderFrequency ReminderFrequency
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
