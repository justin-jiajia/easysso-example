package main

import (
	"github.com/justin-jiajia/easysso-example/config"
	"github.com/justin-jiajia/easysso-example/router"
	"github.com/justin-jiajia/easysso-example/session"
)

func main() {
	config.LoadConfig()
	session.InitSessions()
	r := router.InitRouter()
	r.Run(config.Config.Host)
}
