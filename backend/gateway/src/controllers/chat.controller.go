package controllers

import (
	"net/http"

	chatpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/chat"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/services"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatController struct{}

func (*ChatController) ChatWebSocket(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		utils.LoggerInfof("Failed to upgrade HTTP connection to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Start a goroutine to receive messages from the WebSocket connection and send them to the chat microservice
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				utils.LoggerErrorf("Failed to read message from WebSocket:", err)
				return
			}

			// Send the message to the chat microservice via gRPC
			_, err = services.SendMessage(&chatpb.SendMessageRequest{
				RoomId:  "your_room_id",
				UserId:  "your_user_id",
				Content: string(msg),
			})
			if err != nil {
				utils.LoggerErrorf("Failed to send message to chat microservice:", err)
				return
			}
		}
	}()

	// Start a goroutine to receive messages from the chat microservice and send them to the WebSocket connection
	go func() {
		stream, err := services.JoinRoom(&chatpb.JoinRoomRequest{
			Id: "your_room_id",
		})
		if err != nil {
			utils.LoggerErrorf("Failed to join chat room:", err)
			return
		}

		for {
			chatMsg, err := stream.Recv()
			if err != nil {
				utils.LoggerErrorf("Failed to receive message from chat microservice:", err)
				return
			}

			// Send the message to the WebSocket connection
			err = conn.WriteMessage(websocket.TextMessage, []byte(chatMsg.Content))
			if err != nil {
				utils.LoggerErrorf("Failed to write message to WebSocket:", err)
				return
			}
		}
	}()

	// Keep the goroutine alive
	select {}
}

func (*ChatController) CreateRoom(c *gin.Context) {}

func (*ChatController) GetAllRooms(c *gin.Context) {}
