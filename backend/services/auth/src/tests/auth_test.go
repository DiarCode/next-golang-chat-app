package tests

import (
	"context"
	"testing"

	authpb "github.com/DiarCode/next-golang-chat-app/auth/src/gen/auth"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSignup(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(BufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	reqAuth := &authpb.SignupRequest{
		Username: "testUsername",
		Email:    "test@example.com",
		Password: "password",
	}

	authResponse, err := client.Signup(ctx, reqAuth)

	assert.NoError(t, err)
	assert.NotNil(t, authResponse)
	assert.NotEmpty(t, authResponse.Id)
	assert.NotEmpty(t, authResponse.Token)
	assert.Equal(t, reqAuth.Username, authResponse.Username)
	assert.Equal(t, reqAuth.Email, authResponse.Email)
}

func TestSignupTheLogin(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(BufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	signupReq := &authpb.SignupRequest{
		Username: "testUsername",
		Email:    "test@example.com",
		Password: "password",
	}
	signupResp, err := client.Signup(ctx, signupReq)

	loginReq := &authpb.LoginRequest{
		Email:    signupReq.Email,
		Password: signupReq.Password,
	}
	loginResp, err := client.Login(ctx, loginReq)

	assert.NoError(t, err)
	assert.NotNil(t, loginResp)
	assert.NotEmpty(t, loginResp.Id)
	assert.NotEmpty(t, loginResp.Token)
	assert.Equal(t, loginResp.Username, signupResp.Username)
	assert.Equal(t, loginResp.Email, signupResp.Email)
}
