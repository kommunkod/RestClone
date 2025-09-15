package server

import (
	"net/http"

	"github.com/kommunkod/restclone/pkg/config"
)

type RestClone struct {
	Config *config.Server
}

var RC *RestClone

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
	listener, tlsListener, err := Init()
	if err != nil {
		panic(err)
	}

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

	RC.Config.Println("Server is running...")
	RC.Config.Println("Press Ctrl+C to stop the server.")
	for {
	}
}

func RunBackground(port string, tlsPort string) error {
	listener, tlsListener, err := Init()
	if err != nil {
		return err
	}

	listener.Addr = port
	tlsListener.Addr = tlsPort

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

	RC.Config.Println("Server is running in background...")
	RC.Config.Println("Press Ctrl+C to stop the server.")

	return nil
}

func Init() (http.Server, http.Server, error) {
	cfg, err := config.Init()
	if err != nil {
		return http.Server{}, http.Server{}, err
	}

	RC = &RestClone{
		Config: cfg,
	}

	tlsc, err := RC.Config.TLS.GetTLSConfig()
	if err != nil {
		return http.Server{}, http.Server{}, err
	}

	router := GetRouter()

	listener := http.Server{
		Addr:      RC.Config.Listen,
		Handler:   router,
		TLSConfig: tlsc,
	}

	tlsListener := http.Server{
		Addr:      RC.Config.ListenTLS,
		Handler:   router,
		TLSConfig: tlsc,
	}

	RC.Config.Println("Listening on", RC.Config.Listen)
	RC.Config.Println("[TLS] Listening on", RC.Config.ListenTLS)
	RC.Config.Println("[TLS] Certificate:", tlsc.Certificates[0].Leaf.Subject)

	return listener, tlsListener, nil
}
