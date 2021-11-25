package task_category

import "time"

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
