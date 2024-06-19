package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Service struct {
	secretKey []byte
	expireSec int64
}

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// NewService creates a new JWT service with the given secret key
func NewService(secretKey string, expireSec int64) *Service {
	return &Service{
		secretKey: []byte(secretKey),
		expireSec: expireSec,
	}
}

func (s *Service) GenerateJWT(id string, email string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(s.expireSec * int64(time.Second)))
	claims := &Claims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *Service) ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
