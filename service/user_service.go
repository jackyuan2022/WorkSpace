package service

import (
	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	core "github.com/jackyuan2022/workspace/core"
)

type UserService interface {
	Login(ctx *gin.Context, u *dto.LoginRequest) (r *dto.LoginResponse, err *core.AppError)
	Register(ctx *gin.Context, u *dto.RegisterRequest) *core.AppError
	RefreshToken(ctx *gin.Context, u *dto.RefreshTokenRequest) (r *dto.RefreshTokenResponse, err *core.AppError)
}
