package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReplceText struct {
	ClientID    string
	RedirectURI string
	Scope       []string
}

func Authorize(c *gin.Context) {
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	scope := c.Query("scope")
	responseType := c.Query("response_type")
	state := c.Query("state")

	if clientID == "" || redirectURI == "" || scope == "" || responseType == "" || state == "" {
		c.JSON(http.StatusBadRequest, "invalid parameter")
		return
	}
	if responseType != "code" {
		c.JSON(http.StatusBadRequest, "invalid parameter")
		return
	}
	replceText := ReplceText{
		ClientID:    clientID,
		RedirectURI: redirectURI,
		Scope:       strings.Split(scope, " "),
	}
	// Content-Type を "text/html" に変換したり、template.Executeも行い、HTMLをレンダリングする
	c.HTML(http.StatusOK, "authorize.html", replceText)
	return
}
