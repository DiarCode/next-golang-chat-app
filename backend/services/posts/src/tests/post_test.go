package tests

import (
	"context"
	"testing"

	postspb "github.com/DiarCode/next-golang-chat-app/posts/src/gen/posts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreatePost(t *testing.T) {
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

	client := postspb.NewPostsServiceClient(conn)

	reqPost := &postspb.CreatePostRequest{
		Title:    "Post title",
		Body:     "Post body",
		AuthorId: 1,
	}

	post, err := client.CreatePost(ctx, reqPost)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.NotEmpty(t, post.Id)
	assert.Equal(t, reqPost.Title, post.Title)
	assert.Equal(t, reqPost.Body, post.Body)
	assert.Equal(t, reqPost.AuthorId, post.AuthorId)
	assert.NotZero(t, post.PublishedAt)
}

func TestCreatePostThenGetById(t *testing.T) {
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

	client := postspb.NewPostsServiceClient(conn)

	postCreate, err := client.CreatePost(ctx, &postspb.CreatePostRequest{
		Title:    "Post title",
		Body:     "Post body",
		AuthorId: 1,
	})
	
	if err != nil {
		t.Errorf("TestCreatePostThenGetById() Failed: %v", err)
	}

	postGet, err := client.GetPostById(ctx, &postspb.GetPostByIdRequest{Id: postCreate.Id})

	assert.NoError(t, err)
	assert.NotNil(t, postGet)
	assert.NotEmpty(t, postGet.Id)
	assert.Equal(t, postGet.Title, postCreate.Title)
	assert.Equal(t, postGet.Body, postCreate.Body)
	assert.Equal(t, postGet.AuthorId, postCreate.AuthorId)
	assert.Equal(t, postGet.PublishedAt, postCreate.PublishedAt)
}
