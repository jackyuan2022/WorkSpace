package service

import (
	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	core "github.com/jackyuan2022/workspace/core"
)

type CaseTaskService interface {
	GetCaseTaskList(ctx *gin.Context, r *dto.GetCaseTaskListRequest) (res *dto.DataListResponse[dto.CaseTaskDTO], err *core.AppError)
	CreateCaseTask(ctx *gin.Context, r *dto.DataRequest[dto.CaseTaskDTO]) (res *dto.DataResponse[dto.CaseTaskDTO], err *core.AppError)
	UpdateCaseTask(ctx *gin.Context, r *dto.DataRequest[dto.CaseTaskDTO]) (res *dto.DataResponse[dto.CaseTaskDTO], err *core.AppError)
	DeleteCaseTask(ctx *gin.Context, r *dto.DataRequest[dto.CaseTaskDTO]) (res *dto.DataResponse[dto.CaseTaskDTO], err *core.AppError)
}
