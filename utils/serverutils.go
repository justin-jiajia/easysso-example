package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/justin-jiajia/easysso-example/config"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type CodeToTokenRequest struct {
	Code         string `json:"code"`
	ClientSecret string `json:"client_secret"`
	ClientID     string `json:"client_id"`
}

type CodeToTokenResponse struct {
	Token string `json:"token"`
}

func CodeToToken(code string) string {
	url, err := url.JoinPath(config.Config.Server, "/api/oath2/gettoken/")
	if err != nil {
		log.Panic(err)
	}

	req := CodeToTokenRequest{
		Code:         code,
		ClientSecret: config.Config.ClientSecret,
		ClientID:     config.Config.ClientID,
	}
	reqjsoned, err := json.Marshal(req)
	if err != nil {
		log.Panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqjsoned))
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var res ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			log.Panic(err)
		}
		log.Panic(res.Error)
	}

	var res CodeToTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Panic(err)
	}
	return res.Token
}

type InfoRequest struct {
	Token        string `json:"token"`
	ClientSecret string `json:"client_secret"`
	ClientID     string `json:"client_id"`
}

type InfoResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	JoinTime int    `json:"jointime"`
}

func GetUserInfo(token string) InfoResponse {
	url, err := url.JoinPath(config.Config.Server, "/api/oath2/information/")
	if err != nil {
		log.Panic(err)
	}

	req := InfoRequest{
		Token:        token,
		ClientSecret: config.Config.ClientSecret,
		ClientID:     config.Config.ClientID,
	}

	reqjsoned, err := json.Marshal(req)
	if err != nil {
		log.Panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqjsoned))
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var res ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			log.Panic(err)
		}
		log.Panic(res.Error)
	}

	var res InfoResponse
	// s, _ := io.ReadAll(resp.Body)
	// println(string(s))
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

const state_rand_char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRedirectURL() (string, string) {
	urlparsed, err := url.Parse(config.Config.Server)
	if err != nil {
		log.Panic(err)
	}
	state := ""
	for i := 0; i < 16; i++ {
		state += string(state_rand_char[rand.Intn(len(state_rand_char))])
	}
	urlparsed.Path = "/ui/authorize.html"
	params := url.Values{}
	params.Add("client_id", config.Config.ClientID)
	params.Add("state", state)
	urlparsed.RawQuery = params.Encode()
	return urlparsed.String(), state
}
