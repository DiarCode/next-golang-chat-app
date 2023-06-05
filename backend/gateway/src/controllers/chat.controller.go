package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	chatpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/chat"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/services"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatController struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// func (*ChatController) ChatWebSocket(c *gin.Context) {
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		fmt.Println("Failed to set websocket upgrade: %+v", err)
// 		return
// 	}

// 	for {
// 		t, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("Error in loop:", err)
// 			break
// 		}
// 		conn.WriteMessage(t, msg)
// 	}
// }

func (*ChatController) ChatWebSocket(c *gin.Context) {
	log.Println("EEEEEENTERRRR TO WEB")

	id := c.Param("id")
	if id == "" {
		utils.SendJsonError(c, http.StatusBadRequest, "Empty query parametr id", errors.New("missing id"))
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		utils.LoggerInfof("Failed to upgrade HTTP connection to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Start a goroutine to receive messages from the chat microservice and send them to the WebSocket connection

	stream, err := services.JoinRoom(&chatpb.JoinRoomRequest{
		Id: id,
	})
	if err != nil {
		utils.LoggerErrorf("Failed to join chat room:", err)
		return
	}

	for {
		fmt.Println("In Stream")
		chatMsg, err := stream.Recv()
		if err != nil {
			utils.LoggerErrorf("Failed to receive message from chat microservice:", err)
			return
		}

		fmt.Println("In Stream with message: ", chatMsg)

		err = conn.WriteJSON(chatMsg)
		if err != nil {
			utils.LoggerErrorf("Failed to write message to WebSocket:", err)
			return
		}
	}
}

func (*ChatController) SendMessage(c *gin.Context) {
	var req *chatpb.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendJsonError(c, http.StatusBadRequest, "Failed to parse request json", err)
		return
	}

	resp, err := services.SendMessage(req)
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to send message", err)
		return
	}

	utils.SendJson(c, http.StatusOK, resp)
}

func (*ChatController) CreateRoom(c *gin.Context) {
	var req chatpb.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendJsonError(c, http.StatusBadRequest, "Failed to parse request json", err)
		return
	}

	resp, err := services.CreateRoom(&req)
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to create room", err)
		return
	}

	utils.SendJson(c, http.StatusOK, resp)
}

func (*ChatController) GetAllMessagesByRoom(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		utils.SendJsonError(c, http.StatusBadRequest, "Empty query parametr id", errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to convert id", err)
		return
	}

	resp, err := services.GetAllMessagesByRoom(&chatpb.GetMessagesByRoomIdRequest{
		RoomId: id,
	})

	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to get message", err)
		return
	}

	utils.SendJson(c, http.StatusOK, resp.Messages)
}

func (*ChatController) GetAllRooms(c *gin.Context) {
	resp, err := services.GetAllRooms()

	if err != nil {
		utils.LoggerErrorf("Failed to get all rooms:", err)
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to get all rooms", err)
		return
	}

	utils.SendJson(c, http.StatusOK, resp.Rooms)
}

func (*ChatController) GetRoomById(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		utils.SendJsonError(c, http.StatusBadRequest, "Empty query parametr id", errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to convert id", err)
		return
	}

	resp, err := services.GetRoomById(&chatpb.GetRoomByIdRequest{
		Id: id,
	})

	if err != nil {
		utils.LoggerErrorf("Failed to get room:", err)
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to get room", err)
		return
	}

	utils.SendJson(c, http.StatusOK, resp)
}
