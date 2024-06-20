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



type UnencryptedEmailCredential struct {

	Id string `json:"id,omitempty"`

	Credential string `json:"credential,omitempty"`

	IssuedAt time.Time `json:"issued_at,omitempty"`

	ExpireAt time.Time `json:"expire_at,omitempty"`
}

// AssertUnencryptedEmailCredentialRequired checks if the required fields are not zero-ed
func AssertUnencryptedEmailCredentialRequired(obj UnencryptedEmailCredential) error {
	return nil
}

// AssertUnencryptedEmailCredentialConstraints checks if the values respects the defined constraints
func AssertUnencryptedEmailCredentialConstraints(obj UnencryptedEmailCredential) error {
	return nil
}
