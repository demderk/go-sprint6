package main

import (
	serverLogs "go-sprint6/internal/logs"
	"go-sprint6/internal/server"
)

func main() {
	logger := serverLogs.Main
	server := server.NewServer(logger)
	server.Run()
}
