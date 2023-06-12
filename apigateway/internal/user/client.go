package user

import (
	"log"

	"github.com/ikadgzl/inventory/apigateway/internal/config"
	"github.com/ikadgzl/inventory/common/proto/user"
	"google.golang.org/grpc"
)

func NewUserServiceClient(c *config.Config) user.UserServiceClient {
	// TODO: tls
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return user.NewUserServiceClient(cc)
}
