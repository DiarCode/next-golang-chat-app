package tests

import (
	"context"
	"testing"

	chatpb "github.com/DiarCode/next-golang-chat-app/chat/src/gen/chat"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreateChat(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(BufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := chatpb.NewChatServiceClient(conn)

	reqChat := &chatpb.CreateRoomRequest{
		Name: "RoomTest",
	}

	chat, err := client.CreateRoom(ctx, reqChat)

	assert.NoError(t, err)
	assert.NotNil(t, chat)
	assert.NotEmpty(t, chat.Id)
	assert.Equal(t, reqChat.Name, chat.Name)
}

func TestCreateChatThenGetById(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(BufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := chatpb.NewChatServiceClient(conn)

	chatCreate, err := client.CreateRoom(ctx, &chatpb.CreateRoomRequest{
		Name: "RoomTest",
	})

	if err != nil {
		t.Errorf("TestCreatePostThenGetById() Failed: %v", err)
	}

	chatGet, err := client.GetRoomById(ctx, &chatpb.GetRoomByIdRequest{Id: chatCreate.Id})

	assert.NoError(t, err)
	assert.NotNil(t, chatGet)
	assert.NotEmpty(t, chatGet.Id)
	assert.Equal(t, chatGet.Name, chatCreate.Name)
	assert.Equal(t, chatGet.Id, chatCreate.Id)
}
