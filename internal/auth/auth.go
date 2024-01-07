package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthIncluded = errors.New("No Authorization header included")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthIncluded
	}

	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) < 2 || authSplit[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return authSplit[1], nil
}
