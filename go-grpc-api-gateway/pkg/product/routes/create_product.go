package routes

import (
	"context"
	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(c *gin.Context, psc pb.ProductServiceClient) {
	body := CreateProductRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()

	res, err := psc.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusCreated, &res)

}
