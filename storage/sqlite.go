package storage

import (
	"database/sql"

	"github.com/ignoxx/sl3/checkout/types"
)

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(path string) *SqliteStorage {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        panic(err)
    }

    return &SqliteStorage{
        db: db,
    }
}

func (s *SqliteStorage) CreateOrder(order types.Order) error {
	return nil
}

func (s *SqliteStorage) GetOrder(id int) (types.Order, error) {
	return types.Order{}, nil
}

func (s *SqliteStorage) UpdateOrder(order types.Order) error {
	return nil
}

var _ Storage = (*SqliteStorage)(nil)
