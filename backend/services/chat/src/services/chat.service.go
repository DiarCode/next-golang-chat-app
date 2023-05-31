package services

import (
	"context"
	"encoding/json"

	chatpb "github.com/DiarCode/next-golang-chat-app/chat/src/gen/chat"
	"github.com/DiarCode/next-golang-chat-app/chat/src/utils"
	"github.com/streadway/amqp"
)

type ChatService struct {
	conn *amqp.Connection
}

func NewChatServiceServer(conn *amqp.Connection) *ChatService {
	return &ChatService{conn: conn}
}

func (s *ChatService) GetAllRooms(ctx context.Context, req *chatpb.Empty) (*chatpb.GetAllRoomsResponse, error) {
	return nil, nil
}

func (s *ChatService) CreateRoom(ctx context.Context, req *chatpb.CreateRoomRequest) (*chatpb.Room, error) {
	// Logic to create a new chat room
	// ...

	roomID := 32
	roomName := "Name"

	return &chatpb.Room{Id: int64(roomID), Name: roomName}, nil
}

func (s *ChatService) JoinRoom(req *chatpb.JoinRoomRequest, stream chatpb.ChatService_JoinRoomServer) error {
	roomID := req.Id

	// Receive messages from RabbitMQ and send them to the client
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		roomID, // Queue name is the room ID
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		var chatMsg chatpb.ChatMessage
		err := json.Unmarshal(msg.Body, &chatMsg)
		if err != nil {
			utils.LoggerInfof("Failed to unmarshal message:", err)
			continue
		}

		err = stream.Send(&chatMsg)
		if err != nil {
			utils.LoggerInfof("Failed to send message to client:", err)
		}
	}

	return nil
}

func (s *ChatService) SendMessage(ctx context.Context, req *chatpb.SendMessageRequest) (*chatpb.SendMessageResponse, error) {
	roomID := req.RoomId
	userID := req.UserId
	content := req.Content

	// Publish the message to RabbitMQ
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		roomID, // Queue name is the room ID
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	chatMsg := &chatpb.ChatMessage{
		RoomId:  roomID,
		UserId:  userID,
		Content: content,
	}

	body, err := json.Marshal(chatMsg)
	if err != nil {
		return nil, err
	}

	err = ch.Publish(
		"",         // Exchange
		queue.Name, // Routing key is the queue name
		false,
		false,
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        body,
		},
	)
	if err != nil {
		return nil, err
	}

	return &chatpb.SendMessageResponse{Success: true}, nil
}
