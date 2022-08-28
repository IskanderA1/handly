package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeys = 32

type JwtMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeys {
		return nil, fmt.Errorf("invalid key size: min size must be %d", minSecretKeys)
	}
	return &JwtMaker{secretKey}, nil
}

func (jm *JwtMaker) CreateToken(projectName string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(projectName, duration)
	if err != nil {
		return "", nil, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(jm.secretKey))
	return token, payload, err
}

func (jm *JwtMaker) VerifyToken(token string) (*Payload, error) {
	/// Проверка на выборк секретного ключа для расшифровки токена
	/// Тут проверяется то что метод шифровния сооствстует SigningMethodHMAC,
	/// т.к. SigningMethodHS256 является реализацией SigningMethodHMAC
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jm.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
