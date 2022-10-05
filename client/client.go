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
	time1 := timestamppb.New(time.Now())
	fmt.Println("T1:", time1)
	clientRequest := protos.ClientRequest{
		Timestamp: time1,
	}
	response, err := c.GetTime(context.Background(), &clientRequest)
	if err != nil {
		log.Fatalf("Error when calling server: %s", err)
	}
	time4 := timestamppb.New(time.Now())
	fmt.Println("T2:", response.TimestampRecieved, "\nT3:", response.TimestampSent, "\nT4:", time4)

	// Calculate the time difference
	var timeClient = time4.AsTime().Sub(time1.AsTime())
	var timeServer = response.TimestampSent.AsTime().Sub(response.TimestampRecieved.AsTime())
	var roundTrip = timeClient - timeServer
	//Client sets time to T3 + roundtrip/2
	fmt.Println("bug spotting:", response.TimestampSent.AsTime())
	var clientSyncTime = response.TimestampSent.AsTime().Add(roundTrip / 2)

	//Prints
	fmt.Println("T4 - T1:", time4.AsTime().Sub(time1.AsTime()))
	fmt.Println("T3 - T2:", response.TimestampSent.AsTime().Sub(response.TimestampRecieved.AsTime()))
	fmt.Println("Roundtrip:", roundTrip)
	fmt.Println("Client synced time:", clientSyncTime)
	fmt.Println("Client before correction:", time.Now())
}
