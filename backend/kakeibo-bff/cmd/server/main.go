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
	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer conn.Close()

	userServiceClient := pb.NewUserServiceClient(conn)

	// Resolverの初期化
	resolver := &graph.Resolver{
		UserServiceClient: userServiceClient,
	}

	// GraphQLサーバーの作成
	es := graph.NewExecutableSchema(graph.Config{Resolvers: resolver})
	srv := handler.NewDefaultServer(es)
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		log.Printf("GraphQL error: %v", err)
		return err
	})

	// CORSミドルウェアの設定
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// GraphQLサーバーのハンドラーを設定
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsMiddleware.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
