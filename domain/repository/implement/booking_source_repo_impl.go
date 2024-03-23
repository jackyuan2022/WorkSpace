package implement

import (
	"context"
	"errors"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingSourceRepositoryImpl struct {
	BaseRepository
}

func NewBookingSourceRepository(dbContext core.DbContext) repo.BookingSourceRepository {
	dataRepo := &BookingSourceRepositoryImpl{
		BaseRepository: BaseRepository{
			dbContext: dbContext,
		},
	}
	return dataRepo
}

func (r *BookingSourceRepositoryImpl) FindById(ctx context.Context, id string) (*model.BookingSource, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.BookingSource
	err := db.WithContext(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("Data Recrod not found")
		}
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *BookingSourceRepositoryImpl) Insert(ctx context.Context, d *model.BookingSource) (*model.BookingSource, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	err := db.WithContext(ctx).Omit(clause.Associations).Create(d).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return d, nil
}

func (r *BookingSourceRepositoryImpl) Update(ctx context.Context, d *model.BookingSource) (*model.BookingSource, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.BookingSource
	err := db.WithContext(ctx).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	data.Name = d.Name
	data.Category = d.Category
	err = db.WithContext(ctx).Omit(clause.Associations).Save(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *BookingSourceRepositoryImpl) DeleteById(ctx context.Context, id string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	data := model.BookingSource{}
	err := db.WithContext(ctx).Omit(clause.Associations).Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}

func (r *BookingSourceRepositoryImpl) QueryData(ctx context.Context, query *core.DbQuery) ([]model.BookingSource, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	datas := []model.BookingSource{}
	whereClaues, values := query.GetWhereClause()
	offset := (query.PageNumber - 1) * query.PageSize
	err := db.WithContext(ctx).Where(whereClaues, values).Offset(offset).Limit(query.PageSize + 1).Find(&datas).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return datas, nil
}
