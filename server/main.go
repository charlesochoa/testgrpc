package main

import (
	"context"
	"fmt"
	"net"

	pb "testgrpc/notification"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
}

var requests = 0

const (
	projectID = "project-prometeo-v2"
	topicID   = "test-grpc-pub-sub"
)

func (s *server) Send(ctx context.Context, req *pb.SendItemReq) (*pb.SendItemResp, error) {

	fmt.Println("Receiving request: " + req.Item.Content)
	requests = requests + 1
	if requests > 4 {
		err := publish("4 requests exceded, countdown to 0")
		if err != nil {
			panic("cannot publish >> " + err.Error())
		}
		requests = 0
	}
	return &pb.SendItemResp{
		Id: req.Item.Id,
	}, nil
}

func publish(msg string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("get: %v", err)
	}
	fmt.Println("Published a message; msg ID:", id)
	return nil
}

func main() {

	lstnr, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("cannot create tcp connection")
	}
	serv := grpc.NewServer()
	pb.RegisterNotificationServiceServer(serv, &server{})
	if err = serv.Serve(lstnr); err != nil {
		panic("cannot start server")
	}
}
