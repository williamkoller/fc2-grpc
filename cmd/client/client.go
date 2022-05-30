package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/williamkoller/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)

	// AddUserVerbose(client)

	AddUsers(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id: "0",
		Name: "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id: "0",
		Name: "Joao",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
		break
	}

	if err != nil {
	log.Fatalf("Could not receive the message: %v", err)
	}

	fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id: "w1",
			Name: "William",
			Email: "william@mail.com",
		},
		&pb.User{
			Id: "w2",
			Name: "William 2",
			Email: "william2@mail.com",
		},
		&pb.User{
			Id: "w3",
			Name: "William 3",
			Email: "william3 @mail.com",
		},
		&pb.User{
			Id: "w4",
			Name: "William 4",
			Email: "william4@mail.com",
		},
		&pb.User{
			Id: "w5",
			Name: "William 5",
			Email: "william5@mail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}