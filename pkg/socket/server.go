package socket

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/kommunkod/restclone/pkg/server"
)

type SocketServer struct {
	Log      *log.Logger
	SockAddr string
}

func Run(socketPath string) error {
	srv := &SocketServer{
		Log:      log.New(os.Stdout, "[RestClone] ", log.LstdFlags),
		SockAddr: socketPath,
	}

	srv.Log.Println("Starting RestClone Server....")

	router := server.GetRouter()

	listener, err := net.ListenUnix("unix", &net.UnixAddr{
		Name: socketPath,
	})
	if err != nil {
		return fmt.Errorf("Error starting unix socket server: %w", err)
	}

	defer listener.Close()
	listener.SetUnlinkOnClose(true)

	serve := &http.Server{
		Addr:    socketPath,
		Handler: router,
	}

	srv.Log.Println("Server is running...")

	err = serve.Serve(listener)
	if err != nil {
		return fmt.Errorf("Error starting listener: %w", err)
	}

	srv.Log.Println("Server is listening on", socketPath)
	return nil
}
