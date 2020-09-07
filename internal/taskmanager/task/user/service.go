package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	l "github.com/RodrigoDev/sword-task-manager/internal/logging"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/model"
)

type TaskService struct {
	repository *task.Repository
}

func NewTaskService(repository *task.Repository) *TaskService {
	return &TaskService{
		repository,
	}
}

func (t *TaskService) TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		t.createTask(w, r)
		return
	}

	http.Error(w, "wrong method", http.StatusMethodNotAllowed)
}

func (t *TaskService) createTask(w http.ResponseWriter, r *http.Request) {
	var req model.Task
	ctx := r.Context()

	err := t.decode(r.Body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.repository.AddTask(ctx, req)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not add task to database: %v", req))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *TaskService) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req model.Task
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = t.decode(r.Body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.repository.UpdateTask(ctx, id, req)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not add task to database: %v", req))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *TaskService) decode(body io.ReadCloser, task *model.Task)  error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(task)
	if err != nil {
		return err
	}
	return nil
}