package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
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
			log.Printf("AroundOperations: Context received")
			if ctx == nil {
				log.Printf("AroundOperations: Context is nil, creating new context")
				ctx = context.Background()
			}
			oc := graphql.GetOperationContext(ctx)
			if oc == nil {
				log.Printf("AroundOperations: OperationContext is nil, cannot proceed")
				return &graphql.Response{Errors: gqlerror.List{{Message: "OperationContextがnilです"}}}
			}
			log.Printf("Received GraphQL operation: %s", oc.RawQuery)
			resp := next(ctx)
			return resp(ctx)
		}
	})
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		log.Printf("Error Presenter: Received error: %v", e)
		err := graphql.DefaultErrorPresenter(ctx, e)
		log.Printf("Detailed error: %+v", err)

		return &gqlerror.Error{
			Message: fmt.Sprintf("エラーが発生しました: %s", err.Message),
			Path:    err.Path,
			Extensions: map[string]interface{}{
				"code": "INTERNAL_SERVER_ERROR",
			},
		}
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Printf("Panic occurred: %v", err)
		return fmt.Errorf("内部サーバーエラーが発生しました: %v", err)
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
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		log.Printf("Body: %s", string(body))
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		ctx := r.Context()
		ctx = graphql.WithOperationContext(ctx, &graphql.OperationContext{})
		r = r.WithContext(ctx)
		corsMiddleware.Handler(srv).ServeHTTP(w, r)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
