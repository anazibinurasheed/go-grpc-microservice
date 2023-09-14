package auth

import (
	"fmt"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

// It will give as the interface that contains the api with initializing the connection from the generated code using the address of the service
func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	//using WithInsecure() because no SSL running

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
