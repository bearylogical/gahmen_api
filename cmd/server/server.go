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
	router.HandleFunc("GET /api/v1/ministries", makeHttpHandleFunc(handler.ListMinistries))
	router.HandleFunc("GET /api/v1/ministries/{ministry_id}", makeHttpHandleFunc(handler.GetMinistryByID))
	router.HandleFunc("GET /api/v1/budget/{ministry_id}/documents", makeHttpHandleFunc(handler.ListDocumentByMinistryID))
	router.HandleFunc("GET /api/v1/budget", makeHttpHandleFunc(handler.ListExpenditure))
	router.HandleFunc("POST /api/v1/projects", makeHttpHandleFunc(handler.GetProjectExpenditureByQuery))
	router.HandleFunc("GET /api/v1/projects/{project_id}", makeHttpHandleFunc(handler.GetProjectExpenditureByID))
	router.HandleFunc("GET /api/v1/budget/opts", makeHttpHandleFunc(handler.GetBudgetOpts))
	router.HandleFunc("GET /api/v1/budget/{ministry_id}", makeHttpHandleFunc(handler.ListExpenditureByMinistry))
	router.HandleFunc("GET /api/v1/budget/{ministry_id}/projects", makeHttpHandleFunc(handler.ListProjectExpenditureByMinistryID))
	router.HandleFunc("GET /api/v1/sgdi/{ministry_id}/links", makeHttpHandleFunc(handler.ListSGDILinksByMinistryID))
	router.HandleFunc("GET /api/v1/personnel", makeHttpHandleFunc(handler.ListTopNPersonnelByMinistryID))

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
