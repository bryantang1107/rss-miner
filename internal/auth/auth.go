package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts API Key from headers of HTTP request
// Example:
// Authorization: ApiKey <insert apikey here>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No API Key Found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth header")
	}

	return vals[1], nil
}
