package main

import (
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/handlers"
	"github.com/piroyoung/dynred/internal/log"
	"github.com/piroyoung/dynred/internal/server"
	"os"
)

func main() {
	engine := gin.Default()
	noteUserId := os.Getenv("NOTE_USER_ID")
	repo, err := log.NewBigQueryRepository("pizu-279501", "bi", "tracking_logs")
	if err != nil {
		panic(err)
	}
	note := handlers.NewNoteRedirectHandler(noteUserId, repo)
	s := server.NewServer(engine, note)
	s.Run()
}
