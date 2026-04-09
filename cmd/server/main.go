package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/l10-bhushan/notes_management/internal/router"
)

func main() {
	//log.Println("Notes Management system")

	// Creating an instance of config
	cfg := router.Config{
		Addr: ":8080",
	}

	// Creating an instance of application
	app := router.Application{
		Cfg: cfg,
	}

	// Starting the server
	if err := app.Run(router.Mount()); err != nil {
		slog.Error("Server failed to start")
		log.Printf("Server failed to start: %s", err)
		os.Exit(1)
	}
}
