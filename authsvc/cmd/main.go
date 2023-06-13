package main

import (
	"log"
	"net"

	"github.com/ikadgzl/inventory/authsvc/internal/config"
	"github.com/ikadgzl/inventory/authsvc/internal/handler"
	"github.com/ikadgzl/inventory/common/proto/auth"
	"google.golang.org/grpc"
)

func main() {
	c := config.NewConfig()

	l, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	authHndl := handler.NewAuthHandler(c.JWT)
	auth.RegisterAuthServiceServer(grpcSrv, authHndl)

	log.Printf("auth service listening at %v", l.Addr())
	if err := grpcSrv.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
