package database

import (
	"github.com/golang-jwt/jwt/v5"
	"aniby/medods/util"
)

type User struct {
	Id           string		`pg:",unique"`
	RefreshToken string		`pg:",unique"`
}

func (user User) CheckHash(refreshToken string) error {
	return util.CompareHash(user.RefreshToken, refreshToken)
}

func (user User) Refresh() string {
	user.RefreshToken = util.GenerateBase64UUID()
	return user.RefreshToken
}

func (user User) GenerateAccessToken(address string, secretKey string) (string, error) {
	payload := jwt.MapClaims{
		"id": user.Id,
		"address": address,
	}
	return util.GenerateSignedJWTString(payload, secretKey)
}

func CreateUser(id string) (User, error) {
	refreshToken, err := util.HashToken(util.GenerateBase64UUID())
	if err != nil {
		return User{}, err
	}
	return User{id, refreshToken}, nil
}