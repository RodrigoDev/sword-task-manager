package task

import (
	"context"

	"github.com/RodrigoDev/sword-task-manager/internal/storage"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/model"
)

type Repository struct {
	client *storage.MySQLStorage
}

type TaskManager interface {
	AddTask(ctx context.Context, task model.Task) error
	UpdateTask(ctx context.Context, taskID int64, task model.Task) error
	DeleteTask(ctx context.Context, taskID int64) error
	FindTaskByID(ctx context.Context, taskID int64) (*model.Task, error)
	FindTasksByManagerID(ctx context.Context, userID int64) (*model.Tasks, error)
	GetAllTasks(ctx context.Context) (*model.Tasks, error)
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
		"UPDATE tasks SET"+
			" assigned_to = ?, title = ?, summary = ?, done_at = NOW() "+
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
	conn, err := r.client.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM tasks WHERE id = ?", taskID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindTaskByID(ctx context.Context, taskID int64) (*model.Task, error) {
	var task model.Task

	conn, err := r.client.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM tasks WHERE id = ?", taskID)
	err = row.Scan(
		&task.ID,
		&task.UserID,
		&task.AssignedTo,
		&task.Title,
		&task.Summary,
		&task.CreatedAt,
		&task.DoneAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *Repository) FindTasksByManagerID(ctx context.Context, userID int64) (*model.Tasks, error) {
	var tasks model.Tasks

	conn, err := r.client.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task model.Task

		err = rows.Scan(
			&task.ID,
			&task.UserID,
			&task.AssignedTo,
			&task.Title,
			&task.Summary,
			&task.CreatedAt,
			&task.DoneAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func (r *Repository) GetAllTasks(ctx context.Context) (*model.Tasks, error) {
	var tasks model.Tasks

	conn, err := r.client.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task model.Task

		err = rows.Scan(
			&task.ID,
			&task.UserID,
			&task.AssignedTo,
			&task.Title,
			&task.Summary,
			&task.CreatedAt,
			&task.DoneAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
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
