package repository

import (
	"context"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
)

type BookingRepository interface {
	FindById(ctx context.Context, id string) (*model.Booking, *core.AppError)
	Insert(ctx context.Context, data *model.Booking) (*model.Booking, *core.AppError)
	Update(ctx context.Context, data *model.Booking) (*model.Booking, *core.AppError)
	DeleteById(ctx context.Context, id string) (bool, *core.AppError)
	QueryData(ctx context.Context, query *core.DbQuery) ([]model.Booking, *core.AppError)
}
