package repository

import (
	"context"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
)

type OAuthSessionRepository interface {
	FindById(ctx context.Context, id string) (*model.OAuthSession, *core.AppError)
	Insert(ctx context.Context, session *model.OAuthSession) (*model.OAuthSession, *core.AppError)
	Update(ctx context.Context, session *model.OAuthSession) (*model.OAuthSession, *core.AppError)
	DeleteById(ctx context.Context, id string) (bool, *core.AppError)
}
