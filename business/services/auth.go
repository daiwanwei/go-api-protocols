package services

import (
	"context"
	"go-api-protocols/utils/security"
)

type AuthService interface {
	Login(ctx context.Context, loginDto LoginDto) (passportDto *PassportDto, err error)
	FindAuthByName(ctx context.Context, name string) (security.Authentication, error)
}

type authService struct {
}

func NewAuthService() (AuthService, error) {
	return &authService{}, nil
}

func (service *authService) Login(ctx context.Context, loginDto LoginDto) (passportDto *PassportDto, err error) {
	auth := AuthDto{
		UserID:      "0",
		Authorities: []string{"user"},
	}

	token, err := security.GenerateToken(auth)
	if err != nil {
		return
	}

	passportDto = &PassportDto{
		Token:       token,
		UserID:      "0",
		Authorities: auth.Authorities,
	}
	return
}

func (service *authService) FindAuthByName(ctx context.Context, name string) (security.Authentication, error) {
	return &AuthDto{
		UserID:      name,
		Authorities: []string{"user"},
	}, nil
}

type AuthDto struct {
	UserID      string   `json:"userId"`
	Authorities []string `json:"authorities"`
}

func (dto AuthDto) GetName() string {
	return dto.UserID
}

func (dto AuthDto) GetAuthorities() []string {
	return dto.Authorities
}

type PassportDto struct {
	Token       string   `json:"token"`
	UserID      string   `json:"userId"`
	Authorities []string `json:"authorities"`
}

type LoginDto struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AuthServiceError struct {
	ServiceError
}

func NewAuthServiceError(e ServiceEvent) error {
	return &AuthServiceError{ServiceError{ServiceName: "wrapUserService", Code: e.GetEvent().Code, Msg: e.GetEvent().Msg, Err: nil}}
}
