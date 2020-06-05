package server

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/handlers"
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
	s.engine.Run()
}
