package implement

import (
	"strings"
	"time"

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

type BookingServiceImpl struct {
	dataRepo repo.BookingRepository
}

func NewBookingService() svc.BookingService {
	dataRepo := repoImpl.NewBookingRepository(app.AppContext.APP_DbContext)

	bookingySvc := &BookingServiceImpl{
		dataRepo: dataRepo,
	}

	return bookingySvc
}

func (s *BookingServiceImpl) GetBookingList(ctx *gin.Context, r *dto.GetBookingListRequest) (res *dto.DataListResponse[dto.BookingDTO], err *core.AppError) {
	wheres := []core.DbQueryWhere{}
	filters := []core.DbQueryFilter{}
	if len(r.BookingSourceId) > 0 {
		filters = append(filters, core.NewDbQueryFilter("booking_source_id", []interface{}{r.BookingSourceId}, "EQ", "string"))
	}
	if len(r.UserId) > 0 {
		filters = append(filters, core.NewDbQueryFilter("user_id", []interface{}{r.UserId}, "EQ", "string"))
	}
	wheres = append(wheres, core.NewDbQueryWhere(filters, "AND"))
	query := &core.DbQuery{
		QueryWheres: wheres,
		PageSize:    r.Pagination.PageSize,
		PageNumber:  r.Pagination.PageNumber,
	}
	result, err := s.dataRepo.QueryData(ctx, query)
	if err != nil {
		return nil, core.NewUnexpectedError("Query Booking Data Failure")
	}
	var bookingList []dto.BookingDTO
	index := len(result)
	if index > r.Pagination.PageSize {
		index = r.Pagination.PageSize
	}
	for i := 0; i < index; i++ {
		item := result[i]
		dto := s.convertModel2Dto(item)
		bookingList = append(bookingList, dto)
	}
	res = &dto.DataListResponse[dto.BookingDTO]{
		DataList: bookingList,
		Pagination: dto.PageDTO{
			PageSize:    r.Pagination.PageSize,
			PageNumber:  r.Pagination.PageNumber,
			HasNextPage: len(result) > r.Pagination.PageSize,
		},
	}
	return res, nil
}

func (s *BookingServiceImpl) CreateBooking(ctx *gin.Context, r *dto.DataRequest[dto.BookingDTO]) (res *dto.DataResponse[dto.BookingDTO], err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Data.Title, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	booking := s.convertDto2Model(r.Data)

	result, err := s.dataRepo.Insert(ctx, &booking)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert Category Data Failure")
	}
	res = &dto.DataResponse[dto.BookingDTO]{
		Data: s.convertModel2Dto(*result),
	}
	return res, nil
}

func (s *BookingServiceImpl) UpdateBooking(ctx *gin.Context, r *dto.DataRequest[dto.BookingDTO]) (res *dto.DataResponse[dto.BookingDTO], err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Data.Title, " ") == "" {
		return nil, core.NewValidationError("预约名称不能为空")
	}
	booking := s.convertDto2Model(r.Data)
	result, err := s.dataRepo.Update(ctx, &booking)
	if err != nil {
		return nil, core.NewUnexpectedError("Update Booking Data Failure")
	}
	res = &dto.DataResponse[dto.BookingDTO]{
		Data: s.convertModel2Dto(*result),
	}
	return res, nil
}

func (s *BookingServiceImpl) DeleteBooking(ctx *gin.Context, r *dto.DataRequest[dto.BookingDTO]) (res *dto.DataResponse[dto.BookingDTO], err *core.AppError) {
	_, err = s.dataRepo.DeleteById(ctx, r.Data.Id)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete Booking Data Failure")
	}
	res = &dto.DataResponse[dto.BookingDTO]{
		Data: r.Data,
	}
	return res, nil
}

func (s *BookingServiceImpl) convertModel2Dto(m model.Booking) (d dto.BookingDTO) {
	d = dto.BookingDTO{
		Id:               m.DbBaseModel.Id,
		Title:            m.Title,
		Content:          m.Content,
		BookingStartTime: m.BookingStartTime.UnixMilli(),
		BookingEndTime:   nil,
		BookingSourceId:  m.BookingSourceId,
		UserId:           m.UserId,
		BookingSource: dto.BookingSourceDTO{
			Id:   m.BookingSource.Id,
			Name: m.BookingSource.Name,
		},
		BookingUser: dto.UserDTO{
			UserId:   m.BookingUser.Id,
			UserName: m.BookingUser.Name,
			Mobile:   m.BookingUser.Mobile,
		},
	}
	if m.BookingEndTime != nil {
		endTime := m.BookingEndTime.UnixMilli()
		d.BookingEndTime = &endTime
	}

	return d
}

func (s *BookingServiceImpl) convertDto2Model(d dto.BookingDTO) (m model.Booking) {
	id := d.Id
	if len(id) < 1 {
		id = util.GenerateId()
	}
	m = model.Booking{
		Title:            d.Title,
		Content:          d.Content,
		BookingStartTime: time.UnixMilli(d.BookingStartTime),
		BookingEndTime:   nil,
		BookingSourceId:  d.BookingSourceId,
		UserId:           d.UserId,
		BookingSource: model.BookingSource{
			Name:        d.BookingSource.Name,
			DbBaseModel: core.NewDbBaseModel(d.BookingSource.Id),
		},
		BookingUser: model.User{
			Name:        d.BookingUser.UserName,
			Mobile:      d.BookingUser.Mobile,
			DbBaseModel: core.NewDbBaseModel(d.BookingUser.UserId),
		},
		Status:      d.Status,
		DbBaseModel: core.NewDbBaseModel(id),
	}
	if d.BookingEndTime != nil {
		endTime := time.UnixMilli(*d.BookingEndTime)
		m.BookingEndTime = &endTime
	}

	return m
}
