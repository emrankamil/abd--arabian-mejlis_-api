package domain

import (
	"context"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ForgetPWRequest struct{
	Email  string	`json:"email" validate:"email,required"`
}

type ResetPWRequest struct{
	Email    string	`form:"email" binding:"required,email"`
	Password string	`json:"password" validate:"required,min=8,max=32"`
}

type LoginUsecase interface {
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	LogoutUser(c context.Context, email string) error
}

type ForgetPWUsecase interface {
	ForgetPW(c context.Context, email string, server_address string) error
	ResetPW(c context.Context, request ResetPWRequest) error
	VerifyForgetPWRequest(c context.Context, email string, recover_token string) error
	GenerateRecoveryLink(server_address string, email string, recoveryToken string) (recoveryLink string)
}