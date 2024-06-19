package service

import (
	"context"
	"net/http"
	"net/mail"

	"github.com/NFTGalaxy/zk-ticketing-server/jwt"
	"github.com/NFTGalaxy/zk-ticketing-server/openapi"
	"github.com/NFTGalaxy/zk-ticketing-server/repos"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/users"
	"github.com/NFTGalaxy/zk-ticketing-server/util"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

type APIService struct {
	dbClient   *repos.Client
	jwtService *jwt.Service
}

// NewAPIService creates a default api service
func NewAPIService(
	dbClient *repos.Client,
	jwtService *jwt.Service,
) *APIService {
	return &APIService{
		dbClient:   dbClient,
		jwtService: jwtService,
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

// UsersLoginPost - User login
func (s *APIService) UsersLoginPost(ctx context.Context, userLogin openapi.UserLogin) (openapi.ImplResponse, error) {
	logger := log.Ctx(ctx).With().Str("op", "UsersLoginPost").Logger()

	// TODO: handle user login code - now it's useful to test the API
	email, err := mail.ParseAddress(userLogin.Email)
	if err != nil {
		return openapi.Response(http.StatusBadRequest, "Invalid Email"), nil
	}
	logger = logger.With().Str("email", email.Address).Logger()

	u, err := s.dbClient.Users.GetUserByEmail(ctx, email.Address)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Info().Msg("User not found, creating new user")
			u, err = s.dbClient.Users.CreateUser(ctx, users.CreateUserParams{
				ID:    uuid.NewString(),
				Email: email.Address,
			})
			if err != nil {
				return openapi.Response(http.StatusInternalServerError, nil), err
			}
		} else {
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	// Generate JWT
	token, err := s.jwtService.GenerateJWT(u.ID, u.Email)
	if err != nil {
		logger.Err(err).Msg("Failed to generate JWT")
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	logger.Info().Msg("User logged in")
	return openapi.Response(http.StatusOK, openapi.LoginResponse{Token: token}), nil
}

// UsersMeGet - Get user details
func (s *APIService) UsersMeGet(ctx context.Context) (openapi.ImplResponse, error) {
	userEmail := util.GetUserEmailFromContext(ctx)
	userID := util.GetUserIDFromContext(ctx)
	log.Info().Msgf("userID: %v, userEmail: %v", userID, userEmail)
	if userID == "" || userEmail == "" {
		return openapi.Response(http.StatusUnauthorized, nil), nil
	}

	user, err := s.dbClient.Users.GetUserByID(ctx, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return openapi.Response(http.StatusNotFound, nil), nil
		}
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, MarshalUser(user)), nil
}

// UsersRequestVerificationCodePost - Request an email verification code
func (s *APIService) UsersRequestVerificationCodePost(ctx context.Context, userEmailVerificationRequest openapi.UserEmailVerificationRequest) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
