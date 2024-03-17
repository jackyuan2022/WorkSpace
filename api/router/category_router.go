package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jackyuan2022/workspace/api/controller"
	middleware "github.com/jackyuan2022/workspace/api/middleware"
	app "github.com/jackyuan2022/workspace/app"
	oauth "github.com/jackyuan2022/workspace/domain/oauth"
)

type CategoryRouter struct{}

func (s *CategoryRouter) InitCategoryRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	categoryRouter := router.Group("category")
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("InitCategoryRouter error")
	}
	categoryRouter.Use(middleware.OAuthMiddleware(oauthMaker))
	api := controller.NewCategoryController()
	categoryRouter.POST("category_list", api.GetCategoryList)

	return categoryRouter
}
