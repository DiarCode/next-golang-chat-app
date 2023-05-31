package main

import (
	"fmt"

	"github.com/DiarCode/next-golang-chat-app/gateway/src/router"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/services"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
)

const (
	port = 8080
)

func main() {
	utils.InitLogger()

	services.Clients = services.InitServiceClients()

	router := router.NewRouter()

	utils.LoggerInfof("Running http server listening on port %v", port)
	err := router.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		utils.LoggerFatalf("Failed to run server: %v", err)
	}

}
