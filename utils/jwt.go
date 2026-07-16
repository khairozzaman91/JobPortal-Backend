package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub       int    `json:"sub"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Iat       int64  `json:"iat"`
	Exp       int64  `json:"exp"`
}

func CreateJwt(secret string, data Payload) (string, error) {

	if secret == "" {
		return "", errors.New("secret key is required")
	}

	if data.Sub == 0 {
		return "", errors.New("user id is required")
	}

	if data.Email == "" {
		return "", errors.New("email is required")
	}

	data.Iat = time.Now().Unix()
	data.Exp = time.Now().Add(24 * time.Hour).Unix()

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	headerB64 := Base64UrlEncode(headerBytes)
	payloadB64 := Base64UrlEncode(payloadBytes)

	signingInput := headerB64 + "." + payloadB64

	h := hmac.New(sha256.New, []byte(secret))

	_, err = h.Write([]byte(signingInput))
	if err != nil {
		return "", err
	}

	signature := h.Sum(nil)
	signatureB64 := Base64UrlEncode(signature)

	token := signingInput + "." + signatureB64

	return token, nil
}