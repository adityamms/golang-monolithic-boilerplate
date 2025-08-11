package Controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Common/Middleware"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	"github.com/mahdidl/golang_boilerplate/dto"

	"net/http"
)

type AuthController struct {
	authService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func RegisterRoutes(router *gin.RouterGroup) {

	authenticationPostfix := "/authentication"
	authRepository := NewAuthRepository()
	authService := NewAuthService(authRepository)
	authController := NewAuthController(authService)
	authUserRouter := router.Group(authenticationPostfix).Use(Middleware.AuthMiddleware())
	authLoginRouter := router.Group(authenticationPostfix).Use()
	{
		authUserRouter.POST("/newToken", authController.AccessToken)
		authLoginRouter.POST("/login", authController.LoginUser)
		authUserRouter.DELETE("/logout", authController.Logout)
	}

}

// @Summary      New access token
// @Description  New access token with refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        AccessTokenRequest  body      AuthRequest.AccessTokenRequest  true  "for get new access token"
// @Success      200                {object}  Response.GeneralResponse{data=User.AccessTokenRequest}
// @Failure      400                {object}  Response.GeneralResponse{data=object} ""
// @Router       /authentication/newToken [post]
//
// LoginUser for get access token
func (authController *AuthController) AccessToken(context *gin.Context) {
	var accessTokenReq dto.AccessTokenRequest
	Helper.Decode(context.Request, &accessTokenReq)

	validationError := Validator.ValidationCheck(accessTokenReq)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	token, err := authController.authService.CreateAccessToken(accessTokenReq)
	if err != nil {
		response := Response.GeneralResponse{Error: true, Message: err.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: "successful", Data: dto.AccessTokenResponse{AccessToken: token}}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}

// LoginUser
// @Summary      Login user
// @Description  Login user with username and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        LoginUserRequest  body      Request.CreateUserRequest  true  "Create user request"
// @Success      200                {object}  Response.GeneralResponse{data=UserResponse.LoginUserResponse}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user not exist or password is incorrect"
// @Router       /authentication/login [post]
//
// LoginUser for get access token
func (authController *AuthController) LoginUser(context *gin.Context) {
	var userRequest dto.LoginUserRequest
	context.ShouldBindJSON(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	userResponse, responseError := authController.authService.LoginUser(userRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	var loginResponse dto.LoginUserResponse
	loginResponse = userResponse
	response := Response.GeneralResponse{Error: false, Message: "your login is successful", Data: loginResponse}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// LogoutUser
// @Summary      Logout user
// @Description  Logout user with access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        LogoutUserRequest  body      Request.LogoutRequest  true  "logout user"
// @Success      200                {object}  Response.GeneralResponse{data=string}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when access token is not valid"
// @Router       /authentication/logout [delete]
//
// Logout user with access token
func (authController *AuthController) Logout(context *gin.Context) {
	var userRequest dto.LogoutRequest
	Helper.Decode(context.Request, &userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	logoutResponse, logoutResponseError := authController.authService.LogoutUser(userRequest)

	if logoutResponseError != nil {
		response := Response.GeneralResponse{Error: true, Message: logoutResponseError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: logoutResponse}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
