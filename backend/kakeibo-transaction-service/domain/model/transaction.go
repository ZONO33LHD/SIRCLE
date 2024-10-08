package model

import "time"

type Transaction struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	Amount float64 `json:"amount"`
	Type string `json:"type"`
	CategoryId int32 `json:"category_id"`
	CategoryName string `json:"category_name"`
	Date time.Time `json:"date"`
	Description string `json:"description"`
	IsRecurring bool `json:"is_recurring"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IncomeExpenseSummary struct {
    IncomeItems  []*SummaryItem
    ExpenseItems []*SummaryItem
    Balance      float64
}

type SummaryItem struct {
    Title  string
    Amount float64
}
