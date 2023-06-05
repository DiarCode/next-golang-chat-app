package controllers

import (
	"errors"
	"net/http"
	"strconv"

	authpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/auth"
	userspb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/users"
	models "github.com/DiarCode/next-golang-chat-app/gateway/src/models/auth"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/services"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (*AuthController) Login(c *gin.Context) {
	var request models.LoginJson
	err := c.ShouldBindJSON(&request)
	if err != nil {
		utils.SendJsonError(c, http.StatusBadRequest, "Failed to parse json", err)
	}

	resp, err := services.Login(&authpb.LoginRequest{Email: request.Email, Password: request.Password})
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to login", err)
	}

	utils.SendJson(c, http.StatusOK, resp)
}

func (*AuthController) Signup(c *gin.Context) {
	var request models.SignupJson
	err := c.ShouldBindJSON(&request)
	if err != nil {
		utils.SendJsonError(c, http.StatusBadRequest, "Failed to parse json", err)
	}

	resp, err := services.Signup(&authpb.SignupRequest{
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	})
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to signup", err)
	}

	utils.SendJson(c, http.StatusOK, resp)
}

func (*AuthController) GetUserById(c *gin.Context) {
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

	resp, err := services.GetUserById(&userspb.GetUserByIdRequest{
		Id: id,
	})

	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to get user", err)
		return
	}

	utils.SendJson(c, http.StatusOK, resp)
}
