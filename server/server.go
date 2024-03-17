package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	migrator "github.com/jackyuan2022/workspace/domain/migrator"
	"github.com/unrolled/secure"

	router "github.com/jackyuan2022/workspace/api/router"
	app "github.com/jackyuan2022/workspace/app"
)

func Run() {
	fmt.Println("gpaas server start......")
	app.InitAppContext()

	// fmt.Println(AppContext.APP_DbContext.GetDb() == nil)
	// fmt.Println(AppContext.APP_REDIS == nil)

	fmt.Println("migrate database start......")
	migrator.MigrateDb(app.AppContext.APP_DbContext)
	fmt.Println("migrate database end......")

	fmt.Println("initialization router start......")
	router := initializeRouter()
	httpsRouter := initializeRouter()
	httpsRouter.Use(TlsHandler()) // 处理SSL的中间件
	// router.Run(":8081")
	// if err := router.RunTLS(":443", "./certs/loongkirin.chat.crt", "./certs/loongkirin.chat.key"); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := http.ListenAndServeTLS(":443", "./certs/loongkirin.chat.crt", "./certs/loongkirin.chat.key", router); err != nil {
	// 	fmt.Println(err)
	// }

	go httpsRouter.RunTLS(":443", "./certs/loongkirin.chat.crt", "./certs/loongkirin.chat.key")
	router.Run(":8081")
}

// 初始化gin总路由
func initializeRouter() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	v1Group := Router.Group("v1")
	ginRouterEntry := router.Entry{}
	ginRouterEntry.InitAllRouter(v1Group)

	return Router
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
