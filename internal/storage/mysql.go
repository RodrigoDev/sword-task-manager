package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/RodrigoDev/sword-task-manager/internal/config"
)

type MySQLStorage struct {
	cfg *config.MySQL
}

func NewMySQLStorage(config *config.MySQL) *MySQLStorage {
	return &MySQLStorage{
		cfg: config,
	}
}

func (m *MySQLStorage) GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		m.cfg.User,
		m.cfg.Password,
		m.cfg.URI,
		m.cfg.Database,
	))

	if err != nil {
		return nil, err
	}

	return db, nil
}
