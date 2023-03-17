package service

import (
	"log"
	"myapp/internal/api"
	"myapp/internal/model"
	"myapp/pkg/grpc"
	"net/http"
)

type Service struct {
	apiClient  *api.APIClient
	grpcClient *grpc.MyServiceClient
}

func NewService(apiClient *api.APIClient, grpcClient *grpc.MyServiceClient) *Service {
	return &Service{apiClient: apiClient, grpcClient: grpcClient}
}

func (s *Service) FetchData() ([]model.Person, error) {
	data, err := s.apiClient.FetchData()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := s.FetchData()
	if err != nil {
		log.Printf("Failed to fetch data: %v", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	err = s.SendData(data)
	if err != nil {
		log.Printf("Failed to send data to gRPC service: %v", err)
		http.Error(w, "Failed to send data to gRPC service", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) SendData(data []model.Person) error {
	stream, err := (*s.grpcClient).SendData()
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
