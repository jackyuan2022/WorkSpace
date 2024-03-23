package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jackyuan2022/workspace/api/controller"
	middleware "github.com/jackyuan2022/workspace/api/middleware"
	app "github.com/jackyuan2022/workspace/app"
	oauth "github.com/jackyuan2022/workspace/domain/oauth"
)

type BookingRouter struct{}

func (s *BookingRouter) InitBookingRouter(routerGroup *gin.RouterGroup) (R gin.IRoutes) {
	router := routerGroup.Group("booking")
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("InitBooingRouter error")
	}
	router.Use(middleware.OAuthMiddleware(oauthMaker))
	api := controller.NewBookingController()
	router.POST("booking_list", api.GetBookingList)
	router.POST("create_booking", api.CreateBooking)
	router.POST("update_booking", api.UpdateBooking)
	router.POST("delete_booking", api.DeleteBooking)
	return router
}
