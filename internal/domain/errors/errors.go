package globalerrors

import "errors"

var (
	UserEmailExists = errors.New("email already exists")
	UserUsernameExists = errors.New("username already exists")
	
)
