package main

import (
	"fmt"
	"net"
	"os"

	"github.com/DiarCode/next-golang-chat-app/auth/src/config"
	"github.com/DiarCode/next-golang-chat-app/auth/src/database"
	authpb "github.com/DiarCode/next-golang-chat-app/auth/src/gen/auth"
	userspb "github.com/DiarCode/next-golang-chat-app/auth/src/gen/users"
	"github.com/DiarCode/next-golang-chat-app/auth/src/services"
	"github.com/DiarCode/next-golang-chat-app/auth/src/utils"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger()

	config.AppConfig = &config.AppConfigType{
		APP_PORT:    os.Getenv("AUTH_APP_PORT"),
		JWT_KEY:     os.Getenv("AUTH_JWT_KEY"),
		DB_HOST:     os.Getenv("AUTH_DB_HOST"),
		DB_USER:     os.Getenv("AUTH_DB_USER"),
		DB_PASSWORD: os.Getenv("AUTH_DB_PASSWORD"),
		DB_NAME:     os.Getenv("AUTH_DB_NAME"),
		DB_PORT:     os.Getenv("AUTH_DB_PORT"),
	}
	// config.AppConfig = &config.AppConfigType{
	// 	APP_PORT:    50051,
	// 	JWT_KEY:     "SSH256KEY",
	// 	DB_USER:     "postgres",
	// 	DB_PASSWORD: "postgres",
	// 	DB_NAME:     "meowchat_auth",
	// 	DB_PORT:     "5432",
	// 	DB_HOST:     "localhost",
	// }

	database.ConnectDB()

	server := grpc.NewServer()
	authpb.RegisterAuthServiceServer(server, &services.AuthService{})
	userspb.RegisterUserServiceServer(server, &services.UserService{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.AppConfig.APP_PORT))
	if err != nil {
		utils.LoggerFatalf("Failed to listen: %v", err)
	}

	utils.LoggerInfof("Server is running on port %v", config.AppConfig.APP_PORT)

	err = server.Serve(lis)
	if err != nil {
		utils.LoggerFatalf("Failed to serve: %v", err)
	}
}
