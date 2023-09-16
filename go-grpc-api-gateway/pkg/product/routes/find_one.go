package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type FindOneData struct {
	Id    int64
	Name  string
	Sku   string
	Stock int64
	Price int64
}

type FindOneResponse struct {
	Status int64
	Error  string
	Data   FindOneData
}

func FindOne(c *gin.Context, psc pb.ProductServiceClient) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusExpectationFailed, errors.New("invalid param"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()

	res, err := psc.FindOne(ctx, &pb.FindOneRequest{

		Id: int64(id)})

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusCreated, FindOneResponse{
		Status: res.Status,
		Error:  res.Error,
		Data: FindOneData{
			Id:    res.Data.Id,
			Name:  res.Data.Name,
			Sku:   res.Data.Sku,
			Stock: res.Data.Stock,
			Price: res.Data.Price,
		},
	},
	)

}
