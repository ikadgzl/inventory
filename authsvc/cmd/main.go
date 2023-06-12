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
	log.Printf("auth service listening at %v", l.Addr())

	gs := grpc.NewServer()
	as := handler.NewAuthServer(c.JWT)
	auth.RegisterAuthServiceServer(gs, as)

	if err := gs.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
