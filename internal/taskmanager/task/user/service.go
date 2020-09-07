package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/model"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/storage"
)

type TaskService struct {
	storage *storage.MySQLStorage
}

func NewTaskService(storage *storage.MySQLStorage) *TaskService {
	return &TaskService{
		storage,
	}
}

func (t *TaskService) TaskHandler(w http.ResponseWriter, r *http.Request) {
	var req model.Task

	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		//add some log
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.addNewTask(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *TaskService) addNewTask(ctx context.Context, req model.Task) error {
	return nil
}
