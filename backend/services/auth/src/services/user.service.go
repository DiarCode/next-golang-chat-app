package services

import (
	"context"

	"github.com/DiarCode/next-golang-chat-app/auth/src/database"
	userspb "github.com/DiarCode/next-golang-chat-app/auth/src/gen/users"
	"github.com/DiarCode/next-golang-chat-app/auth/src/models"
)

type UserService struct{}

func (*UserService) GetUserById(ctx context.Context, req *userspb.GetUserByIdRequest) (*userspb.User, error) {
	var user models.User
	res := database.DB.First(&user, req.Id)

	if res.Error != nil {
		return nil, res.Error
	}

	return &userspb.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
