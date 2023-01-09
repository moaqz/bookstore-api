package jwtToken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/techwithmat/bookery-api/pkg/utils/env"
)

type jwtCustomClaims struct {
	Email   string `json:"email"`
	IsStaff bool   `json:"is_staff"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, isStaff bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwtCustomClaims{
		Email:   email,
		IsStaff: isStaff,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	secret := []byte(env.MustGet("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(env.MustGet("JWT_SECRET"))

	// Validate the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
