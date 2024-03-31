package repository

import (
	"context"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
)

type CaseTaskRepository interface {
	FindById(ctx context.Context, id string) (*model.CaseTask, *core.AppError)
	Insert(ctx context.Context, data *model.CaseTask) (*model.CaseTask, *core.AppError)
	Update(ctx context.Context, data *model.CaseTask) (*model.CaseTask, *core.AppError)
	DeleteById(ctx context.Context, id string) (bool, *core.AppError)
	QueryData(ctx context.Context, query *core.DbQuery) ([]model.CaseTask, *core.AppError)
}
