package main

import (
	"log"
	"net"

	"github.com/ikadgzl/inventory/authsvc/internal/config"
	"github.com/ikadgzl/inventory/authsvc/internal/db"
	"github.com/ikadgzl/inventory/authsvc/internal/handler"
	"github.com/ikadgzl/inventory/authsvc/internal/repository"
	"github.com/ikadgzl/inventory/common/proto/auth"
	"google.golang.org/grpc"
)

func main() {
	c := config.NewConfig()

	l, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := db.NewPostgresDB(c.DB)
	authRepo := repository.NewAuthRepository(db)
	authHndl := handler.NewAuthHandler(c.JWT, authRepo)

	grpcSrv := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcSrv, authHndl)

	log.Printf("auth service listening at %v", l.Addr())
	if err := grpcSrv.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
