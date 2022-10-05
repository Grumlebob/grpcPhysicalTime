package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	message := protos.ClientRequest{Text: "Client sent first handshake, with Syn flag True and Seq 0", Ack: 0, Seq: 0}
	fmt.Println(message.Text)

	//First handshake
	firstHandshake, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}
	fmt.Println(firstHandshake)

}
