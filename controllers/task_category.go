package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dshum/school/models"
	"github.com/gorilla/mux"
)

var taskCategory models.TaskCategory

type TaskCategoryController struct{}

func (*TaskCategoryController) GetTaskCategories(w http.ResponseWriter, r *http.Request) {
	taskCategories, err := taskCategory.GetList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(taskCategories)
	}
}

func (*TaskCategoryController) GetTaskCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["task_category_id"])

	_, err := taskCategory.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(taskCategory)
	}
}
