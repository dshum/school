package controllers

import (
	"net/http"
	"strconv"

	"github.com/dshum/school/models"
	"github.com/gin-gonic/gin"
)

type TaskCategoryController struct{}

func (*TaskCategoryController) GetTaskCategories(c *gin.Context) {
	var taskCategory models.TaskCategory
	task_categories, err := taskCategory.GetList()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"meta": "Task categories",
			"data": task_categories,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (*TaskCategoryController) GetTaskCategory(c *gin.Context) {
	var taskCategory models.TaskCategory
	id := c.Param("task_category_id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter ID",
			"error":   err.Error(),
		})
		return
	}

	_, err2 := taskCategory.Get(idInt)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
			"error":   err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": "Task category",
		"data": taskCategory,
	})
}
