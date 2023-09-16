package client

import (
	"context"
	"fmt"
0	"time"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cc, err := grpc.DialContext(ctx, url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
	return c
}

func (c *ProductServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return c.Client.FindOne(context.Background(), req)

}

func (c *ProductServiceClient) DecreaseStock(productId, orderId, qty int64) (*pb.DecreaseStockResponse, error) {

	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
		Qty:     qty,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
