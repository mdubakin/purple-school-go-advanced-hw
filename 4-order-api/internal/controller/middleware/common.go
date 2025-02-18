package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
	WroteBytes int
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func (w *WrapperWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.WroteBytes += n
	return n, err
}

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range middlewares {
			next = m(next)
		}
		return next
	}
}
