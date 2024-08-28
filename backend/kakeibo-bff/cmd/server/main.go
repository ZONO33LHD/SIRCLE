package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-bff/graph"
	transactionpb "github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/pkg/grpc/pb"
	pb "github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/pkg/grpc/pb"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("USERS_PORT")
	if port == "" {
		port = defaultPort
	}

	// gRPCクライアントの作成
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// ユーザーサービスのgRPCサーバーに接続
	userConn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	userServiceClient := pb.NewUserServiceClient(userConn)

	// トランザクションサービスのgRPCサーバーに接続
	transactionConn, err := grpc.Dial("localhost:50052", opts...)
	if err != nil {
		log.Fatalf("Failed to connect to transaction service: %v", err)
	}
	defer transactionConn.Close()

	transactionServiceClient := transactionpb.NewTransactionServiceClient(transactionConn)

	// Resolverの初期化
	resolver := &graph.Resolver{
		UserServiceClient:        userServiceClient,
		TransactionServiceClient: transactionServiceClient,
	}

	// GraphQLサーバーの作成
	es := graph.NewExecutableSchema(graph.Config{Resolvers: resolver})
	srv := handler.NewDefaultServer(es)
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		return func(ctx context.Context) *graphql.Response {
			log.Printf("Received GraphQL operation: %s", graphql.GetOperationContext(ctx).RawQuery)
			return next(ctx)(ctx)
		}
	})
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		if err.Message == "variable.startDate must be defined" {
			return &gqlerror.Error{
				Message: "開始日（startDate）は必須です",
				Path:    err.Path,
				Extensions: map[string]interface{}{
					"code": "BAD_USER_INPUT",
				},
			}
		} else if err.Message == "variable.endDate must be defined" {
			err.Message = "終了日を指定してください"
			return &gqlerror.Error{
				Message: err.Message,
				Path:    err.Path,
				Extensions: map[string]interface{}{
					"code": "BAD_USER_INPUT",
				},
			}
		}
		err.Message = "内部サーバーエラーが発生しました"
		return err
	})

	// CORSミドルウェアの設定
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            false,
	})

	// GraphQLサーバーのハンドラーを設定
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsMiddleware.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
