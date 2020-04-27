package db

import "errors"

var (
	// ErrNotFound возвращается, если пользователь не зарегистрирован.
	ErrNotFound = errors.New("not registered")
	// ErrBlocked возвращается, если пользователь заблокирован в системе.
	ErrBlocked = errors.New("blocked")
	// ErrAlreadyRegisterd возвращается, если пользователь уже зарегистрирован.
	ErrAlreadyRegisterd = errors.New("already registered")
	// ErrInvalidPassword возвращается в случае неверного паролья пользователя.
	ErrInvalidPassword = errors.New("invalid password")
	// ErrBadToken возвращается, если токен на зарегистрирован.
	ErrBadToken = errors.New("bad token")
	// ErrEmptyEmail возвращается, если email адрес пустой.
	ErrEmptyEmail = errors.New("empty email")
)
