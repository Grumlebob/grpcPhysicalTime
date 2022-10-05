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
	time2 := timestamppb.New(time.Now().Local().UTC())
	fmt.Println("T2:", time2.AsTime())
	fmt.Println("Server current time:", time.Now())
	var response = &protos.ServerResponse{
		TimestampRecieved: time2,
		TimestampSent:     timestamppb.New(time.Now()),
	}
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
