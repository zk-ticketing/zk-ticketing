package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateEmailSigninCode(t *testing.T) {
	apiService := &APIService{}

	code, err := apiService.generateEmailSigninCode()

	// Check that no error occurred
	assert.NoError(t, err)

	// Check that the code length is correct
	assert.Equal(t, 6, len(code))

	// Check that the code contains only digits
	for _, char := range code {
		assert.Contains(t, "0123456789", string(char))
	}
}
