package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "not found",
		"message": "We drifted to Egypt in order to investigate the truth of that...",
	})
}
