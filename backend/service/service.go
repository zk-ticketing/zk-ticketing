package service

import (
	"context"
	"net/http"

	"github.com/NFTGalaxy/zk-ticketing-server/openapi"
	"github.com/NFTGalaxy/zk-ticketing-server/repos"
)

type APIService struct {
	dbClient *repos.Client
}

// NewAPIService creates a default api service
func NewAPIService(dbClient *repos.Client) *APIService {
	return &APIService{
		dbClient: dbClient,
	}
}

// EventsEventIdGet - Get event details
func (s *APIService) EventsEventIdGet(ctx context.Context, eventId string) (openapi.ImplResponse, error) {
	event, err := s.dbClient.Events.GetEventByID(ctx, eventId)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	return openapi.Response(http.StatusOK, MarshalEvent(event)), nil
}

// EventsGet - Get list of events
func (s *APIService) EventsGet(ctx context.Context) (openapi.ImplResponse, error) {
	events, err := s.dbClient.Events.ListEvents(ctx)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, MarshalEvents(events)), nil
}

// HealthGet - Check the health of the API
func (s *APIService) HealthGet(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusOK, "OK"), nil
}
