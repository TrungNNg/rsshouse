package api

import "errors"

var ErrPasswordTooShort = errors.New("password too short")
var ErrPasswordTooLong = errors.New("password too long")
var ErrUsernameTooShort = errors.New("username too short")
var ErrUsernameTooLong = errors.New("username too long")
