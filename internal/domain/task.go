package domain

import (
	"errors"

	"github.com/rickalon/FlowManagerAPI/internal/repositories"
)

type Task struct {
	Task_id   int    `json:"task_id"`
	Status    string `json:"status"`
	Content   string `json:"content"`
	ProyectId int    `json:"proyect_id"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func ValidateTask(t *Task) error {
	if t.ProyectId == 0 {
		return errors.New("proyect id is empty")
	}

	if t.Content == "" {
		return errors.New("content is empty")
	}
	return nil
}

func CreateTask(db *repositories.PqDB, t *Task) error {
	if t.Status == "" {
		_, err := db.DB.Exec("INSERT INTO TASKS (content,proyect_id,user_id) values ($1,$2,$3)", t.Content, t.ProyectId, t.UserId)
		return err
	} else {
		_, err := db.DB.Exec("INSERT INTO TASKS (content,proyect_id,user_id,status) values ($1,$2,$3,$4)", t.Content, t.ProyectId, t.UserId, t.Status)
		return err
	}
}

func GetTaskByIds(db *repositories.PqDB, t *Task) error {
	return db.DB.QueryRow("SELECT content,status,proyect_id,created_at from tasks").Scan(&t.Content, &t.Status, &t.ProyectId, &t.CreatedAt)
}
