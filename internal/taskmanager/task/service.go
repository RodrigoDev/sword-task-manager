package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	l "github.com/RodrigoDev/sword-task-manager/internal/logging"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/model"
)

type TaskService struct {
	repository *Repository
}

func NewTaskService(repository *Repository) *TaskService {
	return &TaskService{
		repository,
	}
}

func (t *TaskService) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req model.Task
	ctx := r.Context()

	err := t.decode(r.Body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.repository.AddTask(ctx, req)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not create a task: %v", req))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *TaskService) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req model.Task

	ctx := r.Context()
	vars := mux.Vars(r)

	taskID, err := strconv.ParseInt(vars["taskID"], 10, 64)
	if err != nil {
		http.Error(w, "invalid taskID", http.StatusBadRequest)
		return
	}

	err = t.decode(r.Body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.repository.UpdateTask(ctx, taskID, req)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not update task: %v", req))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *TaskService) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	taskID, err := strconv.ParseInt(vars["taskID"], 10, 64)
	if err != nil {
		http.Error(w, "invalid taskID", http.StatusBadRequest)
		return
	}

	res, err := t.repository.FindTaskByID(ctx, taskID)
	if err != nil {
		l.Logger(ctx).Error(err.Error())
		http.Error(
			w,
			fmt.Sprintf("could not fetch requested task: %d", taskID),
			http.StatusInternalServerError,
		)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (t *TaskService) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := t.repository.GetAllTasks(ctx)
	if err != nil {
		l.Logger(ctx).Error(err.Error())
		http.Error(
			w,
			"could not fetch tasks",
			http.StatusInternalServerError,
		)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (t *TaskService) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	taskID, err := strconv.ParseInt(vars["taskID"], 10, 64)
	if err != nil {
		http.Error(w, "invalid taskID", http.StatusBadRequest)
		return
	}

	err = t.repository.DeleteTask(ctx, taskID)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not delete task with id: %v", taskID))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *TaskService) FindTasksByManagerID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	userID, err := strconv.ParseInt(vars["userID"], 10, 64)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		return
	}

	res, err := t.repository.FindTasksByManagerID(ctx, userID)
	if err != nil {
		l.Logger(ctx).Error(err.Error())
		http.Error(
			w,
			fmt.Sprintf("could not fetch tasks for user: %d", userID),
			http.StatusInternalServerError,
		)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (t *TaskService) decode(body io.ReadCloser, task *model.Task) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(task)
	if err != nil {
		return err
	}
	return nil
}
