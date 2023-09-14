package services

import (
	"context"
	"net/http"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-auth-svc/pkg/db"
	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-auth-svc/pkg/models"
	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-auth-svc/pkg/pb"
	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-auth-svc/pkg/utils"
)

// Server is a struct containing DB as H and JwtWrapper as Jwt
type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{

			Status: http.StatusConflict,
			Error:  "E-mail already exists",
		}, nil
	}

	user.Email = req.Email
	var ok bool
	user.Password, ok = utils.HashPassword(req.Password)

	if !ok {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to hash password",
		}, nil
	}

	s.H.DB.Create(&user)
	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{

			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
