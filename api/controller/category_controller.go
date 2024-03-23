package controller

import (
	"github.com/gin-gonic/gin"
	response "github.com/jackyuan2022/workspace/api/core"
	dto "github.com/jackyuan2022/workspace/api/dto"
	service "github.com/jackyuan2022/workspace/service"
	serviceImpl "github.com/jackyuan2022/workspace/service/implement"
)

type CategoryController struct {
	dataService service.CategoryService
}

func NewCategoryController() *CategoryController {
	svc := serviceImpl.NewCategoryService()

	return &CategoryController{
		dataService: svc,
	}
}

func (t *CategoryController) GetCategoryList(c *gin.Context) {
	var r dto.GetCategoryListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.GetCategoryList(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "获取类别成功", res)
	}
}

func (t *CategoryController) CreateCategory(c *gin.Context) {
	var r dto.DataRequest[dto.CategoryDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.CreateCategory(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "创建类别成功", res)
	}
}

func (t *CategoryController) UpdateCategory(c *gin.Context) {
	var r dto.DataRequest[dto.CategoryDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.UpdateCategory(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "更新类别成功", res)
	}
}

func (t *CategoryController) DeleteCategory(c *gin.Context) {
	var r dto.DataRequest[dto.CategoryDTO]
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}
	res, err := t.dataService.DeleteCategory(c, &r)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "删除类别成功", res)
	}
}
