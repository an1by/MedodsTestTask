package rest

import (
	"aniby/medods/database"
	"aniby/medods/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetTokensResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var JwtSercetKey string

func setCookieTokens(writer gin.ResponseWriter, accessToken string, refreshToken string) {
	http.SetCookie(writer, &http.Cookie{
		Name: "medods_access_token",
		Value: accessToken,
		Path: "/",
		HttpOnly: true,
		MaxAge: 0,
	})
	http.SetCookie(writer, &http.Cookie{
		Name: "medods_refresh_token",
		Value: refreshToken,
		Path: "/",
		HttpOnly: true,
		MaxAge: 0,
	})
}

func GetTokens(c *gin.Context) {
	id := c.Param(":id")
	
	user, err := database.CreateUser(id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	_, err = database.InsertUser(user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	accessToken, err := user.GenerateAccessToken(c.ClientIP(), JwtSercetKey)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	
	setCookieTokens(c.Writer, accessToken, user.RefreshToken)
	c.Status(http.StatusOK)
}

func PatchTokens(c *gin.Context) {
	// Access Token processing
	accessToken, err := c.Cookie("medods_access_token")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	claims, err := util.DecodeSignedJWTString(accessToken, JwtSercetKey)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}

	address, ok := claims["address"].(string)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}
	id, ok := claims["id"].(string)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	// Refresh Token processing
	refreshToken, err := c.Cookie("medods_refresh_token")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := database.FindUserById(id)
	if err != nil || user.CheckHash(refreshToken) != nil {
		c.Status(http.StatusForbidden)
		return
	}

	// Address notify
	if address != c.ClientIP() {
		util.SendMailMock()	
	}

	// Refresh
	accessToken, err = user.GenerateAccessToken(address, JwtSercetKey)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user.Refresh()
	setCookieTokens(c.Writer, accessToken, user.RefreshToken)
	c.Status(http.StatusOK)
}