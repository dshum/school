package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dshum/school/config"
	"github.com/dshum/school/controllers"
	"github.com/gorilla/mux"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	conn, err := config.InitializeDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := mux.NewRouter()
	var taskCategoryController controllers.TaskCategoryController
	r.HandleFunc("/task_categories", taskCategoryController.GetTaskCategories).Methods("GET")
	r.HandleFunc("/task_categories/{task_category_id:[0-9]+}", taskCategoryController.GetTaskCategory).Methods("GET")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9991", r))
}
