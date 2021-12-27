package api

import (
	"github.com/dshum/school/internal/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func (s *Server) Routes() *mux.Router {
	router := s.router.StrictSlash(true)

	// Log handlers
	router.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})

	// Recovery errors
	router.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(false)))

	// Check JWT token
	authRoutes := router.PathPrefix("/api").Subrouter()
	authRoutes.Use(middlewares.Auth)

	router.HandleFunc("/", s.ApiStatus).Methods("GET")
	router.HandleFunc("/login", s.authService.Login).Methods("POST")
	router.HandleFunc("/logout", s.authService.Logout).Methods("GET", "POST")

	authRoutes.HandleFunc("/task_categories", s.taskCategoryService.GetTaskCategories).Methods("GET")
	authRoutes.HandleFunc("/task_categories/{task_category_id:[0-9]+}", s.taskCategoryService.GetTaskCategory).Methods("GET")

	return router
}
