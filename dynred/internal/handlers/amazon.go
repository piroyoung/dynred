package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/log"
	"net/http"
)

type AmazonRedirectHandler struct {
	repo log.Repository
}

func NewAmazonRedirectHandler(repo log.Repository) RedirectHandler {
	return &AmazonRedirectHandler{repo: repo}
}

func (y *AmazonRedirectHandler) getUrl(c *gin.Context) string {
	return "https://amzn.to/" + c.Param("id")
}

func (y *AmazonRedirectHandler) HandleWithMeta(c *gin.Context) {
	err := y.repo.Dump(log.NewLog(c))
	if err != nil {
		panic(err)
	}
	handleWithMeta(y.getUrl(c), c)
}

func (y *AmazonRedirectHandler) HandleWith301(c *gin.Context) {
	err := y.repo.Dump(log.NewLog(c))
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusMovedPermanently, y.getUrl(c))
}
