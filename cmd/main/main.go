package main

import (
	"log"
	"net"
	"net/http"

	"github.com/Bojong-Ireng/backend-server/config"
	. "github.com/Bojong-Ireng/backend-server/internal/delivery/http"
)

func run() error {
	// Uncomment if not run on docker
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("failed to load .env: %s", err)
	// }

	srv := NewServer()
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	log.Printf("listening on http://%s\n", httpServer.Addr)
	err := httpServer.ListenAndServe()
	return err
}

func main() {
	if err := run(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening and serving: %s\n", err)
	}
}
