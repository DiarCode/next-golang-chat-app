package services

import (
	"context"

	"github.com/DiarCode/next-golang-chat-app/auth/src/database"
	authpb "github.com/DiarCode/next-golang-chat-app/auth/src/gen/auth"
	"github.com/DiarCode/next-golang-chat-app/auth/src/models"
	"github.com/DiarCode/next-golang-chat-app/auth/src/utils"
	"github.com/badoux/checkmail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthService struct{}

func (*AuthService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.AuthResponse, error) {
	user := models.User{}
	query := models.User{Email: req.Email}
	err := database.DB.First(&user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	if !utils.ComparePasswords(user.Password, req.Password) {
		return nil, status.Error(codes.InvalidArgument, "Incorrect password")

	}
	tokenString, err := utils.GenerateToken(user)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create token")
	}

	response := &authpb.AuthResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Token:    tokenString,
	}

	return response, nil
}

func (*AuthService) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.AuthResponse, error) {
	password := utils.HashPassword([]byte(req.Password))
	err := checkmail.ValidateFormat(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Email is incorrect")
	}

	user := models.User{
		Password: password,
		Email:    req.Email,
		Username: req.Username,
	}

	found := models.User{}
	query := models.User{Email: req.Email}
	err = database.DB.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return nil, status.Error(codes.InvalidArgument, "User already exists")
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create user")
	}

	tokenString, err := utils.GenerateToken(user)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create token")
	}

	response := &authpb.AuthResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Token:    tokenString,
	}

	return response, nil
}
