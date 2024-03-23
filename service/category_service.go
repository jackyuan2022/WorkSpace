package service

import (
	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	core "github.com/jackyuan2022/workspace/core"
)

type CategoryService interface {
	GetCategoryList(ctx *gin.Context, r *dto.GetCategoryListRequest) (res *dto.DataListResponse[dto.CategoryDTO], err *core.AppError)
	CreateCategory(ctx *gin.Context, r *dto.DataRequest[dto.CategoryDTO]) (res *dto.DataResponse[dto.CategoryDTO], err *core.AppError)
	UpdateCategory(ctx *gin.Context, r *dto.DataRequest[dto.CategoryDTO]) (res *dto.DataResponse[dto.CategoryDTO], err *core.AppError)
	DeleteCategory(ctx *gin.Context, r *dto.DataRequest[dto.CategoryDTO]) (res *dto.DataResponse[dto.CategoryDTO], err *core.AppError)
}
