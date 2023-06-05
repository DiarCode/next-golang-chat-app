package services

import (
	"context"

	"github.com/DiarCode/next-golang-chat-app/posts/src/database"
	postspb "github.com/DiarCode/next-golang-chat-app/posts/src/gen/posts"
)

type PostsService struct{}

func (*PostsService) CreatePost(ctx context.Context, req *postspb.CreatePostRequest) (*postspb.Post, error) {
	postspb := &postspb.Post{
		Title:    req.Title,
		Body:     req.Body,
		AuthorId: req.AuthorId,
	}

	resultQuery := database.DB.Create(&postspb)
	if resultQuery.Error != nil {
		return nil, resultQuery.Error
	}

	return postspb, nil
}

func (*PostsService) GetPostById(ctx context.Context, req *postspb.GetPostByIdRequest) (*postspb.Post, error) {
	var post *postspb.Post

	resultQuery := database.DB.First(&post, req.Id)
	if resultQuery.Error != nil {
		return nil, resultQuery.Error
	}
	
	return post, nil
}

func (*PostsService) GetAllPosts(context.Context, *postspb.EmptyRequest) (*postspb.GetAllPostsResponse, error) {
	var posts []*postspb.Post

	resultQuery := database.DB.Find(&posts)
	if resultQuery.Error != nil {
		return nil, resultQuery.Error
	}

	response := &postspb.GetAllPostsResponse{
		Posts:  posts,
	}
	
	return response, nil
}
