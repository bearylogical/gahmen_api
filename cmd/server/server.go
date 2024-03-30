package server

import (
	"log"
	"net/http"

	"gahmen-api/config"
	storage "gahmen-api/db"
	"gahmen-api/handlers"
	"gahmen-api/helpers"
	"gahmen-api/middleware"
)

type APIServer struct {
	listenAddr string
	store      storage.Storage
	config     config.Config
}

func NewAPIServer(listenAddr string, store storage.Storage, config *config.Config) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
		config:     *config,
	}
}

func (s *APIServer) Run() {

	handler := handlers.NewHandler(s.store)
	router := http.NewServeMux()

	stack := middleware.CreateStack(middleware.Logging, middleware.AllowCors)

	// high level route to list all ministries and get a specific ministry by id
	router.HandleFunc("GET /ministry", makeHttpHandleFunc(handler.ListMinistry))
	router.HandleFunc("GET /ministry/{ministry_id}", makeHttpHandleFunc(handler.GetMinistryByID))
	router.HandleFunc("GET /ministry/{ministry_id}/budget/documents", makeHttpHandleFunc(handler.ListDocumentByMinistryID))
	router.HandleFunc("GET /ministry/{ministry_id}/budget/projects", makeHttpHandleFunc(handler.ListProjectByMinistryID))
	router.HandleFunc("GET /ministry/{ministry_id}/sgdi/links", makeHttpHandleFunc(handler.ListSGDILinksByMinistryID))

	log.Print("JSON API server running on port", s.listenAddr)
	server := http.Server{
		Addr:    s.listenAddr,
		Handler: stack(router),
	}
	server.ListenAndServe()
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHttpHandleFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
