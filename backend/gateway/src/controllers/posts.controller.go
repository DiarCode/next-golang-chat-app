package controllers

import (
	"net/http"
	"strconv"

	postspb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/posts"
	models "github.com/DiarCode/next-golang-chat-app/gateway/src/models/posts"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/services"
	"github.com/DiarCode/next-golang-chat-app/gateway/src/utils"
	"github.com/gin-gonic/gin"
)

type PostsController struct{}

func (*PostsController) GetAllPosts(c *gin.Context) {
	resp, err := services.GetAllPosts()
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to get all posts", err)
	}

	utils.SendJson(c, http.StatusOK, resp)
}

func (*PostsController) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.SendJsonError(c, http.StatusBadRequest, "Invalid query parametr", err)
		return
	}

	resp, err := services.GetPost(&postspb.GetPostRequest{Id: int64(id)})
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to get post", err)
	}

	utils.SendJson(c, http.StatusOK, resp)
}

func (*PostsController) CreatePost(c *gin.Context) {
	var request models.CreatePostJson
	err := c.ShouldBindJSON(&request)
	if err != nil {
		utils.SendJsonError(c, http.StatusBadRequest, "Failed to parse json", err)
	}

	resp, err := services.CreatePost(&postspb.CreatePostRequest{
		Title:    request.Title,
		Body:     request.Body,
		AuthorId: int64(request.AuthorId),
	})
	if err != nil {
		utils.SendJsonError(c, http.StatusInternalServerError, "Failed to create post", err)
	}

	utils.SendJson(c, http.StatusOK, resp)
}
