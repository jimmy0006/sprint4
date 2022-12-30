package grpc

import (
	"context"
	"log"
	"net"
	pb "sprint4/grpc/user"

	"sprint4/redisConnection"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserStoreServer
}

func (s *server) Get(ctx context.Context, in *pb.Token) (*pb.User, error) {
	log.Printf("Received Token: %v", in.GetToken())
	var dbConnector = new(redisConnection.DBconnector)
	temp := &pb.User{}
	dbConnector.GetHash(ctx, "temp")
	temp.Id = 10
	temp.Email = "asdfasdf"
	temp.Name = "HI!"
	return temp, nil
}

func GRPC() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserStoreServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
