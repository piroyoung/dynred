package log

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Log struct {
	CookieId     string    `bigquery:"cookie_id"`
	LoggedAt     time.Time `bigquery:"logged_at"`
	Path         string    `bigquery:"path"`
	TrackingCode string    `bigquery:"tracking_code"`
	Query        string    `bigquery:"query"`
	Referer      string    `bigquery:"referer"`
	UserAgent    string    `bigquery:"user_agent"`
	Language     string    `bigquery:"language"`
}

func NewLog(c *gin.Context) Log {
	return Log{
		CookieId:     getId(c),
		LoggedAt:     time.Now(),
		Path:         c.Request.URL.Path,
		TrackingCode: c.Query("t"),
		Query:        c.Request.URL.RawQuery,
		Referer:      c.Request.Header.Get("Referer"),
		UserAgent:    c.Request.Header.Get("User-Agent"),
		Language:     c.Request.Header.Get("Accept-Language"),
	}
}

func getId(c *gin.Context) string {
	id, err := c.Cookie("id")
	if err != nil {
		newId, _ := uuid.NewUUID()
		domain := strings.Split(c.Request.Host, ":")[0]
		c.SetCookie("id", newId.String(), 86400*365, "/", domain, false, true)
		return newId.String()
	}
	return id
}