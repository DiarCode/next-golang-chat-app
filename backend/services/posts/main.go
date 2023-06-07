package main

import (
	"fmt"
	"net"
	"os"

	"github.com/DiarCode/next-golang-chat-app/posts/src/config"
	"github.com/DiarCode/next-golang-chat-app/posts/src/database"
	postspb "github.com/DiarCode/next-golang-chat-app/posts/src/gen/posts"
	"github.com/DiarCode/next-golang-chat-app/posts/src/services"
	"github.com/DiarCode/next-golang-chat-app/posts/src/utils"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger()

	config.AppConfig = &config.AppConfigType{
		APP_PORT:    os.Getenv("POSTS_APP_PORT"),
		DB_HOST:     os.Getenv("POSTS_DB_HOST"),
		DB_USER:     os.Getenv("POSTS_DB_USER"),
		DB_PASSWORD: os.Getenv("POSTS_DB_PASSWORD"),
		DB_NAME:     os.Getenv("POSTS_DB_NAME"),
		DB_PORT:     os.Getenv("POSTS_DB_PORT"),
	}
	// config.AppConfig = &config.AppConfigType{
	// 	APP_PORT:    50052,
	// 	JWT_KEY:     "SSH256KEY",
	// 	DB_USER:     "postgres",
	// 	DB_PASSWORD: "postgres",
	// 	DB_NAME:     "meowchat_posts",
	// 	DB_PORT:     "5432",
	// 	DB_HOST:     "localhost",
	// }

	database.ConnectDB()

	server := grpc.NewServer()
	postspb.RegisterPostsServiceServer(server, &services.PostsService{})

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
