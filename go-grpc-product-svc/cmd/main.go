package main

import (
	"fmt"
	"log"
	"net"

	"gihub.com/anazibinurasheed/go-grpc-microservice/go-grpc-product-svc/pkg/config"
	"gihub.com/anazibinurasheed/go-grpc-microservice/go-grpc-product-svc/pkg/db"
	"gihub.com/anazibinurasheed/go-grpc-microservice/go-grpc-product-svc/pkg/pb"
	"gihub.com/anazibinurasheed/go-grpc-microservice/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	fmt.Println(c.Port)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	fmt.Println("Product Svc on:", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
