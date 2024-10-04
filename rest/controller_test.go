package rest

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/http/cookiejar"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func isBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

const serverAddress = "http://localhost:8080/"

func TestGetTokens(t *testing.T) {
	id := uuid.NewString()

	// Creating tokens
	client := http.Client{}
	resp, err := client.Get(serverAddress + id)
	if err != nil {
		t.Fatalf("HTTP Request error:\n%v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Status code is not OK: %d", resp.StatusCode)
	}

	// Checking cookies
	t.Log("Received with cookies:")
	for _, cookie := range resp.Cookies() {
		t.Logf("%s = %s\n", cookie.Name, cookie.Value)
		if cookie.Name == "medods_refresh_token" {
			t.Logf("Refresh Token is in Base64 format? - %v\n", isBase64(cookie.Value))
		}
	}
}

func getCookie(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}

func TestPatchTokens(t *testing.T) {
	id := uuid.NewString()

	// Creating tokens
	jar, _ := cookiejar.New(nil)

	client := http.Client{
		Jar: jar,
	}
	resp, err := client.Get("http://localhost:8080/" + id)
	if err != nil {
		t.Fatalf("HTTP Request error:\n%v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Status code is not OK: %d", resp.StatusCode)
	}

	// Get current cookies
	refreshTokenBefore := getCookie(resp.Cookies(), "medods_refresh_token")
	accessTokenBefore := getCookie(resp.Cookies(), "medods_access_token")

	// Patch
	req, _ := http.NewRequest("PATCH", serverAddress, bytes.NewBufferString(""))
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("HTTP Request error:\n%v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Status code is not OK: %d", resp.StatusCode)
	}

	refreshTokenAfter := getCookie(resp.Cookies(), "medods_refresh_token")
	accessTokenAfter := getCookie(resp.Cookies(), "medods_access_token")

	assert.NotEqual(t, accessTokenBefore, accessTokenAfter)
	assert.NotEqual(t, refreshTokenBefore, refreshTokenAfter)
}
