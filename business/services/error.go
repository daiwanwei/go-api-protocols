package services

import "fmt"

type ServiceError struct {
	ServiceName string
	Code        int
	Msg         string
	Err         error
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("code(%d): Msg(%s)", e.Code, e.ServiceName+":"+e.Msg)
}

func (e ServiceError) GetCode() int {
	return e.Code
}

func (e ServiceError) GetMsg() string {
	return e.Msg
}

type ServiceEvent int

type Event struct {
	Code int
	Msg  string
}

const (
	UserNotFound           ServiceEvent = 201
	UserRegistered         ServiceEvent = 202
	UserNameBeenRegistered ServiceEvent = 203
	PasswordWrong          ServiceEvent = 204
)

func (e ServiceEvent) GetEvent() *Event {
	switch e {
	case UserNotFound:
		return &Event{int(e), "users not found"}
	case UserRegistered:
		return &Event{int(e), "users has registered"}
	case UserNameBeenRegistered:
		return &Event{int(e), "users name have been registered"}
	case PasswordWrong:
		return &Event{int(e), "password is wrong"}
	default:
		return &Event{int(e), "unknown"}
	}
}
