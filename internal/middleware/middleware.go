package middleware

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"
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

		type errorData struct {
			Status int
			Error  string
		}

		if wrapped.statusCode >= 400 {
			data := errorData{
				Status: wrapped.statusCode,
				Error:  wrapped.body.String(),
			}
			w.Header().Set("Content-Type", "text/html")
			tmpl := template.Must(template.ParseFiles("web/templates/error.html"))
			tmpl.Execute(w, data)
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
