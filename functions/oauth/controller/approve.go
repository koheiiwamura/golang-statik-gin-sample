package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Approve(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
