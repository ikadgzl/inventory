package auth

import (
	"log"

	"github.com/ikadgzl/inventory/apigateway/internal/config"
	"github.com/ikadgzl/inventory/common/proto/auth"
	"google.golang.org/grpc"
)

func AuthServiceClient(c *config.Config) auth.AuthServiceClient {
	// TODO: tls
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return auth.NewAuthServiceClient(cc)
}
