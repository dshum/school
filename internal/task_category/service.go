package task_category

import (
	"net/http"
	"strconv"

	"github.com/dshum/school/internal/utils"
	"github.com/gorilla/mux"
)

type Service interface {
	GetTaskCategories(w http.ResponseWriter, r *http.Request)
	GetTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskCategoryService struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &taskCategoryService{
		storage: storage,
	}
}

func (t *taskCategoryService) GetTaskCategories(w http.ResponseWriter, _ *http.Request) {
	// panic("Unexpected error!")

	if taskCategories, err := t.storage.GetList(); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		utils.JSONResponse(w, http.StatusOK, taskCategories)
	}
}

func (t *taskCategoryService) GetTaskCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["task_category_id"])

	if taskCategory, err := t.storage.Get(id); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		utils.JSONResponse(w, http.StatusOK, taskCategory)
	}
}
