package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/charlesochoa/testgrpc/notification/notification"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *server) Send(ctx context.Context, req *pb.SendItemReq) (*pb.SendItemResp, error) {
	fmt.Println("Receiving request: " + req.Item.Content)
	return &pb.SendItemResp{
		Id: req.Item.Id,
	}, nil
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
