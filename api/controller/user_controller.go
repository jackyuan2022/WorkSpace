package controller

import (
	"github.com/gin-gonic/gin"
	response "github.com/jackyuan2022/workspace/api/core"
	dto "github.com/jackyuan2022/workspace/api/dto"
	service "github.com/jackyuan2022/workspace/service"
	serviceImpl "github.com/jackyuan2022/workspace/service/implement"
)

type UserController struct {
	dataService service.UserService
}

func NewUserController() *UserController {
	svc := serviceImpl.NewUserService()

	return &UserController{
		dataService: svc,
	}
}

func (t *UserController) GetUserList(c *gin.Context) {
	var r dto.GetUserListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.GetUserList(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "获取用户成功", res)
	}
}

func (t *UserController) UpdateUser(c *gin.Context) {
	var r dto.DataRequest[dto.UserDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.UpdateUser(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "更新用户成功", res)
	}
}

func (t *UserController) DeleteUser(c *gin.Context) {
	var r dto.DataRequest[dto.UserDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.DeleteUser(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "删除用户成功", res)
	}
}
