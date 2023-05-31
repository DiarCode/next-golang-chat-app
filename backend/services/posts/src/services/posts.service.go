package services

import (
	"context"

	postspb "github.com/DiarCode/next-golang-chat-app/posts/src/gen/posts"
)

type PostsService struct{}

func (*PostsService) CreatePost(context.Context, *postspb.CreatePostRequest) (*postspb.Post, error) {
	return nil, nil
}

func (*PostsService) GetPost(context.Context, *postspb.GetPostRequest) (*postspb.Post, error) {
	return nil, nil
}

func (*PostsService) GetAllPosts(context.Context, *postspb.EmptyRequest) (*postspb.GetAllPostsResponse, error) {
	return nil, nil
}
