package api

import (
	"context"
	"fmt"
	"myapp/internal/model"
	"myapp/pkg/grpc"
)

type APIServer struct {
	grpcServer *grpc.Server
}

func NewAPIServer() *APIServer {
	return &APIServer{grpcServer: grpc.NewServer()}
}

func (s *APIServer) StartGRPCServer(address string) error {
	listener, err := grpc.CreateListener(address)
	if err != nil {
		return err
	}

	grpc.RegisterMyServiceServer(s.grpcServer, s)

	return s.grpcServer.Serve(listener)
}

func (s *APIServer) SendData(ctx context.Context, req *grpc.DataRequest) (*grpc.DataResponse, error) {
	var people []model.Person
	for _, person := range req.People {
		people = append(people, model.Person{Name: person.Name, Phone: person.Phone})
	}

	fmt.Printf("Received %d people\n", len(people))

	return &grpc.DataResponse{Success: true}, nil
}
