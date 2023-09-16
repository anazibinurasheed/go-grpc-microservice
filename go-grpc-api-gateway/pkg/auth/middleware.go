package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (a *AuthMiddlewareConfig) AuthRequired(c *gin.Context) {
	authorization := c.Request.Header.Get("authorization")

	if authorization == "" {
		fmt.Println("1")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer")
	if len(token) < 2 {
		fmt.Println("2")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := a.svc.Client.Validate(ctx, &pb.ValidateRequest{
		Token: token[len(token)-1],
	})

	fmt.Println(res.Status, err)

	if err != nil || res.Status != http.StatusOK {
		fmt.Println("3")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println("Authenticated")
	c.Set("userId", res.UserId)
	c.Next()
}
