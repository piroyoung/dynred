package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/log"
	"net/http"
)

type YouTubeRedirectHandler struct {
	repo log.Repository
}

func NewYouTubeRedirectHandler(repo log.Repository) RedirectHandler {
	return &YouTubeRedirectHandler{repo: repo}
}

func (y *YouTubeRedirectHandler) getUrl(c *gin.Context) string {
	return "https://youtu.be/" + c.Param("id")
}

func (y *YouTubeRedirectHandler) HandleWithMeta(c *gin.Context) {
	err := y.repo.Dump(log.NewLog(c))
	if err != nil {
		panic(err)
	}
	handleWithMeta(y.getUrl(c), c)
}

func (y *YouTubeRedirectHandler) HandleWith301(c *gin.Context) {
	err := y.repo.Dump(log.NewLog(c))
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusMovedPermanently, y.getUrl(c))
}