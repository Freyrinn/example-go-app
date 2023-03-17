package api

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"myapp/pkg/grpc"
)

func TestAPIClient_SendData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStream := grpc.NewMockMyService_SendDataClient(ctrl)

	mockGRPCClient := grpc.NewMockMyServiceClient(ctrl)
	mockGRPCClient.EXPECT().SendData(gomock.Any(), gomock.Any()).Return(mockStream, nil)

	client := &MyServiceClient{grpcClient: mockGRPCClient}

	data := []*Person{
		{Name: "John Doe", Phone: "123-456-7890"},
		{Name: "Jane Doe", Phone: "987-654-3210"},
	}

	mockStream.EXPECT().Send(&DataRequest{People: data[0:1]}).Return(nil)
	mockStream.EXPECT().Send(&DataRequest{People: data[1:2]}).Return(nil)
	mockStream.EXPECT().CloseSend().Return(nil)
	mockStream.EXPECT().Recv().Return(nil, errors.New("something went wrong"))

	err := client.SendData(context.Background(), data)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestAPIClient_SendData_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGRPCClient := grpc.NewMockMyServiceClient(ctrl)

	expectedErr := errors.New("failed to send data")

	mockGRPCClient.EXPECT().SendData(gomock.Any(), gomock.Any()).Return(nil, expectedErr)

	client := &MyServiceClient{grpcClient: mockGRPCClient}

	data := []*Person{
		{Name: "John Doe", Phone: "123-456-7890"},
		{Name: "Jane Doe", Phone: "987-654-3210"},
	}

	err := client.SendData(context.Background(), data)
	if err != expectedErr {
		t.Errorf("unexpected error: expected=%v, got=%v", expectedErr, err)
	}
}
