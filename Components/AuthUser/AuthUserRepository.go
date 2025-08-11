package Controller

import (
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	"github.com/mahdidl/golang_boilerplate/dto"
)

type AuthUserRepository struct {
}

func NewAuthUserRepository() *AuthUserRepository {
	return &AuthUserRepository{}
}

func (userRepository *AuthUserRepository) LogOut(logoutReq dto.LogoutRequest, payload *token.Payload) error {

	err := Config.Redis.Set(payload.Username, logoutReq.Token, 0).Err()

	if err != nil {
		return nil
	}
	return err
}
