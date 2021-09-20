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
	projectID          = "project-prometeo-v2"
	topicID            = "test-grpc-pub-sub"
	instanceConnection = "project-prometeo-v2:europe-north1:grpc-test"
	databaseName       = "test_click_impressions"
	user               = "root"
	password           = "root"
	createSchema       = "USE test_click_impressions; CREATE TABLE click (track_host VARCHAR(255), media_id VARCHAR(255), clickID INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(clickID));"
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

	// cfg := mysql.Cfg(instanceConnection, "root", "root")
	// cfg.DBName = databaseName
	// db, err := mysql.DialCfg(cfg)
	// if err != nil {
	// 	panic("Cannot connect to db >>" + err.Error())
	// }
	// fmt.Println(db)
	// _, err = db.Exec(createSchema)
	// if err != nil {
	// 	panic(err)
	// }

	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("Database connected successfully")
	// 	fmt.Println(db)
	// }

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
