package main

import (
	"log"
	"os"

	"github.com/kommunkod/restclone/pkg/socket"
)

func main() {
	// Usage
	// rcsocket [<socket path>]

	if len(os.Args) < 2 {
		log.Println("Usage: rcsocket [...flags] [<socket path>]")
		os.Exit(1)
	}

	socketPath := os.Args[len(os.Args)-1]

	log.Println("Starting RestClone Server Socket....")

	socket.Run(socketPath)
}
