package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Grumlebob/grpcPhysicalTime/protos"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	protos.ChatServiceServer
}

func (s *Server) GetTime(ctx context.Context, clientMessage *protos.ClientRequest) (*protos.ServerResponse, error) {
	t := time.Now()
	time2 := timestamppb.New(t)
	fmt.Println("T2:", time2)
	var response = &protos.ServerResponse{
		Timestamp: time2,
	}
	t3 := time.Now()
	time3 := timestamppb.New(t3)
	fmt.Println("T3:", time3)
	return response, nil
}

func main() {
	// Create listener tcp on port 9080
	listener, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterChatServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
