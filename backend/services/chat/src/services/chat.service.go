package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DiarCode/next-golang-chat-app/chat/src/config"
	"github.com/DiarCode/next-golang-chat-app/chat/src/database"
	chatpb "github.com/DiarCode/next-golang-chat-app/chat/src/gen/chat"
	"github.com/DiarCode/next-golang-chat-app/chat/src/utils"
	"github.com/segmentio/kafka-go"
)

type ChatService struct {
}

func NewChatServiceServer() *ChatService {
	return &ChatService{}
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
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.QueueConfig.KafkaURI},
		Topic:   fmt.Sprintf("chat-%v", req.Id),
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			utils.Logger.Sugar().Errorf("Error reading message from Kafka:", err)
			continue
		}

		var message chatpb.ChatMessage
		err = json.Unmarshal(m.Value, &message)
		if err != nil {
			utils.LoggerErrorf("Failed to unmarshal message:", err)
			continue
		}

		err = stream.Send(&message)
		if err != nil {
			utils.LoggerInfof("Failed to send message to client:", err)
			continue
		}

		err = reader.CommitMessages(context.Background(), m)
		if err != nil {
			utils.LoggerErrorf("Error committing message offset:", err)
			continue
		}
	}
}

func (s *ChatService) SendMessage(ctx context.Context, req *chatpb.SendMessageRequest) (*chatpb.SendMessageResponse, error) {
	chatMsg := chatpb.ChatMessage{
		RoomId:   req.RoomId,
		UserId:   req.UserId,
		Content:  req.Content,
		SendedAt: time.Now().Unix(),
	}

	// Create message in db
	dbQueryResult := database.DB.Create(&chatMsg)
	if dbQueryResult.Error != nil {
		return nil, dbQueryResult.Error
	}

	body, err := json.Marshal(&chatMsg)
	if err != nil {
		return nil, err
	}

	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.QueueConfig.KafkaURI},
		Topic:   fmt.Sprintf("chat-%v", req.RoomId),
	})
	err = producer.WriteMessages(ctx, kafka.Message{
		Value: body,
	})
	if err != nil {
		return nil, err
	}

	return &chatpb.SendMessageResponse{Success: true}, nil
}
