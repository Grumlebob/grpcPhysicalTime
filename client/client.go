package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Grumlebob/grpcPhysicalTime/protos"
)

func main() {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := protos.NewChatServiceClient(conn)

	timeSync(c)
}

func timeSync(c protos.ChatServiceClient) {
	t := time.Now()
	time1 := timestamppb.New(t)
	fmt.Println("T1:", time1)
	clientRequest := protos.ClientRequest{
		Timestamp: time1,
	}
	response, err := c.GetTime(context.Background(), &clientRequest)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}
	time4 := timestamppb.New(t)
	fmt.Println("T2:", response.TimestampRecieved, "\n T3:", response.TimestampSent, "\n T4:", time4)

}
