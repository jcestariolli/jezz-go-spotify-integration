package auth

const (
	BasicAuthorizationType  AuthorizationType = "Basic"
	BearerAuthorizationType AuthorizationType = "Bearer"
)

type (
	AuthorizationType   string
	AuthorizationHeader struct {
		Key   string
		Value string
	}
	AccessToken string
)

func (authType AuthorizationType) String() string {
	return string(authType)
}

func (accessToken AccessToken) String() string {
	return string(accessToken)
}
