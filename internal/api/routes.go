package api

import (
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/rizalta/urlshort/internal/handlers"
	"github.com/rizalta/urlshort/internal/repo"
	"github.com/rizalta/urlshort/web"
)

func SetupRoutes(mux *http.ServeMux, rdb *redis.Client) {
	r := repo.NewRepo(rdb)
	urlHandler := handlers.NewURLHandler(r)

	mux.Handle("/assets/", http.FileServerFS(web.FS))
	mux.HandleFunc("/{id}", urlHandler.RedirectToURL)
	mux.HandleFunc("/", urlHandler.HomePage)
	mux.HandleFunc("POST /shorten", urlHandler.AddURL)
}
