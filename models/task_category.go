package models

import (
	"context"
	"time"

	"github.com/dshum/school/config"
)

type TaskCategory struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	Order       int       `json:"order"`
	FullContent *string   `json:"fullcontent"`
	Hide        bool      `json:"hide"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (*TaskCategory) GetList() ([]TaskCategory, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, name, url, \"order\", fullcontent, hide, created_at, updated_at FROM task_categories order by id")
	allTaskCategories := []TaskCategory{}
	if err != nil {
		return allTaskCategories, err
	}

	for rows.Next() {
		var currentTaskCategory TaskCategory
		err2 := rows.Scan(
			&currentTaskCategory.Id,
			&currentTaskCategory.Name,
			&currentTaskCategory.URL,
			&currentTaskCategory.Order,
			&currentTaskCategory.FullContent,
			&currentTaskCategory.Hide,
			&currentTaskCategory.CreatedAt,
			&currentTaskCategory.UpdatedAt)
		if err2 != nil {
			return allTaskCategories, err2
		}
		allTaskCategories = append(allTaskCategories, currentTaskCategory)
	}

	return allTaskCategories, err
}

func (taskCategory *TaskCategory) Get(id int) (*TaskCategory, error) {
	row := config.DB.QueryRow(context.Background(),
		"SELECT id, name, url, \"order\", fullcontent, hide, created_at, updated_at FROM task_categories WHERE id = $1", id)
	err := row.Scan(
		&taskCategory.Id,
		&taskCategory.Name,
		&taskCategory.URL,
		&taskCategory.Order,
		&taskCategory.FullContent,
		&taskCategory.Hide,
		&taskCategory.CreatedAt,
		&taskCategory.UpdatedAt)

	return taskCategory, err
}
