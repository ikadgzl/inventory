generate:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	common/proto/**/*.proto

gateway:
	go run ./apigateway/cmd/main.go

authsvc:
	go run ./authsvc/cmd/main.go