package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rizalta/urlshort/internal/api"
	"github.com/rizalta/urlshort/internal/database"
	"github.com/rizalta/urlshort/internal/middleware"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	rdbAddr := os.Getenv("REDIS_ADDRESS")

	db := database.InitDB(rdbAddr)
	defer db.Close()

	mux := http.NewServeMux()
	api.SetupRoutes(mux, db)

	middlewareStack := middleware.StackMiddleware(
		middleware.Logging,
		middleware.ErrorHandler,
	)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: middlewareStack(mux),
	}

	fmt.Printf("Running server on port: %s\n", port)
	server.ListenAndServe()
}
