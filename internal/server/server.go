package server

import (
	"fmt"
	"go-sprint6/internal/handlers"
	serverLogs "go-sprint6/internal/logs"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Log        *log.Logger
	HttpServer *http.Server
}

func NewServer(log *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.GetMain)
	router.HandleFunc("/upload", handlers.PostUpload)

	var s Server

	s.HttpServer = &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     log,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	s.Log = log

	return &s
}

func (s *Server) Run() {
	serverLogs.Main.Println("Starting server")
	if err := s.HttpServer.ListenAndServe(); err != nil {
		s.Log.Fatal(fmt.Errorf("failed to run server: %w", err))
		return
	}
}
