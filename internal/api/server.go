package api

import (
	"net/http"

	"github.com/dshum/school/internal/task_category"
	"github.com/gorilla/mux"
)

type Server struct {
	router              *mux.Router
	taskCategoryService task_category.TaskCategoryService
}

func NewServer(router *mux.Router, taskCategoryService task_category.TaskCategoryService) *Server {
	return &Server{
		router:              router,
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
