package router

import (
	"github.com/DiarCode/next-golang-chat-app/gateway/src/controllers"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"github.com/gin-gonic/gin"
)

type RouterControllersType struct {
	Auth *controllers.AuthControllers
}

var Controllers *RouterControllersType

func NewRouter() *gin.Engine {
	Controllers := &RouterControllersType{
		Auth: &controllers.AuthControllers{},
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")

	// Auth
	api.POST("/auth/login", Controllers.Auth.Login)
	api.POST("/auth/signup", Controllers.Auth.Signup)

	api.GET("/", func(c *gin.Context) {
		utils.LoggerInfo("Entered into home page")
		utils.SendJson(c, 200, "Hello World")
	})

	return r
}
