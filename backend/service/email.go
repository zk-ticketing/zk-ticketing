package service

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/rs/zerolog/log"
)

func (s *APIService) generateEmailSigninCode() (string, error) {
	const codeLength = 6
	const charset = "0123456789"

	code := make([]byte, codeLength)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[num.Int64()]
	}

	return string(code), nil
}

func (s *APIService) sendSigninCodeToEmail(ctx context.Context, email, code string) error {
	logger := log.Ctx(ctx)
	logger.Info().Msgf("Sending signin code %s to email %s", code, email)

	// TODO
	return nil
}
