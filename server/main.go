package main

import (
	pb "github.com/charlesochoa/testgrpc/notification"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
}
