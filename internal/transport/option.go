package transport

import (
	"net/http"

	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task/user"
)

// Option is a Handler modifier.
type Option func(*Handler) error

// Health endpoints
func Health() Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/_live", func(w http.ResponseWriter, r *http.Request) {})
		return nil
	}
}

func Task(taskService *user.TaskService) Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/task", taskService.TaskHandler)
		h.router.HandleFunc("/task/{id}", taskService.UpdateTaskHandler)
		return nil
	}
}
