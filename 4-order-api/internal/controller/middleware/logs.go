package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func WithJSONLogs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
			"method":       r.Method,
			"url":          r.URL.Path,
			"code":         wrapper.StatusCode,
			"path":         r.URL.Path,
			"statusCode":   wrapper.StatusCode,
			"duration":     time.Since(start),
			"userAgent":    r.UserAgent(),
			"responseSize": wrapper.WroteBytes,
			"requestID":    r.Context().Value("requestID"),
		}).Info("request")
	})
}
