package service

import (
	"github.com/NFTGalaxy/zk-ticketing-server/openapi"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/events"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/users"
)

func MarshalEvent(event events.Event) openapi.Event {
	return openapi.Event{
		Id:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		Url:         event.Url,
		AdminCode:   event.AdminCode,
	}
}

func MarshalEvents(events []events.Event) []openapi.Event {
	marshaledEvents := make([]openapi.Event, len(events))
	for i, event := range events {
		marshaledEvents[i] = MarshalEvent(event)
	}
	return marshaledEvents
}

func MarshalUser(user users.User) openapi.User {
	return openapi.User{
		Id:                         user.ID,
		Email:                      user.Email,
		IdentityCommitment:         user.IdentityCommitment,
		EncryptedInternalNullifier: user.EncryptedInternalNullifier,
		EncryptedIdentitySecret:    user.EncryptedIdentitySecret,
		CreatedAt:                  user.CreatedAt.Time,
	}
}
