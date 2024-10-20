package middleware

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/rizalta/urlshort/web/components"
)

type Middlware func(http.Handler) http.Handler

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func newWrappedWriter(w http.ResponseWriter) *wrappedWriter {
	return &wrappedWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           &bytes.Buffer{},
	}
}

func StackMiddleware(xs ...Middlware) Middlware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *wrappedWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v\n", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		wrapped := newWrappedWriter(w)
		next.ServeHTTP(wrapped, r)

		if wrapped.statusCode >= 400 {
			data := components.ErrorProps{
				Status: wrapped.statusCode,
				Error:  wrapped.body.String(),
			}
			w.Header().Set("Content-Type", "text/html")
			components.Error(data).Render(r.Context(), w)

		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := newWrappedWriter(w)
		next.ServeHTTP(wrapped, r)
		next.ServeHTTP(w, r)
		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
