package implement

import (
	"context"
	"errors"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	"gorm.io/gorm"
)

type CaseTaskRepositoryImpl struct {
	BaseRepository
}

func NewCaseTaskRepository(dbContext core.DbContext) repo.CaseTaskRepository {
	dataRepo := &CaseTaskRepositoryImpl{
		BaseRepository: BaseRepository{
			dbContext: dbContext,
		},
	}
	return dataRepo
}

func (r *CaseTaskRepositoryImpl) FindById(ctx context.Context, id string) (*model.CaseTask, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.CaseTask
	err := db.WithContext(ctx).Where("id = ?", id).Preload("CaseTaskDetails").First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.NewNotFoundError("Data Recrod not found")
		}
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *CaseTaskRepositoryImpl) Insert(ctx context.Context, d *model.CaseTask) (*model.CaseTask, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	err := db.WithContext(ctx).Select("CaseTaskDetails").Create(d).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return d, nil
}

func (r *CaseTaskRepositoryImpl) Update(ctx context.Context, d *model.CaseTask) (*model.CaseTask, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	var data model.CaseTask
	err := db.WithContext(ctx).Where("id = ?", d.Id).Preload("CaseTaskDetails").First(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	data.Name = d.Name
	data.Content = d.Content
	data.CaseTime = d.CaseTime
	data.Status = d.Status
	data.UserId = d.UserId
	data.CaseTaskDetails = d.CaseTaskDetails
	err = db.WithContext(ctx).Select("CaseTaskDetails").Save(&data).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return &data, nil
}

func (r *CaseTaskRepositoryImpl) DeleteById(ctx context.Context, id string) (bool, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return false, appErr
	}
	data := model.CaseTask{}
	err := db.WithContext(ctx).Select("CaseTaskDetails").Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return false, core.AsAppError(err)
	}
	return true, nil
}

func (r *CaseTaskRepositoryImpl) QueryData(ctx context.Context, query *core.DbQuery) ([]model.CaseTask, *core.AppError) {
	db, appErr := r.getDb()
	if appErr != nil {
		return nil, appErr
	}
	datas := []model.CaseTask{}
	whereClaues, values, order := query.GetWhereClause()
	offset := (query.PageNumber - 1) * query.PageSize
	err := db.WithContext(ctx).Where(whereClaues, values...).Order(order).Offset(offset).Limit(query.PageSize + 1).Preload("CaseTaskDetails").Find(&datas).Error
	if err != nil {
		return nil, core.AsAppError(err)
	}
	return datas, nil
}
