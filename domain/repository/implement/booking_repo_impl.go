package implement

import (
	"context"
	"errors"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	"gorm.io/gorm"
)

type BookingRepositoryImpl struct {
	BaseRepository
}

func NewBookingRepository(dbContext core.DbContext) repo.BookingRepository {
	dataRepo := &BookingRepositoryImpl{
		BaseRepository: BaseRepository{
			dbContext: dbContext,
		},
	}
	return dataRepo
}

func (r *BookingRepositoryImpl) FindById(ctx context.Context, id string) (*model.Booking, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.Booking
	err := db.WithContext(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("Data Recrod not found")
		}
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *BookingRepositoryImpl) Insert(ctx context.Context, d *model.Booking) (*model.Booking, *core.AppError) {
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

func (r *BookingRepositoryImpl) Update(ctx context.Context, d *model.Booking) (*model.Booking, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.Booking
	err := db.WithContext(ctx).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	data.Title = d.Title
	data.Content = d.Content
	data.BookingStartTime = d.BookingStartTime
	data.BookingEndTime = d.BookingEndTime
	data.Status = d.Status
	data.UserId = d.UserId
	data.BookingSourceId = d.BookingSourceId
	err = db.WithContext(ctx).Omit("BookingSource.*", "BookingUser.*").Save(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *BookingRepositoryImpl) DeleteById(ctx context.Context, id string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	data := model.Booking{}
	err := db.WithContext(ctx).Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}

func (r *BookingRepositoryImpl) QueryData(ctx context.Context, query *core.DbQuery) ([]model.Booking, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	datas := []model.Booking{}
	whereClaues, values := query.GetWhereClause()
	offset := (query.PageNumber - 1) * query.PageSize
	err := db.WithContext(ctx).Where(whereClaues, values).Offset(offset).Limit(query.PageSize + 1).Find(&datas).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return datas, nil
}
