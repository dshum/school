package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dshum/school/config"
	"github.com/dshum/school/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
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

	router := gin.Default()

	var taskCategoryController controllers.TaskCategoryController
	router.GET("/task_categories", taskCategoryController.GetTaskCategories)
	router.GET("/task_categories/:task_category_id", taskCategoryController.GetTaskCategory)
	router.Run(":9991")
}
