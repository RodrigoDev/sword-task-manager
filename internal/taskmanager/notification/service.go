package notification

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	l "github.com/RodrigoDev/sword-task-manager/internal/logging"
)

type ErrorResponse struct {
	Err string
}

type Service struct {
	repository *Repository
}

func NewNotificationService(repository *Repository) *Service {
	return &Service{
		repository,
	}
}

func (s *Service) CreateNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var req Notification
	ctx := r.Context()

	err := s.decode(r.Body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.repository.AddNotification(&req)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not create a notifications: %v", req))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) GetNotificationsByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	userID, err := strconv.ParseInt(vars["userID"], 10, 64)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		return
	}
	res, err := s.repository.GetNotificationsByUserID(userID)
	if err != nil {
		l.Logger(ctx).Error(err.Error())
		http.Error(
			w,
			"could not fetch notifications",
			http.StatusInternalServerError,
		)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (s *Service) decode(body io.ReadCloser, notification *Notification) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(notification)
	if err != nil {
		return err
	}
	return nil
}
