package util

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateSignedJWTString(payload jwt.MapClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES512, payload)
	return token.SignedString(secretKey)
}

func DecodeSignedJWTString(jwtSignedString string, secretKey string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtSignedString, claims, func(token *jwt.Token) (interface{}, error) {
	    return []byte(secretKey), nil
	})
	return claims, err
}