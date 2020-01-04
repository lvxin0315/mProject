package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Demo(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", nil)
}
