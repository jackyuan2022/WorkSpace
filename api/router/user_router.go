package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jackyuan2022/workspace/api/controller"
	middleware "github.com/jackyuan2022/workspace/api/middleware"
	app "github.com/jackyuan2022/workspace/app"
	oauth "github.com/jackyuan2022/workspace/domain/oauth"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := router.Group("user")
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("InitUserRouter error")
	}
	userRouter.Use(middleware.OAuthMiddleware(oauthMaker))
	api := controller.NewUserController()
	userRouter.POST("user_list", api.GetUserList)
	userRouter.POST("update_user", api.UpdateUser)
	userRouter.POST("delete_user", api.DeleteUser)
	return userRouter
}
