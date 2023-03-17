package service

import (
	"errors"
	"myapp/internal/api"
	"myapp/internal/model"
	"myapp/pkg/grpc"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestService_FetchData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPIClient := api.NewMockAPIClient(ctrl)

	data := []model.Person{
		{Name: "John Doe", Phone: "123-456-7890"},
		{Name: "Jane Doe", Phone: "987-654-3210"},
	}

	mockAPIClient.EXPECT().FetchData().Return(data, nil)

	svc := NewService(mockAPIClient, nil)

	result, err := svc.FetchData()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != len(data) {
		t.Errorf("unexpected result: expected=%v, got=%v", data, result)
	}
}

func TestService_FetchData_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAPIClient := api.NewMockAPIClient(ctrl)

	expectedErr := errors.New("failed to fetch data")

	mockAPIClient.EXPECT().FetchData().Return(nil, expectedErr)

	svc := NewService(mockAPIClient, nil)

	_, err := svc.FetchData()
	if err == nil {
		t.Errorf("expected error: %v, got nil", expectedErr)
	}
}

func TestService_SendData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGRPCClient := grpc.NewMockMyServiceClient(ctrl)

	data := []model.Person{
		{Name: "John Doe", Phone: "123-456-7890"},
		{Name: "Jane Doe", Phone: "987-654-3210"},
	}

	mockStream := grpc.NewMockMyService_SendDataClient(ctrl)

	mockStream.EXPECT().Send(&grpc.Person{Name: "John Doe", Phone: "123-456-7890"}).Return(nil)
	mockStream.EXPECT().Send(&grpc.Person{Name: "Jane Doe", Phone: "987-654-3210"}).Return(nil)
	mockStream.EXPECT().CloseAndRecv().Return(nil, nil)

	mockGRPCClient.EXPECT().SendData(gomock.Any()).Return(mockStream, nil)

	svc := NewService(nil, &mockGRPCClient)

	err := svc.SendData(data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestService_SendData_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGRPCClient := grpc.NewMockMyServiceClient(ctrl)

	expectedErr := errors.New("failed to send data")

	mockStream := grpc.NewMockMyService_SendDataClient(ctrl)

	mockGRPCClient.EXPECT().SendData(gomock.Any()).Return(mockStream, expectedErr)

	svc := NewService(nil, &mockGRPCClient)

	err := svc.SendData(nil)
	if err == nil {
		t.Errorf("expected error: %v, got nil", expectedErr)
	}
}
