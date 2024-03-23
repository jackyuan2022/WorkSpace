package implement

import (
	"strings"

	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	app "github.com/jackyuan2022/workspace/app"
	core "github.com/jackyuan2022/workspace/core"
	"github.com/jackyuan2022/workspace/domain/model"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	repoImpl "github.com/jackyuan2022/workspace/domain/repository/implement"
	svc "github.com/jackyuan2022/workspace/service"
	"github.com/jackyuan2022/workspace/util"
)

type BookingSourceServiceImpl struct {
	dataRepo repo.BookingSourceRepository
}

func NewBookingSourceService() svc.BookingSourceService {
	dataRepo := repoImpl.NewBookingSourceRepository(app.AppContext.APP_DbContext)

	bookingSourceSvc := &BookingSourceServiceImpl{
		dataRepo: dataRepo,
	}

	return bookingSourceSvc
}

func (s *BookingSourceServiceImpl) GetBookingSourceList(ctx *gin.Context, r *dto.GetBookingSourceListRequest) (res *dto.DataListResponse[dto.BookingSourceDTO], err *core.AppError) {
	values := []interface{}{r.CategoryId}
	filters := []core.DbQueryFilter{core.NewDbQueryFilter("category_id", values, "EQ", "string")}
	wheres := []core.DbQueryWhere{core.NewDbQueryWhere(filters, "AND")}
	query := &core.DbQuery{
		QueryWheres: wheres,
		PageSize:    r.Pagination.PageSize,
		PageNumber:  r.Pagination.PageNumber,
	}
	result, err := s.dataRepo.QueryData(ctx, query)
	if err != nil {
		return nil, core.NewUnexpectedError("Query BookingSource Data Failure")
	}
	var dataList []dto.BookingSourceDTO
	index := len(result)
	if index > r.Pagination.PageSize {
		index = r.Pagination.PageSize
	}
	for i := 0; i < index; i++ {
		item := result[i]
		dto := s.convertModel2Dto(item)
		dataList = append(dataList, dto)
	}
	res = &dto.DataListResponse[dto.BookingSourceDTO]{
		DataList: dataList,
		Pagination: dto.PageDTO{
			PageSize:    r.Pagination.PageSize,
			PageNumber:  r.Pagination.PageNumber,
			HasNextPage: len(result) > r.Pagination.PageSize,
		},
	}
	return res, nil
}

func (s *BookingSourceServiceImpl) CreateBookingSource(ctx *gin.Context, r *dto.DataRequest[dto.BookingSourceDTO]) (res *dto.DataResponse[dto.BookingSourceDTO], err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Data.Name, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	data := s.convertDto2Model(r.Data)

	result, err := s.dataRepo.Insert(ctx, &data)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert BookingSource Data Failure")
	}
	res = &dto.DataResponse[dto.BookingSourceDTO]{
		Data: s.convertModel2Dto(*result),
	}
	return res, nil
}

func (s *BookingSourceServiceImpl) UpdateBookingSource(ctx *gin.Context, r *dto.DataRequest[dto.BookingSourceDTO]) (res *dto.DataResponse[dto.BookingSourceDTO], err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Data.Name, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	data := s.convertDto2Model(r.Data)

	result, err := s.dataRepo.Update(ctx, &data)
	if err != nil {
		return nil, core.NewUnexpectedError("Update Booking Source Data Failure")
	}
	res = &dto.DataResponse[dto.BookingSourceDTO]{
		Data: s.convertModel2Dto(*result),
	}
	return res, nil
}

func (s *BookingSourceServiceImpl) DeleteBookingSource(ctx *gin.Context, r *dto.DataRequest[dto.BookingSourceDTO]) (res *dto.DataResponse[dto.BookingSourceDTO], err *core.AppError) {
	_, err = s.dataRepo.DeleteById(ctx, r.Data.Id)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete BookingSource Data Failure")
	}
	res = &dto.DataResponse[dto.BookingSourceDTO]{
		Data: r.Data,
	}
	return res, nil
}

func (s *BookingSourceServiceImpl) convertModel2Dto(m model.BookingSource) (d dto.BookingSourceDTO) {
	d = dto.BookingSourceDTO{
		Id:   m.DbBaseModel.Id,
		Name: m.Name,
		Category: dto.CategoryDTO{
			Id:           m.Category.Id,
			Name:         m.Category.Name,
			CategoryType: m.Category.CategoryType,
		},
	}

	return d
}

func (s *BookingSourceServiceImpl) convertDto2Model(d dto.BookingSourceDTO) (m model.BookingSource) {
	id := d.Id
	if len(id) < 1 {
		id = util.GenerateId()
	}
	m = model.BookingSource{
		Name: d.Name,
		Category: model.Category{
			Name:         d.Category.Name,
			CategoryType: d.Category.CategoryType,
			DbBaseModel:  core.NewDbBaseModel(d.Category.Id),
		},
		DbBaseModel: core.NewDbBaseModel(id),
	}

	return m
}
