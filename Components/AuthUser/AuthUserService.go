package Controller

import (
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	"github.com/mahdidl/golang_boilerplate/dto"
)

type AuthUserService struct {
	AuthUserRepository *AuthUserRepository
}

func NewAuthUserService(authUserRepository *AuthUserRepository) *AuthUserService {
	return &AuthUserService{}
}

func (authUserService AuthUserService) LogoutUser(request dto.LogoutRequest) (response string, err error) {
	payload, _ := token.MakerPaseto.VerifyToken(request.Token)

	err = authUserService.AuthUserRepository.LogOut(request, payload)
	if err != nil {
		return "logout failed", err
	}

	return "logout successfully", err
}
