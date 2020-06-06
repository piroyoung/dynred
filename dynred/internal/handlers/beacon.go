package handlers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/piroyoung/dynred/internal/log"
	"net/http"
)

type BeaconHandler struct {
	repo log.Repository
	img  *[]byte
}

func NewBeaconHandler(repo log.Repository) BeaconHandler {
	b64MinimumPng := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVQI12NgYAAAAAMAASDVlMcAAAAASUVORK5CYII="
	img, err := base64.StdEncoding.DecodeString(b64MinimumPng)
	if err != nil {
		panic(err)
	}
	return BeaconHandler{repo: repo, img: &img}
}

func (b *BeaconHandler) Handle(c *gin.Context) {
	err := b.repo.Dump(log.NewLog(c))
	if err != nil {
		panic(err)
	}
	c.Header("Cache-Control", "no-cache")
	c.Data(http.StatusOK, "image/png", *b.img)
}
