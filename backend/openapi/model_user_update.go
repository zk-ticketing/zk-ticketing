// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ZK Ticketing API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

package openapi




type UserUpdate struct {

	IdentityCommitment string `json:"identity_commitment,omitempty"`

	EncryptedInternalNullifier string `json:"encrypted_internal_nullifier,omitempty"`

	EncryptedIdentitySecret string `json:"encrypted_identity_secret,omitempty"`
}

// AssertUserUpdateRequired checks if the required fields are not zero-ed
func AssertUserUpdateRequired(obj UserUpdate) error {
	return nil
}

// AssertUserUpdateConstraints checks if the values respects the defined constraints
func AssertUserUpdateConstraints(obj UserUpdate) error {
	return nil
}
