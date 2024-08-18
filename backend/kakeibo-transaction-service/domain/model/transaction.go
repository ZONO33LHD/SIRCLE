package model

import "time"

type Transaction struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	Amount float64 `json:"amount"`
	Type string `json:"type"`
	CategoryId string `json:"category_id"`
	Date time.Time `json:"date"`
	Description string `json:"description"`
	IsRecurring bool `json:"is_recurring"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}