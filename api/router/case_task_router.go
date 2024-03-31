package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jackyuan2022/workspace/api/controller"
	middleware "github.com/jackyuan2022/workspace/api/middleware"
	app "github.com/jackyuan2022/workspace/app"
	oauth "github.com/jackyuan2022/workspace/domain/oauth"
)

type CaseTaskRouter struct{}

func (s *CaseTaskRouter) InitCaseTaskRouter(routerGroup *gin.RouterGroup) (R gin.IRoutes) {
	router := routerGroup.Group("case_task")
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("InitCaseTaskRouter error")
	}
	router.Use(middleware.OAuthMiddleware(oauthMaker))
	api := controller.NewCaseTaskController()
	router.POST("case_task_list", api.GetCaseTaskList)
	router.POST("create_case_task", api.CreateCaseTask)
	router.POST("update_case_task", api.UpdateCaseTask)
	router.POST("delete_case_task", api.DeleteCaseTask)
	return router
}
