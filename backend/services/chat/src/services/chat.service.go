package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/DiarCode/next-golang-chat-app/chat/src/database"
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

func (s *ChatService) GetMessagesByRoomId(ctx context.Context, req *chatpb.GetMessagesByRoomIdRequest) (*chatpb.GetMessagesByRoomIdResponse, error) {
	var messages []*chatpb.ChatMessage

	queryResult := database.DB.Where("room_id = ?", req.RoomId).Find(&messages)

	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	return &chatpb.GetMessagesByRoomIdResponse{
		Messages: messages,
	}, nil
}

func (s *ChatService) GetAllRooms(ctx context.Context, req *chatpb.Empty) (*chatpb.GetAllRoomsResponse, error) {
	var rooms []*chatpb.Room
	res := database.DB.Find(&rooms)

	if res.Error != nil {
		return nil, res.Error
	}

	response := &chatpb.GetAllRoomsResponse{
		Rooms: rooms,
	}

	return response, nil
}

func (s *ChatService) GetRoomById(ctx context.Context, req *chatpb.GetRoomByIdRequest) (*chatpb.Room, error) {
	var room chatpb.Room

	res := database.DB.First(&room, req.Id)

	if res.Error != nil {
		return nil, res.Error
	}

	return &room, nil
}

func (s *ChatService) CreateRoom(ctx context.Context, req *chatpb.CreateRoomRequest) (*chatpb.Room, error) {
	room := chatpb.Room{
		Name: req.Name,
	}

	result := database.DB.Create(&room)

	if result.Error != nil {
		return nil, result.Error
	}

	return &chatpb.Room{Id: room.Id, Name: room.Name}, nil
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

		fmt.Println("Sending message ", chatMsg)
		err = stream.Send(&chatMsg)
		if err != nil {
			utils.LoggerInfof("Failed to send message to client:", err)
		}

		msg.Ack(true)
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
		strconv.FormatInt(int64(roomID), 10),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	chatMsg := chatpb.ChatMessage{
		RoomId:  roomID,
		UserId:  userID,
		Content: content,
	}

	// Create message in db]
	fmt.Println("RECCEEIEV MESSAGE", chatMsg)
	dbQueryResult := database.DB.Create(&chatMsg)
	if dbQueryResult.Error != nil {
		return nil, dbQueryResult.Error
	}

	body, err := json.Marshal(&chatMsg)
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
