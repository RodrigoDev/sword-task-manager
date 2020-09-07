package task

import (
	"context"

	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/model"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/storage"
)

type Repository struct {
	client *storage.MySQLStorage
}

type TaskManager interface {
	AddTask(ctx context.Context, task model.Task) error
	UpdateTask(ctx context.Context, taskID int64, task model.Task) error
	DeleteTask(ctx context.Context, taskID int64) error
	FindTaskByID(ctx context.Context, taskID int64) error
	FindTasksByUserID(ctx context.Context, UserID int64) error
}

func (r *Repository) AddTask(ctx context.Context, task model.Task) error {
	conn, err := r.client.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(
		"INSERT INTO tasks "+
			"(user_id, assigned_to, title, summary) "+
			"VALUES (?, ?, ?, ?);",
		task.UserID,
		task.AssignedTo,
		task.Title,
		task.Summary,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTask(ctx context.Context, taskID int64, task model.Task) error {
	conn, err := r.client.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(
		"UPDATE tasks SET" +
			" assigned_to = ?, title = ?, summary = ?, done_at = NOW() " +
			" WHERE id = ?",
		task.AssignedTo,
		task.Title,
		task.Summary,
		taskID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTask(ctx context.Context, taskID int64) error {
	panic("implement me")
}

func (r *Repository) FindTaskByID(ctx context.Context, taskID int64) error {
	panic("implement me")
}

func (r *Repository) FindTasksByUserID(ctx context.Context, UserID int64) error {
	panic("implement me")
}

func NewTaskRepository(client *storage.MySQLStorage) (*Repository, error) {
	r := &Repository{
		client: client,
	}

	conn, err := r.client.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return r, nil
}
