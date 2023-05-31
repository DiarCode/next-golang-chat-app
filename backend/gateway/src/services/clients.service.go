package services

import (
	"fmt"

	authpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/auth"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientsType struct {
	Auth authpb.AuthServiceClient
}

var Clients *ClientsType

const (
	authPort  = 50051
	eventPort = 50052
	queuePort = 50053
)

func InitServiceClients() *ClientsType {
	authClient := getAuthClient()

	clients := &ClientsType{
		Auth: authClient,
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
