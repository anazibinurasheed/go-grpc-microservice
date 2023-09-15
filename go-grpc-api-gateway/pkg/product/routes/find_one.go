package routes

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FindOne(c *gin.Context, psc pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()

	res, err := psc.FindOne(ctx, &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusCreated, &res)

}
