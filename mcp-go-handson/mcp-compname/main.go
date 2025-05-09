package main

import (
	"log"

	opgenserver "github.com/Senoue/handson-2025/mcp-go-handson/mcp-compname/server"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	srv := opgenserver.NewServer()
	if err := server.ServeStdio(srv); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
