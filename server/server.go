package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/logs"
	"server/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
}

func (s *server) Speaking(ctx context.Context, in *messagepb.SpeakRequest) (*messagepb.SpeakResponse, error) {
	log.Printf("Received Client Request : %v \n", in.GetClient_Request())
	log.Printf("server LIstenning 50051 ...\n")

	return &messagepb.SpeakResponse{Server_Response: "Hello 👋 !!! What are you doing " + in.GetClient_Request() + " ?"}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	logs.Error("Error in listening", err)

	// MyServer := grpc.NewServer()

	MyServer := grpc.NewServer(
    grpc.KeepaliveParams(keepalive.ServerParameters{
        MaxConnectionIdle: 5 * time.Minute,           // <--- This fixes it!
    }),
)

	messagepb.RegisterConversationServer(MyServer, &server{})

	fmt.Println("listening on 50051 ... ")
	err = MyServer.Serve(lis)

	logs.Error("Error in serving", err)

}
