// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: grpc/proto/kakeibo.proto

package kakeibo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	KakeiboService_GetUser_FullMethodName                    = "/kakeibo.KakeiboService/GetUser"
	KakeiboService_GetTransactions_FullMethodName            = "/kakeibo.KakeiboService/GetTransactions"
	KakeiboService_GetBudgets_FullMethodName                 = "/kakeibo.KakeiboService/GetBudgets"
	KakeiboService_GetReports_FullMethodName                 = "/kakeibo.KakeiboService/GetReports"
	KakeiboService_GetCategories_FullMethodName              = "/kakeibo.KakeiboService/GetCategories"
	KakeiboService_CreateUser_FullMethodName                 = "/kakeibo.KakeiboService/CreateUser"
	KakeiboService_UpdateUser_FullMethodName                 = "/kakeibo.KakeiboService/UpdateUser"
	KakeiboService_CreateTransaction_FullMethodName          = "/kakeibo.KakeiboService/CreateTransaction"
	KakeiboService_UpdateTransaction_FullMethodName          = "/kakeibo.KakeiboService/UpdateTransaction"
	KakeiboService_DeleteTransaction_FullMethodName          = "/kakeibo.KakeiboService/DeleteTransaction"
	KakeiboService_CreateBudget_FullMethodName               = "/kakeibo.KakeiboService/CreateBudget"
	KakeiboService_UpdateBudget_FullMethodName               = "/kakeibo.KakeiboService/UpdateBudget"
	KakeiboService_CreateCategory_FullMethodName             = "/kakeibo.KakeiboService/CreateCategory"
	KakeiboService_UpdateCategory_FullMethodName             = "/kakeibo.KakeiboService/UpdateCategory"
	KakeiboService_DeleteCategory_FullMethodName             = "/kakeibo.KakeiboService/DeleteCategory"
	KakeiboService_SetGoal_FullMethodName                    = "/kakeibo.KakeiboService/SetGoal"
	KakeiboService_UpdateGoal_FullMethodName                 = "/kakeibo.KakeiboService/UpdateGoal"
	KakeiboService_SetNotificationPreferences_FullMethodName = "/kakeibo.KakeiboService/SetNotificationPreferences"
)

// KakeiboServiceClient is the client API for KakeiboService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// サービスの定義
type KakeiboServiceClient interface {
	// クエリ操作
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	GetBudgets(ctx context.Context, in *GetBudgetsRequest, opts ...grpc.CallOption) (*GetBudgetsResponse, error)
	GetReports(ctx context.Context, in *GetReportsRequest, opts ...grpc.CallOption) (*Report, error)
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	// ミューテーション操作
	CreateUser(ctx context.Context, in *CreateUserInput, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionInput, opts ...grpc.CallOption) (*Transaction, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error)
	CreateBudget(ctx context.Context, in *CreateBudgetInput, opts ...grpc.CallOption) (*Budget, error)
	UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*Budget, error)
	CreateCategory(ctx context.Context, in *CreateCategoryInput, opts ...grpc.CallOption) (*Category, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
	SetGoal(ctx context.Context, in *SetGoalInput, opts ...grpc.CallOption) (*Goal, error)
	UpdateGoal(ctx context.Context, in *UpdateGoalRequest, opts ...grpc.CallOption) (*Goal, error)
	SetNotificationPreferences(ctx context.Context, in *SetNotificationPreferencesRequest, opts ...grpc.CallOption) (*User, error)
}

type kakeiboServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKakeiboServiceClient(cc grpc.ClientConnInterface) KakeiboServiceClient {
	return &kakeiboServiceClient{cc}
}

func (c *kakeiboServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, KakeiboService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, KakeiboService_GetTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) GetBudgets(ctx context.Context, in *GetBudgetsRequest, opts ...grpc.CallOption) (*GetBudgetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBudgetsResponse)
	err := c.cc.Invoke(ctx, KakeiboService_GetBudgets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) GetReports(ctx context.Context, in *GetReportsRequest, opts ...grpc.CallOption) (*Report, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Report)
	err := c.cc.Invoke(ctx, KakeiboService_GetReports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, KakeiboService_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) CreateUser(ctx context.Context, in *CreateUserInput, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, KakeiboService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, KakeiboService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionInput, opts ...grpc.CallOption) (*Transaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Transaction)
	err := c.cc.Invoke(ctx, KakeiboService_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Transaction)
	err := c.cc.Invoke(ctx, KakeiboService_UpdateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTransactionResponse)
	err := c.cc.Invoke(ctx, KakeiboService_DeleteTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) CreateBudget(ctx context.Context, in *CreateBudgetInput, opts ...grpc.CallOption) (*Budget, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Budget)
	err := c.cc.Invoke(ctx, KakeiboService_CreateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*Budget, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Budget)
	err := c.cc.Invoke(ctx, KakeiboService_UpdateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryInput, opts ...grpc.CallOption) (*Category, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Category)
	err := c.cc.Invoke(ctx, KakeiboService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Category)
	err := c.cc.Invoke(ctx, KakeiboService_UpdateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, KakeiboService_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) SetGoal(ctx context.Context, in *SetGoalInput, opts ...grpc.CallOption) (*Goal, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Goal)
	err := c.cc.Invoke(ctx, KakeiboService_SetGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) UpdateGoal(ctx context.Context, in *UpdateGoalRequest, opts ...grpc.CallOption) (*Goal, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Goal)
	err := c.cc.Invoke(ctx, KakeiboService_UpdateGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kakeiboServiceClient) SetNotificationPreferences(ctx context.Context, in *SetNotificationPreferencesRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, KakeiboService_SetNotificationPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KakeiboServiceServer is the server API for KakeiboService service.
// All implementations must embed UnimplementedKakeiboServiceServer
// for forward compatibility.
//
// サービスの定義
type KakeiboServiceServer interface {
	// クエリ操作
	GetUser(context.Context, *GetUserRequest) (*User, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	GetBudgets(context.Context, *GetBudgetsRequest) (*GetBudgetsResponse, error)
	GetReports(context.Context, *GetReportsRequest) (*Report, error)
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	// ミューテーション操作
	CreateUser(context.Context, *CreateUserInput) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	CreateTransaction(context.Context, *CreateTransactionInput) (*Transaction, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*Transaction, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error)
	CreateBudget(context.Context, *CreateBudgetInput) (*Budget, error)
	UpdateBudget(context.Context, *UpdateBudgetRequest) (*Budget, error)
	CreateCategory(context.Context, *CreateCategoryInput) (*Category, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*Category, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error)
	SetGoal(context.Context, *SetGoalInput) (*Goal, error)
	UpdateGoal(context.Context, *UpdateGoalRequest) (*Goal, error)
	SetNotificationPreferences(context.Context, *SetNotificationPreferencesRequest) (*User, error)
	mustEmbedUnimplementedKakeiboServiceServer()
}

// UnimplementedKakeiboServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKakeiboServiceServer struct{}

func (UnimplementedKakeiboServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedKakeiboServiceServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedKakeiboServiceServer) GetBudgets(context.Context, *GetBudgetsRequest) (*GetBudgetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBudgets not implemented")
}
func (UnimplementedKakeiboServiceServer) GetReports(context.Context, *GetReportsRequest) (*Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReports not implemented")
}
func (UnimplementedKakeiboServiceServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedKakeiboServiceServer) CreateUser(context.Context, *CreateUserInput) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedKakeiboServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedKakeiboServiceServer) CreateTransaction(context.Context, *CreateTransactionInput) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedKakeiboServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedKakeiboServiceServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedKakeiboServiceServer) CreateBudget(context.Context, *CreateBudgetInput) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBudget not implemented")
}
func (UnimplementedKakeiboServiceServer) UpdateBudget(context.Context, *UpdateBudgetRequest) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBudget not implemented")
}
func (UnimplementedKakeiboServiceServer) CreateCategory(context.Context, *CreateCategoryInput) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedKakeiboServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedKakeiboServiceServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedKakeiboServiceServer) SetGoal(context.Context, *SetGoalInput) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGoal not implemented")
}
func (UnimplementedKakeiboServiceServer) UpdateGoal(context.Context, *UpdateGoalRequest) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoal not implemented")
}
func (UnimplementedKakeiboServiceServer) SetNotificationPreferences(context.Context, *SetNotificationPreferencesRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotificationPreferences not implemented")
}
func (UnimplementedKakeiboServiceServer) mustEmbedUnimplementedKakeiboServiceServer() {}
func (UnimplementedKakeiboServiceServer) testEmbeddedByValue()                        {}

// UnsafeKakeiboServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KakeiboServiceServer will
// result in compilation errors.
type UnsafeKakeiboServiceServer interface {
	mustEmbedUnimplementedKakeiboServiceServer()
}

func RegisterKakeiboServiceServer(s grpc.ServiceRegistrar, srv KakeiboServiceServer) {
	// If the following call pancis, it indicates UnimplementedKakeiboServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KakeiboService_ServiceDesc, srv)
}

func _KakeiboService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_GetBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBudgetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).GetBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_GetBudgets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).GetBudgets(ctx, req.(*GetBudgetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_GetReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).GetReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_GetReports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).GetReports(ctx, req.(*GetReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).CreateUser(ctx, req.(*CreateUserInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).CreateTransaction(ctx, req.(*CreateTransactionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_UpdateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_DeleteTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_CreateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBudgetInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).CreateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_CreateBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).CreateBudget(ctx, req.(*CreateBudgetInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_UpdateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).UpdateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_UpdateBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).UpdateBudget(ctx, req.(*UpdateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).CreateCategory(ctx, req.(*CreateCategoryInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_SetGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGoalInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).SetGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_SetGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).SetGoal(ctx, req.(*SetGoalInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_UpdateGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).UpdateGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_UpdateGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).UpdateGoal(ctx, req.(*UpdateGoalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KakeiboService_SetNotificationPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNotificationPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KakeiboServiceServer).SetNotificationPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KakeiboService_SetNotificationPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KakeiboServiceServer).SetNotificationPreferences(ctx, req.(*SetNotificationPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KakeiboService_ServiceDesc is the grpc.ServiceDesc for KakeiboService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KakeiboService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kakeibo.KakeiboService",
	HandlerType: (*KakeiboServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _KakeiboService_GetUser_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _KakeiboService_GetTransactions_Handler,
		},
		{
			MethodName: "GetBudgets",
			Handler:    _KakeiboService_GetBudgets_Handler,
		},
		{
			MethodName: "GetReports",
			Handler:    _KakeiboService_GetReports_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _KakeiboService_GetCategories_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _KakeiboService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _KakeiboService_UpdateUser_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _KakeiboService_CreateTransaction_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _KakeiboService_UpdateTransaction_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _KakeiboService_DeleteTransaction_Handler,
		},
		{
			MethodName: "CreateBudget",
			Handler:    _KakeiboService_CreateBudget_Handler,
		},
		{
			MethodName: "UpdateBudget",
			Handler:    _KakeiboService_UpdateBudget_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _KakeiboService_CreateCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _KakeiboService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _KakeiboService_DeleteCategory_Handler,
		},
		{
			MethodName: "SetGoal",
			Handler:    _KakeiboService_SetGoal_Handler,
		},
		{
			MethodName: "UpdateGoal",
			Handler:    _KakeiboService_UpdateGoal_Handler,
		},
		{
			MethodName: "SetNotificationPreferences",
			Handler:    _KakeiboService_SetNotificationPreferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/kakeibo.proto",
}
