// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Proof Pass API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

package openapi


import (
	"time"
)



type UnencryptedTicketCredential struct {

	Id string `json:"id,omitempty"`

	EventId string `json:"event_id,omitempty"`

	Credential string `json:"credential,omitempty"`

	IssuedAt time.Time `json:"issued_at,omitempty"`

	ExpireAt time.Time `json:"expire_at,omitempty"`
}

// AssertUnencryptedTicketCredentialRequired checks if the required fields are not zero-ed
func AssertUnencryptedTicketCredentialRequired(obj UnencryptedTicketCredential) error {
	return nil
}

// AssertUnencryptedTicketCredentialConstraints checks if the values respects the defined constraints
func AssertUnencryptedTicketCredentialConstraints(obj UnencryptedTicketCredential) error {
	return nil
}
