package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jackyuan2022/workspace/api/controller"
	middleware "github.com/jackyuan2022/workspace/api/middleware"
	app "github.com/jackyuan2022/workspace/app"
	oauth "github.com/jackyuan2022/workspace/domain/oauth"
)

type BookingSourceRouter struct{}

func (s *BookingSourceRouter) InitBookingSourceRouter(routerGroup *gin.RouterGroup) (R gin.IRoutes) {
	router := routerGroup.Group("booking_source")
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("InitBooingSourceRouter error")
	}
	router.Use(middleware.OAuthMiddleware(oauthMaker))
	api := controller.NewBookingSourceController()
	router.POST("booking_source_list", api.GetBookingSourceList)
	router.POST("create_booking_source", api.CreateBookingSource)
	router.POST("update_booking_source", api.UpdateBookingSource)
	router.POST("delete_booking_source", api.DeleteBookingSource)
	return router
}
