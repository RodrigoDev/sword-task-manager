package transport

import (
	"net/http"

	"github.com/RodrigoDev/sword-task-manager/internal/auth"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/notification"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/task"
	"github.com/RodrigoDev/sword-task-manager/internal/taskmanager/user"
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

// User endpoints
func User(service *user.Service) Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/login", service.LoginHandler).Methods("POST")
		h.router.HandleFunc("/register", service.CreateUserHandler).Methods("POST")

		return nil
	}
}

// Task endpoints
func Task(taskService *task.TaskService) Option {
	return func(h *Handler) error {
		t := h.router.PathPrefix("/tasks").Subrouter()
		t.Use(auth.JwtVerify)
		t.HandleFunc("/", taskService.CreateTaskHandler).Methods("POST")
		t.HandleFunc("/{taskID}", taskService.GetTaskHandler).Methods("GET")
		t.HandleFunc("/", taskService.GetAllTasksHandler).Methods("GET")
		t.HandleFunc("/{taskID}", taskService.UpdateTaskHandler).Methods("PUT")
		t.HandleFunc("/{taskID}", taskService.DeleteTaskHandler).Methods("DELETE")
		t.HandleFunc("/user/{userID}", taskService.FindTasksByManagerID).Methods("GET")

		return nil
	}
}

func Notification(service *notification.Service) Option {
	return func(h *Handler) error {
		n := h.router.PathPrefix("/notifications").Subrouter()
		n.Use(auth.JwtVerify)
		n.HandleFunc("/", service.CreateNotificationHandler).Methods("POST")
		n.HandleFunc("/user/{userID}", service.CreateNotificationHandler).Methods("GET")

		return nil
	}
}