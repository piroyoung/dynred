package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/log"
	"net/http"
)

type NoteRedirectHandler struct {
	userId string
	repo   log.Repository
}

func NewNoteRedirectHandler(userId string, repo log.Repository) *NoteRedirectHandler {
	return &NoteRedirectHandler{userId: userId, repo: repo}
}

func (n *NoteRedirectHandler) getUrl(c *gin.Context) string {
	return "https://note.com/" + n.userId + "/n/" + c.Param("articleId")
}

func (n *NoteRedirectHandler) HandleWithMeta(c *gin.Context) {
	err := n.repo.Dump(log.NewLog(c))
	if err != nil {
		panic(err)
	}
	handleWithMeta(n.getUrl(c), c)
}

func (n *NoteRedirectHandler) HandleWith301(c *gin.Context) {
	err := n.repo.Dump(log.NewLog(c))
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusMovedPermanently, n.getUrl(c))
}
