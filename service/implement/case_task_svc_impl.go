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

type CaseTaskServiceImpl struct {
	dataRepo repo.CaseTaskRepository
}

func NewCaseTaskService() svc.CaseTaskService {
	dataRepo := repoImpl.NewCaseTaskRepository(app.AppContext.APP_DbContext)

	caseTaskSvc := &CaseTaskServiceImpl{
		dataRepo: dataRepo,
	}

	return caseTaskSvc
}

func (s *CaseTaskServiceImpl) GetCaseTaskList(ctx *gin.Context, r *dto.GetCaseTaskListRequest) (res *dto.DataListResponse[dto.CaseTaskDTO], err *core.AppError) {
	wheres := []core.DbQueryWhere{}
	filters := []core.DbQueryFilter{}
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
		return nil, core.NewUnexpectedError("Query CaseTask Data Failure")
	}
	var dataList []dto.CaseTaskDTO
	index := len(result)
	if index > r.Pagination.PageSize {
		index = r.Pagination.PageSize
	}
	for i := 0; i < index; i++ {
		item := result[i]
		dto := s.convertModel2Dto(item)
		dataList = append(dataList, dto)
	}
	res = &dto.DataListResponse[dto.CaseTaskDTO]{
		DataList: dataList,
		Pagination: dto.PageDTO{
			PageSize:    r.Pagination.PageSize,
			PageNumber:  r.Pagination.PageNumber,
			HasNextPage: len(result) > r.Pagination.PageSize,
		},
	}
	return res, nil
}

func (s *CaseTaskServiceImpl) CreateCaseTask(ctx *gin.Context, r *dto.DataRequest[dto.CaseTaskDTO]) (res *dto.DataResponse[dto.CaseTaskDTO], err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Data.Name, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	data := s.convertDto2Model(r.Data)

	result, err := s.dataRepo.Insert(ctx, &data)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert CaseTask Data Failure")
	}
	res = &dto.DataResponse[dto.CaseTaskDTO]{
		Data: s.convertModel2Dto(*result),
	}
	return res, nil
}

func (s *CaseTaskServiceImpl) UpdateCaseTask(ctx *gin.Context, r *dto.DataRequest[dto.CaseTaskDTO]) (res *dto.DataResponse[dto.CaseTaskDTO], err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Data.Name, " ") == "" {
		return nil, core.NewValidationError("预约名称不能为空")
	}
	data := s.convertDto2Model(r.Data)
	result, err := s.dataRepo.Update(ctx, &data)
	if err != nil {
		return nil, core.NewUnexpectedError("Update CaseTask Data Failure")
	}
	res = &dto.DataResponse[dto.CaseTaskDTO]{
		Data: s.convertModel2Dto(*result),
	}
	return res, nil
}

func (s *CaseTaskServiceImpl) DeleteCaseTask(ctx *gin.Context, r *dto.DataRequest[dto.CaseTaskDTO]) (res *dto.DataResponse[dto.CaseTaskDTO], err *core.AppError) {
	_, err = s.dataRepo.DeleteById(ctx, r.Data.Id)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete CaseTask Data Failure")
	}
	res = &dto.DataResponse[dto.CaseTaskDTO]{
		Data: r.Data,
	}
	return res, nil
}

func (s *CaseTaskServiceImpl) convertModel2Dto(m model.CaseTask) (d dto.CaseTaskDTO) {
	d = dto.CaseTaskDTO{
		Id:      m.DbBaseModel.Id,
		Name:    m.Name,
		Content: m.Content,
		UserId:  m.UserId,
		CaseUser: dto.UserDTO{
			UserId:   m.CaseUser.Id,
			UserName: m.CaseUser.Name,
			Mobile:   m.CaseUser.Mobile,
		},
		Status:          m.Status,
		CaseTaskDetails: []dto.CaseTaskDetailDTO{},
	}

	if len(m.CaseTaskDetails) > 0 {
		lst := []dto.CaseTaskDetailDTO{}
		for _, item := range m.CaseTaskDetails {
			t := s.convertDetailModel2DetailDto(item)
			lst = append(lst, t)
		}
		d.CaseTaskDetails = lst
	}
	return d
}

func (s *CaseTaskServiceImpl) convertDto2Model(d dto.CaseTaskDTO) (m model.CaseTask) {
	id := d.Id
	if len(id) < 1 {
		id = util.GenerateId()
	}
	m = model.CaseTask{
		Name:        d.Name,
		Content:     d.Content,
		UserId:      d.UserId,
		Status:      d.Status,
		CaseTime:    time.UnixMilli(d.CaseTime),
		DbBaseModel: core.NewDbBaseModel(id),
		CaseUser: model.User{
			Name:        d.CaseUser.UserName,
			Mobile:      d.CaseUser.Mobile,
			DbBaseModel: core.NewDbBaseModel(d.CaseUser.UserId),
		},
		CaseTaskDetails: []model.CaseTaskDetail{},
	}
	if len(d.CaseTaskDetails) > 0 {
		lst := []model.CaseTaskDetail{}
		for _, item := range d.CaseTaskDetails {
			t := s.convertDetailDto2DetailModel(item)
			lst = append(lst, t)
		}
		m.CaseTaskDetails = lst
	}

	return m
}

func (s *CaseTaskServiceImpl) convertDetailModel2DetailDto(m model.CaseTaskDetail) (d dto.CaseTaskDetailDTO) {
	d = dto.CaseTaskDetailDTO{
		Id:             m.DbBaseModel.Id,
		Name:           m.Name,
		Content:        m.Content,
		UserId:         m.UserId,
		Status:         m.Status,
		CaseTime:       m.CaseTime.UnixMilli(),
		ExpirationDays: m.ExpirationDays,
		ExpirationTime: nil,
		CaseUser: dto.UserDTO{
			UserId:   m.CaseUser.Id,
			UserName: m.CaseUser.Name,
			Mobile:   m.CaseUser.Mobile,
		},
	}
	if m.ExpirationTime != nil {
		t := m.ExpirationTime.UnixMilli()
		d.ExpirationTime = &t
	}

	return d
}

func (s *CaseTaskServiceImpl) convertDetailDto2DetailModel(d dto.CaseTaskDetailDTO) (m model.CaseTaskDetail) {
	id := d.Id
	if len(id) < 1 {
		id = util.GenerateId()
	}
	m = model.CaseTaskDetail{
		Name:           d.Name,
		Content:        d.Content,
		UserId:         d.UserId,
		Status:         d.Status,
		CaseTime:       time.UnixMilli(d.CaseTime),
		ExpirationDays: d.ExpirationDays,
		ExpirationTime: nil,
		DbBaseModel:    core.NewDbBaseModel(id),
		CaseUser: model.User{
			Name:        d.CaseUser.UserName,
			Mobile:      d.CaseUser.Mobile,
			DbBaseModel: core.NewDbBaseModel(d.CaseUser.UserId),
		},
	}
	if d.ExpirationTime != nil {
		t := time.UnixMilli(*d.ExpirationTime)
		m.ExpirationTime = &t
	}

	return m
}
