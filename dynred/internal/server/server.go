package server

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/handlers"
	"net/http"
)

type Server struct {
	engine  *gin.Engine
	note    handlers.RedirectHandler
	youtube handlers.RedirectHandler
	amazon  handlers.RedirectHandler
	beacon  handlers.BeaconHandler
}

func NewServer(engine *gin.Engine, note handlers.RedirectHandler, youtube handlers.RedirectHandler, amazon handlers.RedirectHandler, beacon handlers.BeaconHandler) Server {
	return Server{engine, note, youtube, amazon, beacon}
}

func (s *Server) Run() {
	s.engine.LoadHTMLGlob("static/templates/*.tmpl")
	s.engine.GET("b.png", s.beacon.Handle)
	s.engine.GET("m/note/:articleId", s.note.HandleWithMeta)
	s.engine.GET("s/note/:articleId", s.note.HandleWith301)
	s.engine.GET("m/yt/:id", s.youtube.HandleWithMeta)
	s.engine.GET("s/yt/:id", s.youtube.HandleWith301)
	s.engine.GET("m/am/:id", s.amazon.HandleWithMeta)
	s.engine.GET("s/am/:id", s.amazon.HandleWith301)
	s.engine.GET("debug", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.tmpl", gin.H{})
	})
	s.engine.NoRoute(handlers.HandleNotFound)
	s.engine.Run()
}
