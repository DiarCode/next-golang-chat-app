package router

import (
	"github.com/DiarCode/next-golang-chat-app/gateway/src/controllers"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"github.com/gin-gonic/gin"
)

type RouterControllersType struct {
	Auth  *controllers.AuthController
	Posts *controllers.PostsController
	Chat  *controllers.ChatController
}

var Controllers *RouterControllersType

func NewRouter() *gin.Engine {
	Controllers := &RouterControllersType{
		Auth:  &controllers.AuthController{},
		Posts: &controllers.PostsController{},
		Chat:  &controllers.ChatController{},
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")

	// Home
	r.GET("/", func(c *gin.Context) {
		utils.SendJson(c, 200, "Welcome to Meowchat API!")
	})

	// Auth
	api.POST("/auth/login", Controllers.Auth.Login)
	api.POST("/auth/signup", Controllers.Auth.Signup)

	// Posts
	api.POST("/posts", Controllers.Posts.CreatePost)
	api.GET("/posts", Controllers.Posts.GetAllPosts)
	api.GET("/posts/:id", Controllers.Posts.GetPost)

	// Chat
	api.GET("/chat/ws", Controllers.Chat.ChatWebSocket)
	api.GET("/chat/rooms", Controllers.Chat.GetAllRooms)
	api.POST("/chat/rooms", Controllers.Chat.CreateRoom)

	return r
}
