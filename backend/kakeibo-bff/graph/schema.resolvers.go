package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"time"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-bff/graph/model"
	transactionpb "github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/pkg/grpc/pb"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/pkg/grpc/pb"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	resp, err := r.UserServiceClient.CreateUser(ctx, &pb.CreateUserRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	notificationPreferences := &model.NotificationPreferences{
		BudgetAlerts: false,
	}

	if resp.NotificationPreferences != nil {
		notificationPreferences.BudgetAlerts = resp.NotificationPreferences.BudgetNotifications
	}

	return &model.User{
		ID:                      resp.Id,
		Name:                    resp.Name,
		Email:                   resp.Email,
		Budgets:                 []*model.Budget{},
		Goals:                   []*model.Goal{},
		NotificationPreferences: notificationPreferences,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*model.Transaction, error) {
	if r.TransactionServiceClient == nil {
		return nil, fmt.Errorf("TransactionServiceClient is not initialized")
	}

	date, err := time.Parse(time.RFC3339, input.Date)
	if err != nil {
		return nil, fmt.Errorf("無効な日付形式です: %v", err)
	}

	resp, err := r.TransactionServiceClient.CreateTransaction(ctx, &transactionpb.CreateTransactionRequest{
		UserId:      input.UserID,
		Amount:      float32(input.Amount),
		Type:        TransactionTypeToString(input.Type),
		CategoryId:  int32(input.CategoryID),
		Date:        date.Format(time.RFC3339),
		Description: *input.Description,
		IsRecurring: input.IsRecurring,
	})

	if err != nil {
		return nil, fmt.Errorf("取引の作成に失敗しました: %v", err)
	}

	return &model.Transaction{
		UserID: resp.Transaction.UserId,
		Amount: float64(resp.Transaction.Amount),
		Type:   model.TransactionType(resp.Transaction.Type),
		Category: &model.Category{
			ID:   fmt.Sprintf("%d", resp.Transaction.CategoryId),
			Name: resp.Transaction.CategoryName,
			Type: model.TransactionType(resp.Transaction.Type),
		},
		Date:        resp.Transaction.Date,
		Description: &resp.Transaction.Description,
		IsRecurring: resp.Transaction.IsRecurring,
	}, nil
}

// UpdateTransaction is the resolver for the updateTransaction field.
func (r *mutationResolver) UpdateTransaction(ctx context.Context, id string, input model.UpdateTransactionInput) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: UpdateTransaction - updateTransaction"))
}

// DeleteTransaction is the resolver for the deleteTransaction field.
func (r *mutationResolver) DeleteTransaction(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTransaction - deleteTransaction"))
}

// CreateBudget is the resolver for the createBudget field.
func (r *mutationResolver) CreateBudget(ctx context.Context, input model.CreateBudgetInput) (*model.Budget, error) {
	panic(fmt.Errorf("not implemented: CreateBudget - createBudget"))
}

// UpdateBudget is the resolver for the updateBudget field.
func (r *mutationResolver) UpdateBudget(ctx context.Context, id string, input model.UpdateBudgetInput) (*model.Budget, error) {
	panic(fmt.Errorf("not implemented: UpdateBudget - updateBudget"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, input model.UpdateCategoryInput) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: UpdateCategory - updateCategory"))
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCategory - deleteCategory"))
}

// SetGoal is the resolver for the setGoal field.
func (r *mutationResolver) SetGoal(ctx context.Context, input model.SetGoalInput) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: SetGoal - setGoal"))
}

// UpdateGoal is the resolver for the updateGoal field.
func (r *mutationResolver) UpdateGoal(ctx context.Context, id string, input model.UpdateGoalInput) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: UpdateGoal - updateGoal"))
}

// GetIncomeExpenseSummary is the resolver for the getIncomeExpenseSummary field.
func (r *mutationResolver) GetIncomeExpenseSummary(ctx context.Context, startDate string, endDate string) (*model.IncomeExpenseSummary, error) {
	// トランザクションサービスクライアントが初期化されていることを確認
	if r.TransactionServiceClient == nil {
		return nil, fmt.Errorf("TransactionServiceClient が初期化されていません")
	}

	// 日付文字列をtime.Time型に変換
	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return nil, fmt.Errorf("無効な開始日付: %v", err)
	}
	end, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return nil, fmt.Errorf("無効な終了日付: %v", err)
	}

	// トランザクションサービスに収支サマリーのリクエストを送信
	resp, err := r.TransactionServiceClient.GetIncomeExpenseSummary(ctx, &transactionpb.GetIncomeExpenseSummaryRequest{
		StartDate: start.Format(time.RFC3339),
		EndDate:   end.Format(time.RFC3339),
	})
	if err != nil {
		return nil, fmt.Errorf("収支サマリーの取得に失敗しました: %v", err)
	}

	// レスポンスをGraphQLの型に変換
	incomeItems := make([]*model.IncomeItem, len(resp.IncomeItems))
	for i, item := range resp.IncomeItems {
		incomeItems[i] = &model.IncomeItem{
			Title:  item.Title,
			Amount: float64(item.Amount),
		}
	}

	expenseItems := make([]*model.ExpenseItem, len(resp.ExpenseItems))
	for i, item := range resp.ExpenseItems {
		expenseItems[i] = &model.ExpenseItem{
			Title:  item.Title,
			Amount: float64(item.Amount),
		}
	}

	return &model.IncomeExpenseSummary{
		IncomeItems:  incomeItems,
		ExpenseItems: expenseItems,
		Balance:      float64(resp.Balance),
	}, nil
}

// SetNotificationPreferences is the resolver for the setNotificationPreferences field.
func (r *mutationResolver) SetNotificationPreferences(ctx context.Context, userID string, input model.NotificationPreferencesInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: SetNotificationPreferences - setNotificationPreferences"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthPayload, error) {
	resp, err := r.UserServiceClient.Login(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token: resp.Token,
		User: &model.User{
			ID:    resp.User.Id,
			Name:  resp.User.Name,
			Email: resp.User.Email,
		},
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, userID string, filter *model.TransactionFilter) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: Transactions - transactions"))
}

// Budgets is the resolver for the budgets field.
func (r *queryResolver) Budgets(ctx context.Context, userID string) ([]*model.Budget, error) {
	panic(fmt.Errorf("not implemented: Budgets - budgets"))
}

// Reports is the resolver for the reports field.
func (r *queryResolver) Reports(ctx context.Context, userID string, typeArg model.ReportType, period model.Period) (*model.Report, error) {
	panic(fmt.Errorf("not implemented: Reports - reports"))
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
