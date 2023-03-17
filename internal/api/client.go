package api

import (
	"context"
	"myapp/internal/model"
	"myapp/pkg/grpc"
)

type APIClientWrapper struct {
	grpcClient grpc.MyServiceClient
}

func NewAPIClientWrapper(address string) (*APIClientWrapper, error) {
	conn, err := grpc.Dial(address)
	if err != nil {
		return nil, err
	}

	client := grpc.NewMyServiceClient(conn)

	return &APIClientWrapper{grpcClient: client}, nil
}

func (c *APIClientWrapper) SendData(data []model.Person) error {
	stream, err := c.grpcClient.SendData(context.Background())
	if err != nil {
		return err
	}

	for _, person := range data {
		err := stream.Send(&grpc.Person{Name: person.Name, Phone: person.Phone})
		if err != nil {
			return err
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}
