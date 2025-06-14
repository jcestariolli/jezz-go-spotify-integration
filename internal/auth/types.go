package auth

import "jezz-go-spotify-integration/internal/model"

type AuthenticationFlow interface {
	Authenticate() (*model.Authentication, error)
}
