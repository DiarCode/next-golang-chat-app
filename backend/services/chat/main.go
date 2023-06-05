package main

import (
	"fmt"
	"net"
	"os"

	"github.com/DiarCode/next-golang-chat-app/chat/src/config"
	"github.com/DiarCode/next-golang-chat-app/chat/src/database"
	chatpb "github.com/DiarCode/next-golang-chat-app/chat/src/gen/chat"
	"github.com/DiarCode/next-golang-chat-app/chat/src/services"
	"github.com/DiarCode/next-golang-chat-app/chat/src/utils"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger()

	config.Config = &config.AppConfig{
		APP_PORT:    os.Getenv("CHATS_APP_PORT"),
		DB_USER:     os.Getenv("CHATS_DB_USER"),
		DB_PASSWORD: os.Getenv("CHATS_DB_PASSWORD"),
		DB_NAME:     os.Getenv("CHATS_DB_NAME"),
		DB_PORT:     os.Getenv("CHATS_DB_PORT"),
	}
	// config.AppConfig = &config.AppConfigType{
	// 	APP_PORT:    50053,
	// 	JWT_KEY:     "SSH256KEY",
	// 	DB_USER:     "postgres",
	// 	DB_PASSWORD: "postgres",
	// 	DB_NAME:     "meowchat_chat",
	// 	DB_PORT:     "5432",
	// 	DB_HOST:     "localhost",
	// }

	config.QueueConfig = &config.QueueConfigType{
		KafkaURI: "localhost:9092",
	}

	database.ConnectDB()

	server := grpc.NewServer()
	chatpb.RegisterChatServiceServer(server, services.NewChatServiceServer())

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
