package utils

type CustomError interface {
	error
	GetCode() int
	GetMsg() string
}
