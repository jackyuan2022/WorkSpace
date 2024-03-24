package implement

import (
	"context"
	"errors"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	BaseRepository
}

func NewCategoryRepository(dbContext core.DbContext) repo.CategoryRepository {
	dataRepo := &CategoryRepositoryImpl{
		BaseRepository: BaseRepository{
			dbContext: dbContext,
		},
	}
	return dataRepo
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, id string) (*model.Category, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.Category
	err := db.WithContext(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("Data Recrod not found")
		}
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *CategoryRepositoryImpl) Insert(ctx context.Context, d *model.Category) (*model.Category, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	err := db.WithContext(ctx).Create(d).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return d, nil
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, d *model.Category) (*model.Category, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.Category
	err := db.WithContext(ctx).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	data.Name = d.Name
	data.Icon = d.Icon
	data.Order = d.Order
	data.CategoryType = d.CategoryType
	err = db.WithContext(ctx).Save(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *CategoryRepositoryImpl) DeleteById(ctx context.Context, id string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	data := model.Category{}
	err := db.WithContext(ctx).Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}

func (r *CategoryRepositoryImpl) QueryData(ctx context.Context, query *core.DbQuery) ([]model.Category, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	datas := []model.Category{}
	whereClaues, values, order := query.GetWhereClause()
	offset := (query.PageNumber - 1) * query.PageSize
	err := db.WithContext(ctx).Where(whereClaues, values...).Order(order).Offset(offset).Limit(query.PageSize + 1).Find(&datas).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return datas, nil
}
