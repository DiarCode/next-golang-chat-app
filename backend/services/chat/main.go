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

	config.AppConfig = &config.AppConfigType{
		APP_PORT:    os.Getenv("CHATS_APP_PORT"),
		DB_HOST:     os.Getenv("CHATS_DB_HOST"),
		DB_USER:     os.Getenv("CHATS_DB_USER"),
		DB_PASSWORD: os.Getenv("CHATS_DB_PASSWORD"),
		DB_NAME:     os.Getenv("CHATS_DB_NAME"),
		DB_PORT:     os.Getenv("CHATS_DB_PORT"),
	}

	kafkaHost := os.Getenv("CHATS_KAFKA_HOST")
	config.QueueConfig = &config.QueueConfigType{
		KafkaURI: fmt.Sprintf("%v:9092", kafkaHost),
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
