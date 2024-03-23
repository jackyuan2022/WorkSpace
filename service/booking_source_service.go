package service

import (
	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	core "github.com/jackyuan2022/workspace/core"
)

type BookingSourceService interface {
	GetBookingSourceList(ctx *gin.Context, r *dto.GetBookingSourceListRequest) (res *dto.DataListResponse[dto.BookingSourceDTO], err *core.AppError)
	CreateBookingSource(ctx *gin.Context, r *dto.DataRequest[dto.BookingSourceDTO]) (res *dto.DataResponse[dto.BookingSourceDTO], err *core.AppError)
	UpdateBookingSource(ctx *gin.Context, r *dto.DataRequest[dto.BookingSourceDTO]) (res *dto.DataResponse[dto.BookingSourceDTO], err *core.AppError)
	DeleteBookingSource(ctx *gin.Context, r *dto.DataRequest[dto.BookingSourceDTO]) (res *dto.DataResponse[dto.BookingSourceDTO], err *core.AppError)
}
