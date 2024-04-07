package main

import (
	"log/slog"

	"gorm.io/gorm"
)

type Server struct {
	addr   string
	db     *gorm.DB
	logger *slog.Logger
}

func NewServer(addr string, db *gorm.DB, logger *slog.Logger) *Server {
	return &Server{
		addr:   addr,
		db:     db,
		logger: logger,
	}
}
