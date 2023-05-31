package services

import (
	"context"

	authpb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/auth"
)

func Login(dto *authpb.LoginRequest) (*authpb.AuthResponse, error) {
	resp, err := Clients.Auth.Login(context.Background(), dto)
	return resp, err
}

func Signup(dto *authpb.SignupRequest) (*authpb.AuthResponse, error) {
	resp, err := Clients.Auth.Signup(context.Background(), dto)
	return resp, err
}
