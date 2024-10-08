// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type Budget struct {
	ID       string    `json:"id"`
	UserID   string    `json:"userId"`
	Amount   float64   `json:"amount"`
	Category *Category `json:"category"`
	Period   Period    `json:"period"`
}

type Category struct {
	ID   string          `json:"id"`
	Name string          `json:"name"`
	Type TransactionType `json:"type"`
}

type CategoryBreakdown struct {
	Category   *Category `json:"category"`
	Amount     float64   `json:"amount"`
	Percentage float64   `json:"percentage"`
}

type CreateBudgetInput struct {
	UserID     string  `json:"userId"`
	Amount     float64 `json:"amount"`
	CategoryID int     `json:"categoryId"`
	Period     Period  `json:"period"`
}

type CreateCategoryInput struct {
	Name string          `json:"name"`
	Type TransactionType `json:"type"`
}

type CreateTransactionInput struct {
	UserID             string              `json:"userId"`
	Amount             float64             `json:"amount"`
	Type               TransactionType     `json:"type"`
	CategoryID         int                 `json:"categoryId"`
	Date               string              `json:"date"`
	Description        *string             `json:"description,omitempty"`
	IsRecurring        bool                `json:"isRecurring"`
	RecurringFrequency *RecurringFrequency `json:"recurringFrequency,omitempty"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ExpenseItem struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

type Goal struct {
	ID            string  `json:"id"`
	UserID        string  `json:"userId"`
	Name          string  `json:"name"`
	TargetAmount  float64 `json:"targetAmount"`
	CurrentAmount float64 `json:"currentAmount"`
	Deadline      *string `json:"deadline,omitempty"`
}

type IncomeExpenseSummary struct {
	IncomeItems  []*IncomeItem  `json:"incomeItems"`
	ExpenseItems []*ExpenseItem `json:"expenseItems"`
	Balance      float64        `json:"balance"`
}

type IncomeItem struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

type Mutation struct {
}

type NotificationPreferences struct {
	BudgetAlerts      bool              `json:"budgetAlerts"`
	ReminderFrequency ReminderFrequency `json:"reminderFrequency"`
}

type NotificationPreferencesInput struct {
	BudgetAlerts      bool              `json:"budgetAlerts"`
	ReminderFrequency ReminderFrequency `json:"reminderFrequency"`
}

type Query struct {
}

type Report struct {
	TotalIncome       float64              `json:"totalIncome"`
	TotalExpense      float64              `json:"totalExpense"`
	Balance           float64              `json:"balance"`
	CategoryBreakdown []*CategoryBreakdown `json:"categoryBreakdown"`
	Trends            []*TrendPoint        `json:"trends"`
}

type SetGoalInput struct {
	UserID       string  `json:"userId"`
	Name         string  `json:"name"`
	TargetAmount float64 `json:"targetAmount"`
	Deadline     *string `json:"deadline,omitempty"`
}

type SummaryItem struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

type Transaction struct {
	ID                 string              `json:"id"`
	UserID             string              `json:"userId"`
	Amount             float64             `json:"amount"`
	Type               TransactionType     `json:"type"`
	Category           *Category           `json:"category"`
	Date               string              `json:"date"`
	Description        *string             `json:"description,omitempty"`
	IsRecurring        bool                `json:"isRecurring"`
	RecurringFrequency *RecurringFrequency `json:"recurringFrequency,omitempty"`
}

type TransactionFilter struct {
	StartDate  *string          `json:"startDate,omitempty"`
	EndDate    *string          `json:"endDate,omitempty"`
	Type       *TransactionType `json:"type,omitempty"`
	CategoryID *int             `json:"categoryId,omitempty"`
}

type TrendPoint struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}

type UpdateBudgetInput struct {
	Amount     *float64 `json:"amount,omitempty"`
	CategoryID *int     `json:"categoryId,omitempty"`
	Period     *Period  `json:"period,omitempty"`
}

type UpdateCategoryInput struct {
	Name *string          `json:"name,omitempty"`
	Type *TransactionType `json:"type,omitempty"`
}

type UpdateGoalInput struct {
	Name          *string  `json:"name,omitempty"`
	TargetAmount  *float64 `json:"targetAmount,omitempty"`
	CurrentAmount *float64 `json:"currentAmount,omitempty"`
	Deadline      *string  `json:"deadline,omitempty"`
}

type UpdateTransactionInput struct {
	Amount             *float64            `json:"amount,omitempty"`
	Type               *TransactionType    `json:"type,omitempty"`
	CategoryID         *int                `json:"categoryId,omitempty"`
	Date               *string             `json:"date,omitempty"`
	Description        *string             `json:"description,omitempty"`
	IsRecurring        *bool               `json:"isRecurring,omitempty"`
	RecurringFrequency *RecurringFrequency `json:"recurringFrequency,omitempty"`
}

type UpdateUserInput struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type User struct {
	ID                      string                   `json:"id"`
	Name                    string                   `json:"name"`
	Email                   string                   `json:"email"`
	Transactions            []*Transaction           `json:"transactions"`
	Budgets                 []*Budget                `json:"budgets"`
	Goals                   []*Goal                  `json:"goals"`
	NotificationPreferences *NotificationPreferences `json:"notificationPreferences"`
}

type Period string

const (
	PeriodMonthly Period = "MONTHLY"
	PeriodYearly  Period = "YEARLY"
)

var AllPeriod = []Period{
	PeriodMonthly,
	PeriodYearly,
}

func (e Period) IsValid() bool {
	switch e {
	case PeriodMonthly, PeriodYearly:
		return true
	}
	return false
}

func (e Period) String() string {
	return string(e)
}

func (e *Period) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Period(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Period", str)
	}
	return nil
}

func (e Period) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RecurringFrequency string

const (
	RecurringFrequencyDaily   RecurringFrequency = "DAILY"
	RecurringFrequencyWeekly  RecurringFrequency = "WEEKLY"
	RecurringFrequencyMonthly RecurringFrequency = "MONTHLY"
	RecurringFrequencyYearly  RecurringFrequency = "YEARLY"
)

var AllRecurringFrequency = []RecurringFrequency{
	RecurringFrequencyDaily,
	RecurringFrequencyWeekly,
	RecurringFrequencyMonthly,
	RecurringFrequencyYearly,
}

func (e RecurringFrequency) IsValid() bool {
	switch e {
	case RecurringFrequencyDaily, RecurringFrequencyWeekly, RecurringFrequencyMonthly, RecurringFrequencyYearly:
		return true
	}
	return false
}

func (e RecurringFrequency) String() string {
	return string(e)
}

func (e *RecurringFrequency) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RecurringFrequency(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RecurringFrequency", str)
	}
	return nil
}

func (e RecurringFrequency) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReminderFrequency string

const (
	ReminderFrequencyDaily   ReminderFrequency = "DAILY"
	ReminderFrequencyWeekly  ReminderFrequency = "WEEKLY"
	ReminderFrequencyMonthly ReminderFrequency = "MONTHLY"
)

var AllReminderFrequency = []ReminderFrequency{
	ReminderFrequencyDaily,
	ReminderFrequencyWeekly,
	ReminderFrequencyMonthly,
}

func (e ReminderFrequency) IsValid() bool {
	switch e {
	case ReminderFrequencyDaily, ReminderFrequencyWeekly, ReminderFrequencyMonthly:
		return true
	}
	return false
}

func (e ReminderFrequency) String() string {
	return string(e)
}

func (e *ReminderFrequency) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReminderFrequency(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReminderFrequency", str)
	}
	return nil
}

func (e ReminderFrequency) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReportType string

const (
	ReportTypeIncomeExpense     ReportType = "INCOME_EXPENSE"
	ReportTypeCategoryBreakdown ReportType = "CATEGORY_BREAKDOWN"
	ReportTypeTrends            ReportType = "TRENDS"
)

var AllReportType = []ReportType{
	ReportTypeIncomeExpense,
	ReportTypeCategoryBreakdown,
	ReportTypeTrends,
}

func (e ReportType) IsValid() bool {
	switch e {
	case ReportTypeIncomeExpense, ReportTypeCategoryBreakdown, ReportTypeTrends:
		return true
	}
	return false
}

func (e ReportType) String() string {
	return string(e)
}

func (e *ReportType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReportType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReportType", str)
	}
	return nil
}

func (e ReportType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "INCOME"
	TransactionTypeExpense TransactionType = "EXPENSE"
)

var AllTransactionType = []TransactionType{
	TransactionTypeIncome,
	TransactionTypeExpense,
}

func (e TransactionType) IsValid() bool {
	switch e {
	case TransactionTypeIncome, TransactionTypeExpense:
		return true
	}
	return false
}

func (e TransactionType) String() string {
	return string(e)
}

func (e *TransactionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionType", str)
	}
	return nil
}

func (e TransactionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
