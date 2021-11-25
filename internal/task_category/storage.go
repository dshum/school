package task_category

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type TaskCategoryStorage interface {
	GetList() ([]TaskCategory, error)
	Get(id int) (TaskCategory, error)
}

type taskCategoryStorage struct {
	db *pgx.Conn
}

func NewTaskCategoryStorage(db *pgx.Conn) TaskCategoryStorage {
	return &taskCategoryStorage{
		db: db,
	}
}

func (s *taskCategoryStorage) GetList() ([]TaskCategory, error) {
	allTaskCategories := []TaskCategory{}

	rows, err := s.db.Query(context.Background(), "SELECT id, name, url, \"order\", fullcontent, hide, created_at, updated_at FROM task_categories order by id")
	if err != nil {
		return allTaskCategories, err
	}

	for rows.Next() {
		var currentTaskCategory TaskCategory

		if err := rows.Scan(
			&currentTaskCategory.Id,
			&currentTaskCategory.Name,
			&currentTaskCategory.URL,
			&currentTaskCategory.Order,
			&currentTaskCategory.FullContent,
			&currentTaskCategory.Hide,
			&currentTaskCategory.CreatedAt,
			&currentTaskCategory.UpdatedAt); err == nil {
			allTaskCategories = append(allTaskCategories, currentTaskCategory)
		}
	}

	return allTaskCategories, err
}

func (s *taskCategoryStorage) Get(id int) (TaskCategory, error) {
	taskCategory := TaskCategory{}

	row := s.db.QueryRow(context.Background(),
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
