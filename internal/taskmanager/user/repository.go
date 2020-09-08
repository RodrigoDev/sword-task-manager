package user

import (
	"context"

	"github.com/RodrigoDev/sword-task-manager/internal/storage"
)

type Repository struct {
	client *storage.MySQLStorage
}

func (r *Repository) AddUser(user *User) error {
	conn, err := r.client.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(
		"INSERT INTO users "+
			"(name, password, email, type) "+
			"VALUES (?, ?, ?, ?);",
		user.Name,
		user.Password,
		user.Email,
		user.Type,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindUserByEmail(email string) (*User, error) {
	var user User

	conn, err := r.client.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM users WHERE email = ?", email)
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Type,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

type UserManager interface {
	AddUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, userID int64, user User) error
	DeleteUser(ctx context.Context, userID int64) error
	FindUserByID(ctx context.Context, userID int64) (*User, error)
	FindUserByEmail(ctx context.Context, userID int64) (*User, error)
	GetAllUsers(ctx context.Context) (*Users, error)
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