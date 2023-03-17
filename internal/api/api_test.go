package api

import (
	"myapp/internal/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewAPIClient(t *testing.T) {
	expected := &APIClient{Endpoint: "http://localhost:8080"}

	client := NewAPIClient("http://localhost:8080")

	if !reflect.DeepEqual(client, expected) {
		t.Errorf("unexpected result: expected=%v, got=%v", expected, client)
	}
}

func TestAPIClient_FetchData(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"name": "John Doe", "phone": "123-456-7890"},{"name": "Jane Doe", "phone": "987-654-3210"}]`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	client := NewAPIClient(server.URL)

	expected := []model.Person{
		{Name: "John Doe", Phone: "123-456-7890"},
		{Name: "Jane Doe", Phone: "987-654-3210"},
	}

	result, err := client.FetchData()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("unexpected result: expected=%v, got=%v", expected, result)
	}
}

func TestAPIClient_FetchData_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	client := NewAPIClient(server.URL)

	_, err := client.FetchData()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestAPIClient_FetchData_Decode_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"people": [{"name": "John Doe", "phone": "123-456-7890"},{"name": "Jane Doe", "phone": "987-654-3210"}]`))
	}))

	client := NewAPIClient(server.URL)

	_, err := client.FetchData()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
