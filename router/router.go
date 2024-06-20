package router

import (
	"github.com/gin-gonic/gin"

	"github.com/justin-jiajia/easysso-example/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", handler.IndexHandler)
	router.GET("/login", handler.RedirectHandler)
	router.GET("/callback", handler.CallbackHandler)
	router.GET("/logout", handler.LogoutHandler)
	return router
}
