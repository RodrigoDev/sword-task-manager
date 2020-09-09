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

// Task endpoints
func Task(taskService *task.TaskService) Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/tasks", taskService.CreateTaskHandler).Methods("POST")
		h.router.HandleFunc("/tasks/{taskID}", taskService.GetTaskHandler).Methods("GET")
		h.router.HandleFunc("/tasks", taskService.GetAllTasksHandler).Methods("GET")
		h.router.HandleFunc("/tasks/{taskID}", taskService.UpdateTaskHandler).Methods("PUT")
		h.router.HandleFunc("/tasks/{taskID}", taskService.DeleteTaskHandler).Methods("DELETE")
		h.router.HandleFunc("/tasks/user/{userID}", taskService.FindTasksByManagerID).Methods("GET")

		return nil
	}
}

func User(service *user.Service) Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/login", service.LoginHandler).Methods("POST")
		h.router.HandleFunc("/register", service.CreateUserHandler).Methods("POST")

		h.router.Use(auth.JwtVerify)
		h.router.HandleFunc("/test", service.TestHandler).Methods("GET")

		return nil
	}
}

func Notification(service *notification.Service) Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/notifications", service.CreateNotificationHandler).Methods("POST")
		h.router.HandleFunc("/notifications/user/{userID}", service.CreateNotificationHandler).Methods("GET")
	}
}