package repository

import (
	"context"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
)

type CategoryRepository interface {
	FindById(ctx context.Context, id string) (*model.Category, *core.AppError)
	Insert(ctx context.Context, data *model.Category) (*model.Category, *core.AppError)
	Update(ctx context.Context, data *model.Category) (*model.Category, *core.AppError)
	DeleteById(ctx context.Context, id string) (bool, *core.AppError)
	QueryData(ctx context.Context, query *core.DbQuery) ([]model.Category, *core.AppError)
}
