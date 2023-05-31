package main

import (
	"fmt"
	"net"

	"github.com/DiarCode/next-golang-chat-app/chat/src/config"
	chatpb "github.com/DiarCode/next-golang-chat-app/chat/src/gen/chat"
	"github.com/DiarCode/next-golang-chat-app/chat/src/services"
	"github.com/DiarCode/next-golang-chat-app/chat/src/utils"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger()

	// config.Config = &config.AppConfig{
	// 	APP_PORT:    50051,
	// 	JWT_KEY:     os.Getenv("JWT_KEY"),
	// 	DB_USER:     os.Getenv("DB_USER"),
	// 	DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	// 	DB_NAME:     os.Getenv("DB_NAME"),
	// 	DB_PORT:     os.Getenv("DB_PORT"),
	// }
	config.AppConfig = &config.AppConfigType{
		APP_PORT:    50053,
		JWT_KEY:     "SSH256KEY",
		DB_USER:     "postgres",
		DB_PASSWORD: "postgres",
		DB_NAME:     "kezek_auth",
		DB_PORT:     "5432",
		DB_HOST:     "localhost",
	}

	config.QueueConfig = &config.RabbitmqConfigType{
		AmqpURI:    "amqp://guest:guest@localhost:5672/",
		Exchange:   "chat_exchange",
		RoutingKey: "chat_message",
	}

	conn, err := amqp.Dial(config.QueueConfig.AmqpURI)
	if err != nil {
		utils.LoggerFatalf("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	// database.ConnectDB()

	server := grpc.NewServer()
	chatpb.RegisterChatServiceServer(server, services.NewChatServiceServer(conn))

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
