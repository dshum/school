package main

import (
	"context"
	"log"
	"os"

	"github.com/dshum/school/internal/api"
	"github.com/dshum/school/internal/config"
	"github.com/dshum/school/internal/db"
	"github.com/dshum/school/internal/task_category"
	"github.com/gorilla/mux"
)

func run() error {
	if err := config.LoadEnv(); err != nil {
		return err
	}

	config := config.NewConfig()

	db, err := db.NewConnection(config)
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	taskCategoryStorage := task_category.NewTaskCategoryStorage(db)
	taskCategoryService := task_category.NewTaskCategoryService(taskCategoryStorage)

	router := mux.NewRouter()
	server := api.NewServer(
		router,
		taskCategoryService,
	)

	if err = server.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
