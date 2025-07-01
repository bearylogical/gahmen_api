package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gahmen-api/config"
	storage "gahmen-api/db"
	"gahmen-api/handlers"
	"gahmen-api/helpers"
	"gahmen-api/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
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

	// Public routes
	publicRouter := http.NewServeMux()
	publicRouter.HandleFunc("GET /health", makeHttpHandleFunc(handler.HealthCheck))

	// Authenticated routes
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /api/v1/ministries", makeHttpHandleFunc(handler.ListMinistries))
	authRouter.HandleFunc("GET /api/v1/ministries/{ministry_id}", makeHttpHandleFunc(handler.GetMinistryByID))
	authRouter.HandleFunc("GET /api/v1/budget/{ministry_id}/documents", makeHttpHandleFunc(handler.ListDocumentByMinistryID))
	authRouter.HandleFunc("GET /api/v1/budget", makeHttpHandleFunc(handler.ListExpenditure))
	authRouter.HandleFunc("GET /api/v1/projects/{project_id}", makeHttpHandleFunc(handler.GetProjectExpenditureByID))
	authRouter.HandleFunc("GET /api/v1/budget/{ministry_id}/programmes", makeHttpHandleFunc(handler.GetProgrammeExpenditureByMinistryID))
	authRouter.HandleFunc("GET /api/v1/budget/opts", makeHttpHandleFunc(handler.GetBudgetOpts))
	authRouter.HandleFunc("GET /api/v1/budget/{ministry_id}", makeHttpHandleFunc(handler.ListExpenditureByMinistry))
	authRouter.HandleFunc("GET /api/v1/budget/{ministry_id}/projects", makeHttpHandleFunc(handler.ListProjectExpenditureByMinistryID))
	authRouter.HandleFunc("GET /api/v1/sgdi/{ministry_id}/links", makeHttpHandleFunc(handler.ListSGDILinksByMinistryID))
	authRouter.HandleFunc("GET /api/v1/personnel", makeHttpHandleFunc(handler.ListTopNPersonnelByMinistryID))
	authRouter.HandleFunc("GET /api/v2/budget", makeHttpHandleFunc(handler.GetMinistryDataV2))
	authRouter.HandleFunc("POST /api/v2/projects", makeHttpHandleFunc(handler.GetProjectExpenditureByQuery))

	// Apply common middleware to both public and authenticated routes
	commonStack := middleware.CreateStack(middleware.Logging, middleware.AllowCors, middleware.RateLimit)

	// Apply AuthMiddleware only to authenticated routes
	authHandler := middleware.AuthMiddleware(&s.config)(authRouter)

	// Main router
	mainRouter := http.NewServeMux()
	mainRouter.Handle("/health", commonStack(publicRouter))
	mainRouter.Handle("/api/", commonStack(authHandler))
	mainRouter.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/docs/swagger.json")))
	mainRouter.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, filepath.Join(wd, "docs", "swagger.json"))
	})

	log.Print("JSON API server running on port", s.listenAddr)
	server := http.Server{
		Addr:    s.listenAddr,
		Handler: mainRouter,
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
