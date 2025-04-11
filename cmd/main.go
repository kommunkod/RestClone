package main

import (
	"log"

	"github.com/kommunkod/restclone/pkg/server"
)

func main() {
	log.Println("Starting RestClone Server....")

	server.Run()
}
