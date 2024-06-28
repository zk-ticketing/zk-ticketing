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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service DefaultAPIServicer
	errorHandler ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) *DefaultAPIController {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() Routes {
	return Routes{
		"EventsEventIdGet": Route{
			strings.ToUpper("Get"),
			"/api/v1/events/{eventId}",
			c.EventsEventIdGet,
		},
		"EventsEventIdRequestTicketCredentialPost": Route{
			strings.ToUpper("Post"),
			"/api/v1/events/{eventId}/request-ticket-credential",
			c.EventsEventIdRequestTicketCredentialPost,
		},
		"EventsGet": Route{
			strings.ToUpper("Get"),
			"/api/v1/events",
			c.EventsGet,
		},
		"HealthGet": Route{
			strings.ToUpper("Get"),
			"/api/v1/health",
			c.HealthGet,
		},
		"UserLoginPost": Route{
			strings.ToUpper("Post"),
			"/api/v1/user/login",
			c.UserLoginPost,
		},
		"UserMeEmailCredentialGet": Route{
			strings.ToUpper("Get"),
			"/api/v1/user/me/email-credential",
			c.UserMeEmailCredentialGet,
		},
		"UserMeEmailCredentialPut": Route{
			strings.ToUpper("Put"),
			"/api/v1/user/me/email-credential",
			c.UserMeEmailCredentialPut,
		},
		"UserMeGet": Route{
			strings.ToUpper("Get"),
			"/api/v1/user/me",
			c.UserMeGet,
		},
		"UserMeRequestEmailCredentialPost": Route{
			strings.ToUpper("Post"),
			"/api/v1/user/me/request-email-credential",
			c.UserMeRequestEmailCredentialPost,
		},
		"UserMeTicketCredentialPut": Route{
			strings.ToUpper("Put"),
			"/api/v1/user/me/ticket-credential",
			c.UserMeTicketCredentialPut,
		},
		"UserMeTicketCredentialsGet": Route{
			strings.ToUpper("Get"),
			"/api/v1/user/me/ticket-credentials",
			c.UserMeTicketCredentialsGet,
		},
		"UserRequestVerificationCodePost": Route{
			strings.ToUpper("Post"),
			"/api/v1/user/request-verification-code",
			c.UserRequestVerificationCodePost,
		},
		"UserUpdatePut": Route{
			strings.ToUpper("Put"),
			"/api/v1/user/update",
			c.UserUpdatePut,
		},
	}
}

// EventsEventIdGet - Get event details
func (c *DefaultAPIController) EventsEventIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventIdParam := params["eventId"]
	if eventIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"eventId"}, nil)
		return
	}
	result, err := c.service.EventsEventIdGet(r.Context(), eventIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// EventsEventIdRequestTicketCredentialPost - Request a new ticket credential for an event
func (c *DefaultAPIController) EventsEventIdRequestTicketCredentialPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventIdParam := params["eventId"]
	if eventIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"eventId"}, nil)
		return
	}
	result, err := c.service.EventsEventIdRequestTicketCredentialPost(r.Context(), eventIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// EventsGet - Get list of events
func (c *DefaultAPIController) EventsGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.EventsGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// HealthGet - Check the health of the API
func (c *DefaultAPIController) HealthGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.HealthGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserLoginPost - User login
func (c *DefaultAPIController) UserLoginPost(w http.ResponseWriter, r *http.Request) {
	userLoginParam := UserLogin{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userLoginParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserLoginRequired(userLoginParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserLoginConstraints(userLoginParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserLoginPost(r.Context(), userLoginParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserMeEmailCredentialGet - Get user email credential
func (c *DefaultAPIController) UserMeEmailCredentialGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserMeEmailCredentialGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserMeEmailCredentialPut - Store user email credential with encrypted data
func (c *DefaultAPIController) UserMeEmailCredentialPut(w http.ResponseWriter, r *http.Request) {
	putEmailCredentialRequestParam := PutEmailCredentialRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&putEmailCredentialRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPutEmailCredentialRequestRequired(putEmailCredentialRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPutEmailCredentialRequestConstraints(putEmailCredentialRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserMeEmailCredentialPut(r.Context(), putEmailCredentialRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserMeGet - Get user details
func (c *DefaultAPIController) UserMeGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserMeGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserMeRequestEmailCredentialPost - Generate a new email credential
func (c *DefaultAPIController) UserMeRequestEmailCredentialPost(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserMeRequestEmailCredentialPost(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserMeTicketCredentialPut - Store user ticket credential with encrypted data
func (c *DefaultAPIController) UserMeTicketCredentialPut(w http.ResponseWriter, r *http.Request) {
	putTicketCredentialRequestParam := PutTicketCredentialRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&putTicketCredentialRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPutTicketCredentialRequestRequired(putTicketCredentialRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPutTicketCredentialRequestConstraints(putTicketCredentialRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserMeTicketCredentialPut(r.Context(), putTicketCredentialRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserMeTicketCredentialsGet - Get user ticket credentials
func (c *DefaultAPIController) UserMeTicketCredentialsGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserMeTicketCredentialsGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserRequestVerificationCodePost - Request an email verification code
func (c *DefaultAPIController) UserRequestVerificationCodePost(w http.ResponseWriter, r *http.Request) {
	userEmailVerificationRequestParam := UserEmailVerificationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userEmailVerificationRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserEmailVerificationRequestRequired(userEmailVerificationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserEmailVerificationRequestConstraints(userEmailVerificationRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserRequestVerificationCodePost(r.Context(), userEmailVerificationRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserUpdatePut - Update user details
func (c *DefaultAPIController) UserUpdatePut(w http.ResponseWriter, r *http.Request) {
	userUpdateParam := UserUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userUpdateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserUpdateRequired(userUpdateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserUpdateConstraints(userUpdateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserUpdatePut(r.Context(), userUpdateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
