package main

import (
	"log"

	"github.com/dshum/school/config"
	"github.com/dshum/school/controllers"
	"github.com/dshum/school/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.InitializeDB()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		// Run all migrations
		migrations.Run()

		router := gin.Default()

		var noteController controllers.NoteController
		router.GET("/notes", noteController.GetAllNotes)
		router.POST("/notes", noteController.CreateNewNote)
		router.GET("/notes/:note_id", noteController.GetSingleNote)
		router.Run(":8000")
	}
}
