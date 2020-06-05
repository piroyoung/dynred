package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RedirectHandler interface {
	HandleWithMeta(c *gin.Context)
	HandleWith301(c *gin.Context)
}

func handleWithMeta(url string, c *gin.Context) {
	c.HTML(http.StatusOK, "redirect.tmpl", gin.H{
		"url": url,
	})
}
