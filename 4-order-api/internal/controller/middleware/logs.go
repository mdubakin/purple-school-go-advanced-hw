package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func WithJSONLogs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
			"method": r.Method,
			"url":    r.URL.Path,
			"code":   wrapper.StatusCode,
		}).Info("request")
	})
}
