package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey Extracts an API key from the headers of an HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")

	if value == "" {
		return "", errors.New("not auth info found")
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		return "", errors.New("invalid auth info")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("invalid auth info. Use ApiKey {your key}")
	}

	return values[1], nil

}
