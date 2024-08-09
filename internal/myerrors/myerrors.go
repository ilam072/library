package myerrors

import "errors"

var (
	ErrUsernameExists    = errors.New("username already exists")
	ErrInvalidUserData   = errors.New("invalid password or username")
	ErrInvalidBookAccess = errors.New("invalid book access")
	ErrBooksNotFound     = errors.New("books not found")
	ErrBookNotFound      = errors.New("book not found")
	ErrInvalidBookId     = errors.New("invalid book id: must be integer")
	ErrBookFileNotFound  = errors.New("file not found")
)
