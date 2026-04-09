package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// We start by creating 3 structs each for DbConfig, Config and Application i.e. Server.

// Database Configuration struct: This struct would have the database connection string
type DbConfig struct {
	Dsn string
}

// Server Configuration struct: This will store Port no and database configuration
type Config struct {
	Addr string
	Db   DbConfig
}

// Application: this struct would have configuration as dependency
type Application struct {
	Cfg Config
}

// Mount function is to initialize and mount the routes
func Mount() http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy ✅"))
	})

	return router
}

// Run function will start the server
func (app *Application) Run(router http.Handler) error {
	srv := &http.Server{
		Addr:         app.Cfg.Addr,
		Handler:      router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Println("Server running on PORT :8080 🚀")

	return srv.ListenAndServe()
}
