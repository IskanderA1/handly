package token

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeys = 32

type JwtMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker[ProjectPayload, ProjectPayloadInput], error) {
	if len(secretKey) < minSecretKeys {
		return nil, fmt.Errorf("invalid key size: min size must be %d", minSecretKeys)
	}
	return &JwtMaker{secretKey}, nil
}

func (jm *JwtMaker) CreateToken(inp ProjectPayloadInput) (string, *ProjectPayload, error) {
	payload, err := NewProjectPayload(inp)
	if err != nil {
		return "", nil, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(jm.secretKey))
	return token, payload, err
}

func (jm *JwtMaker) VerifyToken(token string) (*ProjectPayload, error) {

	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jm.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &ProjectPayload{}, keyfunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*ProjectPayload)

	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
