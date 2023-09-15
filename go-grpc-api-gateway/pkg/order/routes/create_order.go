package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(c *gin.Context, osc pb.OrderServiceClient) {

	body := CreateOrderRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := c.Get("userId")

	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()

	res, err := osc.CreateOrder(ctx, &pb.CreateOrderRequest{

		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusCreated, &res)
}
