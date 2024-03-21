package service

import (
	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	core "github.com/jackyuan2022/workspace/core"
)

type BookingService interface {
	GetBookingList(ctx *gin.Context, r *dto.GetBookingListRequest) (res *dto.GetBookingListResponse, err *core.AppError)
	CreateBooking(ctx *gin.Context, r *dto.BookingRequest) (res *dto.BookingResponse, err *core.AppError)
	UpdateBooking(ctx *gin.Context, r *dto.BookingRequest) (res *dto.BookingResponse, err *core.AppError)
	DeleteBooking(ctx *gin.Context, r *dto.BookingRequest) (res *dto.BookingResponse, err *core.AppError)
}
