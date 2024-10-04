package database

import (
	"aniby/medods/util"
	
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id           string		`pg:",unique"`
	RefreshToken string		`pg:",unique"`
}

func (user User) CheckHash(refreshToken string) error {
	return util.CompareHash(user.RefreshToken, refreshToken)
}

func (user *User) Refresh() (string, error) {
	refreshToken := util.GenerateBase64UUID()
	encodedRefreshToken, err := util.HashToken(refreshToken)
	if err != nil {
		return "", err
	}
	user.RefreshToken = encodedRefreshToken
	return refreshToken, nil
}

func (user User) GenerateAccessToken(address string, expiresAt int64, secretKey string) (string, error) {
	payload := jwt.MapClaims{
		"id": user.Id,
		"address": address,
		"exp": expiresAt,
	}
	return util.GenerateSignedJWTString(payload, secretKey)
}

func CreateUser(id string) (User, string, error) {
	refreshToken := util.GenerateBase64UUID()
	encodedRefreshToken, err := util.HashToken(refreshToken)
	if err != nil {
		return User{}, "", err
	}
	return User{id, encodedRefreshToken}, refreshToken, nil
}