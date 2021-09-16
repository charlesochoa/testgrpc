package main

import (
	"context"
	"fmt"

	pb "github.com/charlesochoa/testgrpc/notification/notification"

	"google.golang.org/grpc"
)

func generateId() string {
	return "1234567"
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	serviceClient := pb.NewNotificationServiceClient(conn)
	res, err := serviceClient.Send(context.Background(), &pb.SendItemReq{
		Item: &pb.Item{
			Id:      generateId(),
			Content: "Esto es un item desde cliente",
		},
	})
	if err != nil {
		panic("cannot send >> " + err.Error())
	}
	fmt.Println(res.Id)
}

// projects/project-prometeo-v2/topics/test-grpc-pub-sub
