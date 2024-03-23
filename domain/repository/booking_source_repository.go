package repository

import (
	"context"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
)

type BookingSourceRepository interface {
	FindById(ctx context.Context, id string) (*model.BookingSource, *core.AppError)
	Insert(ctx context.Context, data *model.BookingSource) (*model.BookingSource, *core.AppError)
	Update(ctx context.Context, data *model.BookingSource) (*model.BookingSource, *core.AppError)
	DeleteById(ctx context.Context, id string) (bool, *core.AppError)
	QueryData(ctx context.Context, query *core.DbQuery) ([]model.BookingSource, *core.AppError)
}
