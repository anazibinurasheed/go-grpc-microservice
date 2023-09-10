package product

import (
	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/auth"
	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)

}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
