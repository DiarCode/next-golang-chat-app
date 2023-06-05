package services

import (
	"context"

	postspb "github.com/DiarCode/next-golang-chat-app/gateway/src/gen/posts"
)

func GetAllPosts() (*postspb.GetAllPostsResponse, error) {
	resp, err := Clients.Posts.GetAllPosts(context.Background(), &postspb.EmptyRequest{})
	return resp, err
}

func GetPostById(dto *postspb.GetPostByIdRequest) (*postspb.Post, error) {
	resp, err := Clients.Posts.GetPostById(context.Background(), dto)
	return resp, err
}

func CreatePost(dto *postspb.CreatePostRequest) (*postspb.Post, error) {
	resp, err := Clients.Posts.CreatePost(context.Background(), dto)
	return resp, err
}
