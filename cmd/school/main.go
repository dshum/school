package main

import (
	"context"
	"github.com/dshum/school/internal/auth"
	"log"
	"os"

	"github.com/dshum/school/internal/api"
	"github.com/dshum/school/internal/db"
	"github.com/dshum/school/internal/redis"
	"github.com/dshum/school/internal/task_category"
	"github.com/gorilla/mux"
)

func run() error {
	db, err := db.NewConnection()
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	redis, err := redis.NewConnection()
	if err != nil {
		return err
	}

	authStorage := auth.NewStorage(db)
	authService := auth.NewService(authStorage, redis)
	taskCategoryStorage := task_category.NewStorage(db)
	taskCategoryService := task_category.NewService(taskCategoryStorage)

	router := mux.NewRouter()
	server := api.NewServer(
		router,
		authService,
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
