package credential

import (
  "golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

// passCred is an username/password implementation of credentials.Credentials.
type passCred struct {
	username string
	password string
	secure   bool
}

// GetRequestMetadata implements the required interface function of
// credentials.Credentials.
func (pc *passCred) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": pc.username,
		"password": pc.password,
	}, nil
}

// RequireTransportSecurity implements the required interface function of
// credentials.Credentials.
func (pc *passCred) RequireTransportSecurity() bool {
	return pc.secure
}

// NewPassCred returns a newly created passCred as credentials.Credentials.
func NewPassCred(username, password string, secure bool) credentials.Credentials {
	return &passCred{
		username: username,
		password: password,
		secure:   secure,
	}
}
