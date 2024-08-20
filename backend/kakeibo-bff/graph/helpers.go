package graph

import "github.com/ZONO33LHD/sircle/backend/kakeibo-bff/graph/model"

// TransactionTypeToString converts a TransactionType to its string representation
func TransactionTypeToString(t model.TransactionType) string {
	switch t {
	case model.TransactionTypeIncome:
		return "INCOME"
	case model.TransactionTypeExpense:
		return "EXPENSE"
	// 他のケースも必要に応じて追加
	default:
		return "UNKNOWN"
	}
}
