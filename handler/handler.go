package handler

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/justin-jiajia/easysso-example/config"
	"github.com/justin-jiajia/easysso-example/session"
	"github.com/justin-jiajia/easysso-example/utils"
)

type Res struct {
	ID          int
	UserName    string
	JoinTime    int
	SettingsURL string
}

func IndexHandler(c *gin.Context) {
	session, err := session.Store.Get(c.Request, "session")
	if err != nil {
		log.Println(err)
		return
	}

	token, ok := session.Values["token"].(string)
	if !ok || token == "" {
		c.HTML(http.StatusOK, "unloginindex.html", nil)
	}
	res := utils.GetUserInfo(token)
	p, _ := url.JoinPath(config.Config.Server, "/ui/settings.html")
	r := Res{
		ID:          res.ID,
		UserName:    res.UserName,
		JoinTime:    res.JoinTime,
		SettingsURL: p,
	}
	c.HTML(http.StatusOK, "index.html", r)
}

func LogoutHandler(c *gin.Context) {
	session, err := session.Store.Get(c.Request, "session")
	if err != nil {
		log.Println(err)
		return
	}
	session.Values["token"] = ""
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func RedirectHandler(c *gin.Context) {
	url, state := utils.GetRedirectURL()
	session, err := session.Store.Get(c.Request, "session")
	if err != nil {
		log.Println(err)
		return
	}
	if session.Values["token"] != "" && session.Values["token"] != nil {
		c.Redirect(http.StatusFound, "/")
	}
	session.Values["state"] = state
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound, url)
}

func CallbackHandler(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	if code == "" {
		panic("code is empty")
	}
	token := utils.CodeToToken(code)
	session, err := session.Store.Get(c.Request, "session")
	if err != nil {
		log.Println(err)
		return
	}
	if state != session.Values["state"] {
		panic("state is not equal")
	}
	session.Values["token"] = token
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound, "/")
}
