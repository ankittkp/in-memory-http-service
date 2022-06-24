package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/jinxankit/in-memory-http-service/internal/handlers"
)

type Application struct {
	port int
}

func NewApplication() (*Application, error) {
	return &Application{
		port: 8080,
	}, nil
}

func (a Application) StartHTTPServer() error {
	// Add handlers
	handler := handlers.NewHandler()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/", handler.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/get/{key}", handler.GetValue).Methods("GET")
	r.HandleFunc("/api/v1/set", handler.SetValue).Methods("POST")
	r.HandleFunc("/api/v1/search", handler.Search).Queries("prefix", "{str}").Methods("GET")
	r.HandleFunc("/api/v1/search", handler.Search).Queries("suffix", "{str}").Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
	//TODO: Register Api group to Auth Middleware
	return nil
}
