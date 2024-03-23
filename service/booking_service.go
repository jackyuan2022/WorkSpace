package service

import (
	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	core "github.com/jackyuan2022/workspace/core"
)

type BookingService interface {
	GetBookingList(ctx *gin.Context, r *dto.GetBookingListRequest) (res *dto.DataListResponse[dto.BookingDTO], err *core.AppError)
	CreateBooking(ctx *gin.Context, r *dto.DataRequest[dto.BookingDTO]) (res *dto.DataResponse[dto.BookingDTO], err *core.AppError)
	UpdateBooking(ctx *gin.Context, r *dto.DataRequest[dto.BookingDTO]) (res *dto.DataResponse[dto.BookingDTO], err *core.AppError)
	DeleteBooking(ctx *gin.Context, r *dto.DataRequest[dto.BookingDTO]) (res *dto.DataResponse[dto.BookingDTO], err *core.AppError)
}
