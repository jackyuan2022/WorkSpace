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
	GetUserList(ctx *gin.Context, r *dto.GetUserListRequest) (res *dto.DataListResponse[dto.UserDTO], err *core.AppError)
	UpdateUser(ctx *gin.Context, r *dto.DataRequest[dto.UserDTO]) (res *dto.DataResponse[dto.UserDTO], err *core.AppError)
	DeleteUser(ctx *gin.Context, r *dto.DataRequest[dto.UserDTO]) (res *dto.DataResponse[dto.UserDTO], err *core.AppError)
}
