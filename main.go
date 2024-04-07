package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	server := Server{
		addr:   ":8080",
		db:     InitDb(),
		logger: logger,
	}

	err := http.ListenAndServe(server.addr, server.routes())
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
