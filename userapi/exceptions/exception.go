package exceptions

import "errors"

var (
	UserNotFound = errors.New("user_not_found")
)
