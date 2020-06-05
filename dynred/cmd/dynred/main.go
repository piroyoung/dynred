package main

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/handlers"
	"github.com/piroyoung/dynred/internal/server"
	"os"
)

func main() {
	engine := gin.Default()
	noteUserId := os.Getenv("NOTE_USER_ID")
	note := handlers.NewNoteRedirectHandler(noteUserId)
	s := server.NewServer(engine, note)
	s.Run()
}
