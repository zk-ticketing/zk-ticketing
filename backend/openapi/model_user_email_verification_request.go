// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ZK Ticketing API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

package openapi




type UserEmailVerificationRequest struct {

	Email string `json:"email,omitempty"`
}

// AssertUserEmailVerificationRequestRequired checks if the required fields are not zero-ed
func AssertUserEmailVerificationRequestRequired(obj UserEmailVerificationRequest) error {
	return nil
}

// AssertUserEmailVerificationRequestConstraints checks if the values respects the defined constraints
func AssertUserEmailVerificationRequestConstraints(obj UserEmailVerificationRequest) error {
	return nil
}
