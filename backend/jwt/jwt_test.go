package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

const (
	testSecretKey = "mysecretkey"
	testExpireSec = 3600
	testUserID    = "12345"
	testUserEmail = "user@example.com"
)

func TestGenerateJWT(t *testing.T) {
	jwtService := NewService(testSecretKey, testExpireSec)

	tokenString, err := jwtService.GenerateJWT(testUserID, testUserEmail)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Decode token to verify the claims
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(testSecretKey), nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)
	assert.Equal(t, testUserID, claims.ID)
	assert.Equal(t, testUserEmail, claims.Email)

	expectedExpiration := time.Now().Add(1 * time.Hour).Unix()
	assert.InDelta(t, expectedExpiration, claims.ExpiresAt, 60)
}

func TestValidateJWT(t *testing.T) {
	jwtService := NewService(testSecretKey, testExpireSec)

	// Generate a token to validate
	tokenString, err := jwtService.GenerateJWT(testUserID, testUserEmail)
	assert.NoError(t, err)

	claims, err := jwtService.ValidateJWT(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, testUserID, claims.ID)
	assert.Equal(t, testUserEmail, claims.Email)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	jwtService := NewService(testSecretKey, testExpireSec)

	invalidToken := "invalid.token.string"

	_, err := jwtService.ValidateJWT(invalidToken)
	assert.Error(t, err)
}
