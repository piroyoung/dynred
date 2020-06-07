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
	projectId := os.Getenv("PROJECT_ID")
	datasetId := os.Getenv("BQ_DATASET_ID")
	tableId := os.Getenv("BQ_TABLE_ID")
	repo, err := log.NewBigQueryRepository(projectId, datasetId, tableId)
	if err != nil {
		panic(err)
	}
	note := handlers.NewNoteRedirectHandler(noteUserId, repo)
	youtube := handlers.NewYouTubeRedirectHandler(repo)
	amazon := handlers.NewAmazonRedirectHandler(repo)
	beacon := handlers.NewBeaconHandler(repo)
	s := server.NewServer(engine, note, youtube, amazon, beacon)
	s.Run()
}
