package api

import (
	"github.com/dshum/school/internal/auth"
	"net/http"

	"github.com/dshum/school/internal/task_category"
	"github.com/gorilla/mux"
)

type Server struct {
	router              *mux.Router
	authService         auth.Service
	taskCategoryService task_category.Service
}

func NewServer(
	router *mux.Router,
	authService auth.Service,
	taskCategoryService task_category.Service,
) *Server {
	return &Server{
		router:              router,
		authService:         authService,
		taskCategoryService: taskCategoryService,
	}
}

func (s *Server) Run() error {
	router := s.Routes()

	if err := http.ListenAndServe(":9991", router); err != nil {
		return err
	}

	return nil
}
