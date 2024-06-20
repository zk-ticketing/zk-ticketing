// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ZK Ticketing API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

package openapi


import (
	"time"
)



type TicketCredential struct {

	Id string `json:"id,omitempty"`

	Email string `json:"email,omitempty"`

	EventId string `json:"event_id,omitempty"`

	Data string `json:"data,omitempty"`

	IssuedAt time.Time `json:"issued_at,omitempty"`

	ExpireAt time.Time `json:"expire_at,omitempty"`
}

// AssertTicketCredentialRequired checks if the required fields are not zero-ed
func AssertTicketCredentialRequired(obj TicketCredential) error {
	return nil
}

// AssertTicketCredentialConstraints checks if the values respects the defined constraints
func AssertTicketCredentialConstraints(obj TicketCredential) error {
	return nil
}