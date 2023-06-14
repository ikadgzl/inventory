generate:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	common/proto/**/*.proto

gateway:
	go run ./apigateway/cmd/main.go

authsvc:
	go run ./authsvc/cmd/main.go


DB_URL=postgres://nexus:nexus@localhost:5432/nexus?sslmode=disable

mcreate:
	migrate create -ext sql -dir ./$(svc)/db/migration -seq $(name)

migrate:
	migrate -path ./$(svc)/db/migration -database ${DB_URL} up