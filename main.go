package main

import (
	"context"
	"log"
	"net"

	pb "CurtainServer/generated"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCurtainServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Hello)
	return &pb.HelloResponse{Hello: "Hello " + in.Hello}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCurtainServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
