package Stores

import (
	"database/sql"
	"phonebook-api/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// UserStore interface for server. Saves information about users using database or slice.
type UserStore interface {
	GetUser(int) (*model.User, error)
	GetAllUser() []*model.User
	AddUser(string, string)
	DeleteUser(int) error
}

// NewStore first implementation with array
func NewStore() UserStore {
	return &arrayUserStore{}
}

// NewTestStore test implementation with mysql database.
func NewTestStore() UserStore {
	db, err := sql.Open("mysql", "admin:password@/go_api")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	return &databaseUserStore{db}
}
