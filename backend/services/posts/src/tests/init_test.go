package tests

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	postspb "github.com/DiarCode/next-golang-chat-app/posts/src/gen/posts"
	"github.com/DiarCode/next-golang-chat-app/posts/src/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func initServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	postspb.RegisterPostsServiceServer(s, &services.PostsService{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func BufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	//Before all
	initServer()

	code := m.Run()

	//After all
	os.Exit(code)
}
