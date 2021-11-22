package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dshum/school/models"
	"github.com/gorilla/mux"
)

var taskCategory models.TaskCategory

type TaskCategoryHandler struct{}

func (*TaskCategoryHandler) GetTaskCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if taskCategories, err := taskCategory.GetList(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(taskCategories)
	}
}

func (*TaskCategoryHandler) GetTaskCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["task_category_id"])

	w.Header().Set("Content-Type", "application/json")

	if _, err := taskCategory.Get(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(taskCategory)
	}
}
