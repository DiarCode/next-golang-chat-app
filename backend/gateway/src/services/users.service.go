package services

import (
	"context"

	userspb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/users"
)

func GetUserById(dto *userspb.GetUserByIdRequest) (*userspb.User, error) {
	resp, err := Clients.Users.GetUserById(context.Background(), dto)
	return resp, err
}
