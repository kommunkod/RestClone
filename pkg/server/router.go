package server

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/kommunkod/restclone/pkg/api/v1"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/kommunkod/restclone/docs" // docs is generated by Swag CLI, you have to import it.
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API V1
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	v1.RegisterRoutes(apiRouter)

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods("GET")

	return router
}
