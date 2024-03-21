package implement

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	app "github.com/jackyuan2022/workspace/app"
	core "github.com/jackyuan2022/workspace/core"
	"github.com/jackyuan2022/workspace/domain/model"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	repoImpl "github.com/jackyuan2022/workspace/domain/repository/implement"
	svc "github.com/jackyuan2022/workspace/service"
	util "github.com/jackyuan2022/workspace/util"
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

func (s *BookingServiceImpl) GetBookingList(ctx *gin.Context, r *dto.GetBookingListRequest) (res *dto.GetBookingListResponse, err *core.AppError) {
	values := []interface{}{r.CategoryId}
	filters := []core.DbQueryFilter{core.NewDbQueryFilter("category_id", values, "EQ", "string")}
	wheres := []core.DbQueryWhere{core.NewDbQueryWhere(filters, "AND")}
	query := &core.DbQuery{
		QueryWheres: wheres,
		PageSize:    r.PageSize,
		PageNumber:  r.PageNumber,
	}
	result, err := s.dataRepo.QueryData(ctx, query)
	if err != nil {
		return nil, core.NewUnexpectedError("Query Booking Data Failure")
	}
	var bookingList []dto.BookingDTO
	index := len(result)
	if index > r.PageSize {
		index = r.PageSize
	}
	for i := 0; i < index; i++ {
		item := result[i]
		dto := dto.BookingDTO{
			Id:               item.DbBaseModel.Id,
			Title:            item.Title,
			Content:          item.Content.String,
			BookingStartTime: item.BookingStartTime,
			BookingEndTime:   &item.BookingEndTime.Time,
			Category: dto.CategoryDTO{
				Id:           item.Category.Id,
				Name:         item.Category.Name,
				CategoryType: item.Category.CategoryType,
			},
			BookingUser: dto.UserDTO{
				UserId:   item.BookingUser.Id,
				UserName: item.BookingUser.Name,
				Mobile:   item.BookingUser.Mobile,
			},
		}
		bookingList = append(bookingList, dto)
	}
	res = &dto.GetBookingListResponse{
		BookingList: bookingList,
		HasNextPage: len(result) >= r.PageSize+1,
	}
	return res, nil
}

func (s *BookingServiceImpl) CreateBooking(ctx *gin.Context, r *dto.BookingRequest) (res *dto.BookingResponse, err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Booking.Title, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	booking := model.Booking{
		Title:       r.Booking.Title,
		DbBaseModel: core.NewDbBaseModel(util.GenerateId()),
	}

	result, err := s.dataRepo.Insert(ctx, &booking)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert Category Data Failure")
	}
	res = &dto.BookingResponse{
		Booking: convertModel2Dto(*result),
	}
	return res, nil
}

func (s *BookingServiceImpl) UpdateBooking(ctx *gin.Context, r *dto.BookingRequest) (res *dto.BookingResponse, err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Booking.Title, " ") == "" {
		return nil, core.NewValidationError("预约名称不能为空")
	}
	booking := convertDto2Model(r.Booking)
	result, err := s.dataRepo.Update(ctx, &booking)
	if err != nil {
		return nil, core.NewUnexpectedError("Update Booking Data Failure")
	}
	res = &dto.BookingResponse{
		Booking: convertModel2Dto(*result),
	}
	return res, nil
}

func (s *BookingServiceImpl) DeleteBooking(ctx *gin.Context, r *dto.BookingRequest) (res *dto.BookingResponse, err *core.AppError) {
	_, err = s.dataRepo.DeleteById(ctx, r.Booking.Id)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete Booking Data Failure")
	}
	res = &dto.BookingResponse{
		Booking: r.Booking,
	}
	return res, nil
}

func convertModel2Dto(m model.Booking) (d dto.BookingDTO) {
	d = dto.BookingDTO{
		Id:               m.DbBaseModel.Id,
		Title:            m.Title,
		Content:          m.Content.String,
		BookingStartTime: m.BookingStartTime,
		BookingEndTime:   nil,
		Category: dto.CategoryDTO{
			Id:           m.Category.Id,
			Name:         m.Category.Name,
			CategoryType: m.Category.CategoryType,
		},
		BookingUser: dto.UserDTO{
			UserId:   m.BookingUser.Id,
			UserName: m.BookingUser.Name,
			Mobile:   m.BookingUser.Mobile,
		},
	}
	if m.BookingEndTime.Valid {
		d.BookingEndTime = &m.BookingEndTime.Time
	}

	return d
}

func convertDto2Model(d dto.BookingDTO) (m model.Booking) {
	m = model.Booking{
		Title:            d.Title,
		Content:          sql.NullString{String: d.Content},
		BookingStartTime: d.BookingStartTime,
		BookingEndTime:   sql.NullTime{},
		Category: model.Category{
			Name:        d.Category.Name,
			DbBaseModel: core.NewDbBaseModel(d.Category.Id),
		},
		BookingUser: model.User{
			Name:        d.BookingUser.UserName,
			Mobile:      d.BookingUser.Mobile,
			DbBaseModel: core.NewDbBaseModel(d.BookingUser.UserId),
		},
		Status:      d.Status,
		DbBaseModel: core.NewDbBaseModel(d.Id),
	}
	if d.BookingEndTime != nil {
		m.BookingEndTime.Time = *d.BookingEndTime
	}

	return m
}
