package router

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/justin-jiajia/easysso-example/handler"
	templatefs "github.com/justin-jiajia/easysso-example/template"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.SetHTMLTemplate(template.Must(template.ParseFS(templatefs.Template, "*.html")))
	router.GET("/", handler.IndexHandler)
	router.GET("/login", handler.RedirectHandler)
	router.GET("/callback", handler.CallbackHandler)
	router.GET("/logout", handler.LogoutHandler)
	return router
}
