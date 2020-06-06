package log

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Log struct {
	id        string
	loggedAt  int64
	path      string
	tag       string
	query     string
	referer   string
	userAgent string
	language  string
}

func NewLog(c *gin.Context) Log {
	return Log{
		id:        c.Request.Header.Get("Cookie"),
		loggedAt:  time.Now().Unix(),
		path:      c.Request.URL.Path,
		tag:       c.Query("t"),
		query:     c.Request.URL.RawQuery,
		referer:   c.Request.Header.Get("Referer"),
		userAgent: c.Request.Header.Get("User-Agent"),
		language:  c.Request.Header.Get("Accept-Language"),
	}
}
