package main

import "C"

import (
	"fmt"
	"log"

	"github.com/kommunkod/restclone/pkg/socket"
)

//export StartSocket
func StartSocket(socketPath *C.char) *C.char {
	sockpath := C.GoString(socketPath)
	log.Println("Starting RestClone Server Socket....")

	err := socket.Run(sockpath)
	if err != nil {
		data := fmt.Sprintf("Error: %s", err)
		return C.CString(data)
	}

	return C.CString("Socket server finished")
}

func main() {
	// This is required for the cgo build to work
	// but we don't need to do anything here
}
