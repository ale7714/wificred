package main

import (
	"context"
	"log"
	"net"
	"sync"
	pb "wificred/server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedWifiServiceServer
	savedCredentials *pb.WifiCredentials
	mu               sync.Mutex
}

func (s *server) SendCredentials(ctx context.Context, in *pb.WifiCredentials) (*pb.Confirmation, error) {
	s.mu.Lock()
	s.savedCredentials = in
	s.mu.Unlock()
	log.Printf("Credentials saved: SSID=%v, Password=%v", in.GetSsid(), in.GetPassword())
	return &pb.Confirmation{Message: "Credentials saved successfully"}, nil
}

func (s *server) GetCredentials(ctx context.Context, in *pb.Request) (*pb.WifiCredentials, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.savedCredentials == nil {
		return nil, status.Errorf(codes.NotFound, "No credentials saved")
	}
	log.Printf("Credentials retrieved: SSID=%v, Password=%v", s.savedCredentials.GetSsid(), s.savedCredentials.GetPassword())
	return s.savedCredentials, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterWifiServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
