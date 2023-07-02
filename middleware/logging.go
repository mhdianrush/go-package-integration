package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.WithFields(logrus.Fields{
			"URL":  r.Host + r.URL.Path,
			"size": r.Method,
		}).Info()
		next.ServeHTTP(w, r)
	})
}
