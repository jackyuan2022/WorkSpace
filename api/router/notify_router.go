package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/jackyuan2022/workspace/api/middleware"
	app "github.com/jackyuan2022/workspace/app"
	oauth "github.com/jackyuan2022/workspace/domain/oauth"
)

type SystemDesignerRouter struct{}

func (s *CategoryRouter) InitSystemDesignerRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	designerRouter := router.Group("notify")
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("InitSystemDesignerRouter error")
	}

	designerRouter.Use(middleware.OAuthMiddleware(oauthMaker))
	designerRouter.GET("/desinger", func(c *gin.Context) {
		c.JSON(http.StatusOK, "disinger ok")
	})
	return designerRouter
}
