package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSignature_Success(t *testing.T) {
	body := `{"foo": "bar"}`
	timestamp := "1544281725.56"
	secret := "ABCD1234"

	signature, err := generateSignature(body, timestamp, secret)

	assert.NoError(t, err)
	assert.Equal(t, signature, "kqKSisBO0iUFuyPdbCgsnRelHo/A3LbxsOqE02/v3Hk=")
}

func TestGenerateSignature_InvalidSecret(t *testing.T) {
	body := `{"foo": "bar"}`
	timestamp := "1544281725.56"
	secret := "!#$"

	signature, err := generateSignature(body, timestamp, secret)

	assert.Error(t, err)
	assert.Equal(t, signature, "")
}
