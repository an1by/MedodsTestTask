package util

import (
	b64 "encoding/base64"

	"github.com/google/uuid"
)

func GenerateBase64UUID() string {
	return b64.StdEncoding.EncodeToString([]byte(uuid.NewString()))
}