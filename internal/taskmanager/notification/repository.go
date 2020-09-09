package notification

import (
	"context"

	"github.com/RodrigoDev/sword-task-manager/internal/storage"
)

type Repository struct {
	client *storage.MySQLStorage
}

func (r *Repository) AddNotification(notification *Notification) error {
	conn, err := r.client.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(
		"INSERT INTO notifications "+
			"(user_id, message) "+
			"VALUES (?, ?);",
		notification.UserID,
		notification.Message,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetNotificationsByUserID(userID int64) (*Notifications, error) {
	var notifications Notifications

	conn, err := r.client.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM notifications WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var notification Notification

		err = rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Message,
		)

		if err != nil {
			return nil, err
		}
	}
	return &notifications, nil
}

type NotificationManager interface {
	AddNotification(ctx context.Context, notification Notification) error
	FindNotificationByID(ctx context.Context, notificationID int64) (*Notification, error)
}

func NewRepository(client *storage.MySQLStorage) (*Repository, error) {
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