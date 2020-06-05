package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NoteRedirectHandler struct {
	userId string
}

func NewNoteRedirectHandler(userId string) RedirectHandler {
	return NoteRedirectHandler{userId: userId}
}

func (n NoteRedirectHandler) getUrl(c *gin.Context) string {
	return "https://note.com/" + n.userId + "/n/" + c.Param("articleId")
}

func (n NoteRedirectHandler) HandleWithMeta(c *gin.Context) {
	handleWithMeta(n.getUrl(c), c)
}

func (n NoteRedirectHandler) HandleWith301(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, n.getUrl(c))
}
