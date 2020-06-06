package server

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/handlers"
	"net/http"
)

type Server struct {
	engine *gin.Engine
	note   handlers.RedirectHandler
}

func NewServer(engine *gin.Engine, note handlers.RedirectHandler) Server {
	return Server{engine, note}
}

func (s *Server) Run() {
	s.engine.LoadHTMLGlob("static/templates/*.tmpl")
	s.engine.GET("m/note/:articleId", s.note.HandleWithMeta)
	s.engine.GET("s/note/:articleId", s.note.HandleWith301)
	s.engine.GET("debug", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.tmpl", gin.H{})
	})
	s.engine.NoRoute(handlers.HandleNotFound)
	s.engine.Run()
}
