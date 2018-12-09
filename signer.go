package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Information about the signing process can be found on Coinbase's official
// documentation: https://docs.pro.coinbase.com/#signing-a-message
func generateSignature(body, timestamp, secret string) (string, error) {
	preHash := fmt.Sprintf("%s%s%s%s", timestamp, "POST", "/orders", body)

	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	signature := hmac.New(sha256.New, decodedSecret)

	if _, err = signature.Write([]byte(preHash)); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature.Sum(nil)), nil
}
