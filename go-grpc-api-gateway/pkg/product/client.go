package product

import (
	"fmt"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/product/pb"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running

	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)

	}

	return pb.NewProductServiceClient(cc)
}
