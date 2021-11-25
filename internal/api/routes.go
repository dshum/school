package api

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (s *Server) Routes() *mux.Router {
	router := s.router

	router.HandleFunc("/", s.ApiStatus).Methods("GET")
	router.HandleFunc("/task_categories", s.taskCategoryService.GetTaskCategories).Methods("GET")
	router.HandleFunc("/task_categories/{task_category_id:[0-9]+}", s.taskCategoryService.GetTaskCategory).Methods("GET")

	router.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})
	router.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	return router
}
