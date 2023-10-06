package globalerrors

import "errors"

var (
	UserEmailExists     = errors.New("email already exists")
	UserUsernameExists  = errors.New("username already exists")
	UserInvalidUsername = errors.New("invalid username")
	UserCouldNotQuery   = errors.New("could not query")

	AuthInvalidToken            = errors.New("invalid token")
	AuthInvalidCredentials      = errors.New("invalid credentials")
	AuthMissingAuthHeader       = errors.New("missing Authorization header")
	AuthInvalidAuthHeaderFormat = errors.New("Invalid Authorization Header Format")

	GlobalInternalServerError = errors.New("internal server error")
)