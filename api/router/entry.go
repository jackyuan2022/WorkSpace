package router

import (
	"github.com/gin-gonic/gin"
)

type Entry struct {
}

func (e *Entry) InitAllRouter(router *gin.RouterGroup) {
	authRouter := AuthorityRouter{}
	authRouter.InitAuthorityRouter(router)

	categoryRouter := CategoryRouter{}
	categoryRouter.InitCategoryRouter(router)

	bookingSourceRouter := BookingSourceRouter{}
	bookingSourceRouter.InitBookingSourceRouter(router)

	bookingRouter := BookingRouter{}
	bookingRouter.InitBookingRouter(router)

	sysDesinger := CategoryRouter{}
	sysDesinger.InitSystemDesignerRouter(router)

}
