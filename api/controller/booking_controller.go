package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	response "github.com/jackyuan2022/workspace/api/core"
	dto "github.com/jackyuan2022/workspace/api/dto"
	service "github.com/jackyuan2022/workspace/service"
	serviceImpl "github.com/jackyuan2022/workspace/service/implement"
)

type BookingController struct {
	dataService service.BookingService
}

func NewBookingController() *BookingController {
	svc := serviceImpl.NewBookingService()

	return &BookingController{
		dataService: svc,
	}
}

func (t *BookingController) GetBookingList(c *gin.Context) {
	var r dto.GetBookingListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.GetBookingList(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "获取预约成功", res)
	}
}

func (t *BookingController) CreateBooking(c *gin.Context) {
	var r dto.DataRequest[dto.BookingDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		fmt.Println(err.Error())
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.CreateBooking(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "创建预约成功", res)
	}
}

func (t *BookingController) UpdateBooking(c *gin.Context) {
	var r dto.DataRequest[dto.BookingDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.UpdateBooking(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "更新预约成功", res)
	}
}

func (t *BookingController) DeleteBooking(c *gin.Context) {
	var r dto.DataRequest[dto.BookingDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.DeleteBooking(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "删除预约成功", res)
	}
}
