package session

import (
	"github.com/gorilla/sessions"
	"github.com/justin-jiajia/easysso-example/config"
)

var Store *sessions.CookieStore

func InitSessions() {
	Store = sessions.NewCookieStore([]byte(config.Config.SessionKey))
}
