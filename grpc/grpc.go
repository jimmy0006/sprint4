package grpc

import (
	"context"
	"log"
	"net"
	pb "sprint4/grpc/user"
	tokenizer "sprint4/token"

	"sprint4/redisConnection"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserStoreServer
}

func (s *server) Get(ctx context.Context, in *pb.Token) (*pb.User, error) {
	log.Printf("Received Token: %v", in.GetToken())
	tokenStruct, err := tokenizer.ExtractTokenEmail(in.GetToken())
	temp := &pb.User{}
	if err != nil {
		temp.Id = 0
		temp.Email = ""
		temp.Name = ""
		return temp, nil
	}
	email := tokenStruct.Email
	var dbConnector = new(redisConnection.DBconnector)
	getResult, err := dbConnector.GetHash(ctx, "temp")
	if err != nil {
		return nil, err
	}
	temp.Id = getResult.Id
	temp.Email = email
	temp.Name = getResult.Name
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
