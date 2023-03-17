package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"myapp/internal/api"
	"myapp/internal/model"
	"myapp/internal/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Set environment variables for the test.
	os.Setenv("API_ENDPOINT", "http://localhost:8080")
	os.Setenv("GRPC_ADDRESS", "localhost:50051")

	// Create a mock API client and gRPC service.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := api.NewMockAPIClient(ctrl)
	mockService := service.NewMockService(ctrl)

	// Set up the gRPC service to return a successful response.
	data := []model.Person{
		{Name: "John Doe", Phone: "123-456-7890"},
		{Name: "Jane Doe", Phone: "987-654-3210"},
	}
	mockService.EXPECT().GetData().Return(data, nil)

	// Set up the API client to return the expected response.
	expected := `[{"name":"John Doe","phone":"123-456-7890"},{"name":"Jane Doe","phone":"987-654-3210"}]`
	mockAPIClient.EXPECT().FetchData().Return(data, nil)

	// Replace the real API client and gRPC service with the mock versions.
	apiClient = mockAPIClient
	service = mockService

	// Start the HTTP server.
	ts := httptest.NewServer(http.HandlerFunc(HandleHTTP))
	defer ts.Close()

	// Send an HTTP GET request to the server.
	resp, err := http.Get(ts.URL + "/data")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Verify that the response body matches the expected value.
	assert.Equal(t, expected, string(body), "unexpected response body")
}
