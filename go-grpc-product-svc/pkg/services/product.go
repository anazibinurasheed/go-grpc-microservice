package services

import (
	"context"
	"fmt"
	"net/http"

	"gihub.com/anazibinurasheed/go-grpc-microservice/go-grpc-product-svc/pkg/db"
	"gihub.com/anazibinurasheed/go-grpc-microservice/go-grpc-product-svc/pkg/models"
	"gihub.com/anazibinurasheed/go-grpc-microservice/go-grpc-product-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price
	fmt.Println(product.Stock)
	fmt.Println(req.Stock)

	id := 0
	err := s.H.DB.Raw(`select id from products where name= $1;`, product.Name).Scan(&id).Error
	if err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to check product",
		}, nil
	}

	if id != 0 {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  "Product already exist with the same name",
		}, nil
	}

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{

		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{

		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.H.DB.Raw("select * from products where id = ? limit 1;", req.Id).Scan(&product); result.Error != nil {

		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {

		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Raw("select * from stock_decrease_logs where order_id = ? limit 1;", req.OrderId).Scan(&log); result.Error == nil {

		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1

	if result := s.H.DB.Raw("update products set stock = ? where id = ?;", req.Id); result.Error != nil {

		return &pb.DecreaseStockResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to update stock",
		}, nil
	}

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{

		Status: http.StatusOK,
	}, nil

}
