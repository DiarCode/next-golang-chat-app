package services

import (
	"fmt"

	authpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/auth"
	chatpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/chat"
	postspb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/posts"
	userspb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/users"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientsType struct {
	Auth  authpb.AuthServiceClient
	Posts postspb.PostsServiceClient
	Chat  chatpb.ChatServiceClient
	Users userspb.UserServiceClient
}

var Clients *ClientsType

const (
	authPort  = 50051
	postsPort = 50052
	chatPort  = 50053
)

func InitServiceClients() *ClientsType {
	authClient := getAuthClient()
	postsClient := getPostsClient()
	chatClient := getChatClient()
	usersClient := getUsersClient()

	clients := &ClientsType{
		Auth:  authClient,
		Posts: postsClient,
		Chat:  chatClient,
		Users: usersClient,
	}

	return clients
}

func getAuthClient() authpb.AuthServiceClient {
	uri := fmt.Sprintf("localhost:%v", authPort)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.LoggerFatalf("Failed to connect to Auth client: %v", err)
	}

	return authpb.NewAuthServiceClient(conn)
}

func getUsersClient() userspb.UserServiceClient {
	uri := fmt.Sprintf("localhost:%v", authPort)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.LoggerFatalf("Failed to connect to Auth client: %v", err)
	}

	return userspb.NewUserServiceClient(conn)
}

func getPostsClient() postspb.PostsServiceClient {
	uri := fmt.Sprintf("localhost:%v", postsPort)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.LoggerFatalf("Failed to connect to Posts client: %v", err)
	}

	return postspb.NewPostsServiceClient(conn)
}

func getChatClient() chatpb.ChatServiceClient {
	uri := fmt.Sprintf("localhost:%v", chatPort)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.LoggerFatalf("Failed to connect to Posts client: %v", err)
	}

	return chatpb.NewChatServiceClient(conn)
}
