package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	response "github.com/jackyuan2022/workspace/api/core"
	dto "github.com/jackyuan2022/workspace/api/dto"
	service "github.com/jackyuan2022/workspace/service"
	serviceImpl "github.com/jackyuan2022/workspace/service/implement"
)

type CaseTaskController struct {
	dataService service.CaseTaskService
}

func NewCaseTaskController() *CaseTaskController {
	svc := serviceImpl.NewCaseTaskService()

	return &CaseTaskController{
		dataService: svc,
	}
}

func (t *CaseTaskController) GetCaseTaskList(c *gin.Context) {
	var r dto.GetCaseTaskListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.GetCaseTaskList(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "获取预约成功", res)
	}
}

func (t *CaseTaskController) CreateCaseTask(c *gin.Context) {
	var r dto.DataRequest[dto.CaseTaskDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		fmt.Println(err.Error())
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.CreateCaseTask(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "创建任务成功", res)
	}
}

func (t *CaseTaskController) UpdateCaseTask(c *gin.Context) {
	var r dto.DataRequest[dto.CaseTaskDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.UpdateCaseTask(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "更新任务成功", res)
	}
}

func (t *CaseTaskController) DeleteCaseTask(c *gin.Context) {
	var r dto.DataRequest[dto.CaseTaskDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.DeleteCaseTask(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "删除任务成功", res)
	}
}
