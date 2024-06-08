package middleware

import (
	"net/http"
	"net/mail"
)

// ParseUserEmailFromAuthHeader parses the user email from the Authorization header
func ParseUserEmailFromAuthHeader(r *http.Request) (string, error) {
	email := r.Header.Get("Authorization")
	if email == "" {
		return "", http.ErrNoCookie
	}

	// Validate email format
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", err
	}

	return email, nil
}
