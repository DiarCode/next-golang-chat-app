package services

import (
	"context"

	chatpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/chat"
)

func SendMessage(dto *chatpb.SendMessageRequest) (*chatpb.SendMessageResponse, error) {
	resp, err := Clients.Chat.SendMessage(context.Background(), dto)
	return resp, err
}

func JoinRoom(dto *chatpb.JoinRoomRequest) (chatpb.ChatService_JoinRoomClient, error) {
	return Clients.Chat.JoinRoom(context.Background(), dto)
}

func CreateRoom(dto *chatpb.CreateRoomRequest) (*chatpb.Room, error) {
	return Clients.Chat.CreateRoom(context.Background(), dto)
}

func GetAllRooms() (*chatpb.GetAllRoomsResponse, error) {
	return Clients.Chat.GetAllRooms(context.Background(), &chatpb.Empty{})
}

func GetRoomById(dto *chatpb.GetRoomByIdRequest) (*chatpb.Room, error) {
	return Clients.Chat.GetRoomById(context.Background(), dto)
}

func GetAllMessagesByRoom(dto *chatpb.GetMessagesByRoomIdRequest) (*chatpb.GetMessagesByRoomIdResponse, error) {
	return Clients.Chat.GetMessagesByRoomId(context.Background(), dto)
}
