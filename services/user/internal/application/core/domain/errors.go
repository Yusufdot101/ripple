package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidToken      = errors.New("invalid token")
	ErrNoIDToken         = errors.New("no id_token field in oauth2 token")
	ErrInvalidNonce      = errors.New("invalid nonce in id_token")
	ErrInvalidTokenUse   = errors.New("invalid token use")
	ErrInvalidTokeType   = errors.New("invalid token type")
)
