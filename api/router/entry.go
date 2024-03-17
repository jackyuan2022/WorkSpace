package router

import (
	"github.com/gin-gonic/gin"
)

type Entry struct {
}

func (e *Entry) InitAllRouter(router *gin.RouterGroup) {
	sysAuth := AuthorityRouter{}
	sysAuth.InitSystemAuthorityRouter(router)

	sysDesinger := CategoryRouter{}
	sysDesinger.InitSystemDesignerRouter(router)

}
