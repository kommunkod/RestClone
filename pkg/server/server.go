package server

import (
	"net/http"

	"github.com/kommunkod/restclone/pkg/config"
)

type RestClone struct {
	Config *config.Server
}

// @title RestClone API
// @version 1.0
// @description RestClone is a modified stateless Rest API for RClone.
// @termsOfService https://github.com/kommunkod/RestClone

// @contact.name Hoglandets IT
// @contact.email support@hoglandet.se

// @license.name GPL-3.0
// @license.url https://www.gnu.org/licenses/gpl-3.0.html

// @BasePath /

// @securityDefinitions.basicAuth BasicAuth
func Run() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	server := &RestClone{
		Config: cfg,
	}

	tlsc, err := server.Config.TLS.GetTLSConfig()
	if err != nil {
		panic(err)
	}

	router := GetRouter()

	listener := http.Server{
		Addr:      server.Config.Listen,
		Handler:   router,
		TLSConfig: tlsc,
	}

	tlsListener := http.Server{
		Addr:      server.Config.ListenTLS,
		Handler:   router,
		TLSConfig: tlsc,
	}

	server.Config.Println("Listening on", server.Config.Listen)
	server.Config.Println("[TLS] Listening on", server.Config.ListenTLS)
	server.Config.Println("[TLS] Certificate:", tlsc.Certificates[0].Leaf.Subject)

	server.Config.Println("Server is starting...")

	go func() {
		if err := listener.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := tlsListener.ListenAndServeTLS("", ""); err != nil {
			panic(err)
		}
	}()

	server.Config.Println("Server is running...")
	server.Config.Println("Press Ctrl+C to stop the server.")
	for {
	}
}
