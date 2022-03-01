package Stores

import (
	"database/sql"
	"go-api-friends/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// UserStore interface for server. Saves information about users using database or slice.
type UserStore interface {
	GetUser(int) (*model.User, error)
	GetAllUser() []*model.User
	AddUser(string, string)
	DeleteUser(int) error
	ValidateUser(string, string) *model.User
	UpdateUser(user *model.User)
}

// NewStore first implementation with array
//func NewStore() UserStore {
//	return &arrayUserStore{}
//}

// NewStore test implementation with mysql database.
func NewStore() UserStore {
	db, err := sql.Open("mysql", "admin:password@/go_api")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	return &databaseUserStore{db}
}
