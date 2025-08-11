package dto

import (
	"time"

	permission "github.com/mahdidl/golang_boilerplate/Components/Permission/Entity"
	role "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	user "github.com/mahdidl/golang_boilerplate/Components/User/Entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePermissionReq struct {
	Name string `json:"Name" validate:"required,min=3"`
}

type CreatePermissionRes struct {
	ID   primitive.ObjectID `json:"Id" bson:"_id"`
	Name string             `json:"Name" bson:"Name"`
}

type GetAllPermissionsReq struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

type GetPermissionsRes struct {
	Permissions []permission.Permission `json:"Permissions"`
}

type Permission struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"Name"`
}

type AccessTokenRequest struct {
	AccessToken string `json:"refreshToken"  validate:"required"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"  validate:"required"`
}

type LogoutRequest struct {
	Token string `json:"token" validate:"required,min=3"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `bson:"Password" validate:"required,min=8"  json:"currentPassword"`
	NewPassword     string `bson:"Password" validate:"required,min=8" json:"newPassword"`
}

type ChangeStatusRequest struct {
	ID string `json:"id" form:"id"`
}

type CreateUserRequest struct {
	// username of the user
	// in: string
	UserName string `json:"username" validate:"required,min=3"`
	// password of the user
	// in: string
	Password string `json:"password" validate:"required,min=8"`
}

type GetAllUsers struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

type GetUser struct {
	ID string `json:"userId" form:"userId"`
}

type GetUserRequest struct {
	UserName string `json:"username" validate:"required,min=3"`
}

type LoginUserRequest struct {
	UserName string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	UserName string `bson:"UserName" json:"userName"`
}
type CreateUserResponse struct {
	UserName string `json:"username" binding:"required"`
}
type GetUserByIdResponse struct {
	UserId   primitive.ObjectID `json:"subject" validate:"required"`
	UserName string             `json:"username" validate:"required,min=3"`
}

type GetUserResponse struct {
	UserId   primitive.ObjectID `json:"subject" validate:"required"`
	UserName string             `json:"username" validate:"required,min=3"`
}

type LoginUserResponse struct {
	Id           string `json:"id" `
	UserName     string `json:"username" `
	AccessToken  string `json:"accessToken" `
	RefreshToken string `json:"refreshToken" `
}

type ResponseAllUsers struct {
	Users []user.User `json:"users" `
}

type AttachRoleReq struct {
	RoleId string `json:"roleId" form:"roleId" validate:"required,min=3"`
	UserId string `json:"userId" form:"userId" validate:"required,min=3"`
}

type DetachRoleReq struct {
	UserId string `json:"userId" form:"userId" validate:"required,min=3"`
}

type AttachRoleRes struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"UserName"`
	IsActive  bool               `bson:"IsActive" `
	RoleID    primitive.ObjectID `bson:"RoleId"`
	CreatedAt time.Time          `bson:"CreatedAt"`
	UpdatedAt *time.Time         `bson:"UpdatedAt"`
	DeletedAt *time.Time         `bson:"DeletedAt"`
}

type DetachRoleRes struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"UserName"`
	IsActive  bool               `bson:"IsActive" `
	RoleID    primitive.ObjectID `bson:"RoleId"`
	CreatedAt time.Time          `bson:"CreatedAt"`
	UpdatedAt *time.Time         `bson:"UpdatedAt"`
	DeletedAt *time.Time         `bson:"DeletedAt"`
}

type CreateTicketRequest struct {
	Subject string `json:"subject" validate:"required"`
	Message string `json:"message" validate:"required"`
	Like    bool   `json:"like" validate:"required"`
	Image   string `json:"image"`
}

type CreateTicketResponse struct {
	Subject string `json:"subject" validate:"required"`
	Message string `json:"message" validate:"required"`
	Like    bool   `json:"like" validate:"required"`
	Image   string `json:"image"`
}

type AttachPermissionReq struct {
	RoleId       string `json:"roleId" form:"roleId" validate:"required,min=3"`
	PermissionId string `json:"permissionId" form:"permissionId" validate:"required,min=3"`
}

type DetachPermissionReq struct {
	RoleId       string `json:"roleId" form:"roleId" validate:"required,min=3"`
	PermissionId string `json:"permissionId" form:"permissionId" validate:"required,min=3"`
}

type AttachRes struct {
	Id            primitive.ObjectID   `bson:"_id"`
	PermissionsId []primitive.ObjectID `bson:"Permissions"`
	Name          string               `bson:"Name"`
	CreatedAt     time.Time            `bson:"CreatedAt"`
	UpdatedAt     *time.Time           `bson:"UpdatedAt"`
	DeletedAt     *time.Time           `bson:"DeletedAt"`
}

type DetachRes struct {
	Id            primitive.ObjectID   `bson:"_id"`
	PermissionsId []primitive.ObjectID `bson:"Permissions"`
	Name          string               `bson:"Name"`
	CreatedAt     time.Time            `bson:"CreatedAt"`
	UpdatedAt     *time.Time           `bson:"UpdatedAt"`
	DeletedAt     *time.Time           `bson:"DeletedAt"`
}

type CreateRoleReq struct {
	Name string `json:"Name" validate:"required,min=3"`
}

type DeleteRoleReq struct {
	Id string `json:"Id" form:"Id" validate:"required,min=3"`
}

type GetAllRoleReq struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

type GetRoleReq struct {
	Id string `json:"roleId" form:"roleId" validate:"required,min=3"`
}

type UpdateRoleReq struct {
	Name string `json:"Name" bson:"Name"  validate:"required"`
}

type CreateRoleRes struct {
	ID   primitive.ObjectID `json:"Id" bson:"_id"`
	Name string             `json:"Name" bson:"Name"  validate:"required"`
}

type DeleteRoleRes struct {
	Id string `json:"Id" form:"Id" validate:"required,min=3"`
}

type GetAllRolesRes struct {
	Roles []role.Role `json:"Role"`
}

type GetRoleRes struct {
	Roles role.Role `json:"Role"`
}

type UpdateRoleRes struct {
	Id   string `json:"Id" validate:"required,min=3"`
	Name string `json:"Name" bson:"Name"  validate:"required"`
}
