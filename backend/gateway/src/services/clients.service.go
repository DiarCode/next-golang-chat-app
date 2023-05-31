package services

import (
	"fmt"

	authpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/auth"
	postspb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/posts"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientsType struct {
	Auth  authpb.AuthServiceClient
	Posts postspb.PostsServiceClient
}

var Clients *ClientsType

const (
	authPort  = 50051
	postsPort = 50052
)

func InitServiceClients() *ClientsType {
	authClient := getAuthClient()
	postsClient := getPostsClient()

	clients := &ClientsType{
		Auth:  authClient,
		Posts: postsClient,
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

func getPostsClient() postspb.PostsServiceClient {
	uri := fmt.Sprintf("localhost:%v", postsPort)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.LoggerFatalf("Failed to connect to Posts client: %v", err)
	}

	return postspb.NewPostsServiceClient(conn)
}
