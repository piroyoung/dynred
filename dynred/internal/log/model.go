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
	Host         string    `bigquery:"host"`
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
		Host:         c.Request.Host,
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
		domain, _ := getDomain(c.Request.Host)
		c.SetCookie("id", newId.String(), 86400*365, "/", domain, false, true)
		return newId.String()
	}
	return id
}

func getDomain(fullHost string) (string, error) {
	withoutPort := strings.Split(fullHost, ":")[0]
	parts := strings.Split(withoutPort, ".")
	var withoutSubdomain string
	if len(parts) == 1 {
		withoutSubdomain = parts[0]
	} else if len(parts) >= 2 {
		n := len(parts)
		withoutSubdomain = strings.Join(parts[n-2:], ".")
	}
	return withoutSubdomain, nil
}