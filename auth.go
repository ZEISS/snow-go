package snowgo

import (
	"context"
	"net/http"
)

// BasicAuth implements the SecurityProvider interface for basic authentication.
type BasicAuth struct {
	username string
	password string
}

// NewBasicAuth returns a new BasicAuth security provider.
func NewBasicAuth(username, password string) (*BasicAuth, error) {
	return &BasicAuth{
		username: username,
		password: password,
	}, nil
}

// Intercept will attach an Authorization header to the request and ensures that
func (s *BasicAuth) Intercept(ctx context.Context, req *http.Request) error {
	req.SetBasicAuth(s.username, s.password)

	return nil
}
