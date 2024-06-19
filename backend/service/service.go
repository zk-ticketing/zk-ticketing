package service

import (
	"context"
	"net/http"

	"github.com/NFTGalaxy/zk-ticketing-server/openapi"
)

type APIService struct {
}

// NewAPIService creates a default api service
func NewAPIService() *APIService {
	return &APIService{}
}

// EventsEventIdGet - Get event details
func (s *APIService) EventsEventIdGet(ctx context.Context, eventId int32) (openapi.ImplResponse, error) {
	mockEvent := openapi.Event{
		Id:          "88c46eb5-cfe8-4892-b0b8-6a5859c642af",
		Name:        "Example Event",
		Description: "This is an example event.",
		Url:         "https://example.event",
		// AdminCode:   "",
	}
	return openapi.Response(http.StatusOK, mockEvent), nil
}

// EventsGet - Get list of events
func (s *APIService) EventsGet(ctx context.Context) (openapi.ImplResponse, error) {
	mockEvents := []openapi.Event{
		{
			Id:          "88c46eb5-cfe8-4892-b0b8-6a5859c642af",
			Name:        "Example Event",
			Description: "This is an example event.",
			Url:         "https://example.event",
			// AdminCode:   "",
		},
	}

	return openapi.Response(http.StatusOK, mockEvents), nil
}

// HealthGet - Check the health of the API
func (s *APIService) HealthGet(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusOK, "OK"), nil
}
