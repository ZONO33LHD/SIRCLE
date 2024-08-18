# backendの実行コマンド

### kakeiboo-bffの実行
```
 go run cmd/server/main.go
```

### kakeibo-user-serviceの実行
```
 go run cmd/server/main.go
```

### kakeibo-bffのgRPCの生成
```
 protoc --go_out=. --go-grpc_out=. backend/kakeibo-bff/grpc/proto/kakeibo.proto
```

### kakeibo-user-serviceのgrpcの生成
```
 protoc --go_out=. --go-grpc_out=. backend/kakeibo-user-service/pkg/grpc/pb/user.proto
```

### GraphQLの生成
```
cd backend/kakeibo-bff
go run github.com/99designs/gqlgen generate
```

## delvの実行

### kakeibo-user-serviceでデバック実行
```
cd backend/kakeibo-user-service
dlv debug cmd/server/main.go
```
### ブレークポイント設定
```
b pkg/grpc/user_service.go:{10行目に設定するときは、10}
```

### デバッグ中のコマンド
```
// 次の行に進む
next or n

// 次のステップに進む
step or s

// ブレークポイントに到達するまでプログラムを実行
continue or s

```

### ブレークポイントの一覧を表示
```
breakpoints or bp
```

### 現在デバッグしている箇所周辺のソースコードを表示
```
list or l
```

### 現在デバッグしている箇所のスコープで有効な変数を表示
```
locals
```

###　式を評価して結果を表示
```
print or p
```

### goroutineの一覧を表示
```
goroutines or grs
```
