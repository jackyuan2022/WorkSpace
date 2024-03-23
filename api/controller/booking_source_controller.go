package controller

import (
	"github.com/gin-gonic/gin"
	response "github.com/jackyuan2022/workspace/api/core"
	dto "github.com/jackyuan2022/workspace/api/dto"
	service "github.com/jackyuan2022/workspace/service"
	serviceImpl "github.com/jackyuan2022/workspace/service/implement"
)

type BookingSourceController struct {
	dataService service.BookingSourceService
}

func NewBookingSourceController() *BookingSourceController {
	svc := serviceImpl.NewBookingSourceService()

	return &BookingSourceController{
		dataService: svc,
	}
}

func (t *BookingSourceController) GetBookingSourceList(c *gin.Context) {
	var r dto.GetBookingSourceListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.GetBookingSourceList(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "获取类别成功", res)
	}
}

func (t *BookingSourceController) CreateBookingSource(c *gin.Context) {
	var r dto.DataRequest[dto.BookingSourceDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.CreateBookingSource(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "创建类别成功", res)
	}
}

func (t *BookingSourceController) UpdateBookingSource(c *gin.Context) {
	var r dto.DataRequest[dto.BookingSourceDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.UpdateBookingSource(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "更新类别成功", res)
	}
}

func (t *BookingSourceController) DeleteBookingSource(c *gin.Context) {
	var r dto.DataRequest[dto.BookingSourceDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.DeleteBookingSource(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "删除预约资源成功", res)
	}
}
