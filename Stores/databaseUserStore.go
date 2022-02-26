package Stores

import (
	"database/sql"
	"errors"
	"log"
	"phonebook-api/model"
)

// an implementation of UserStore but with database.
// query written for mysql database
type databaseUserStore struct {
	db *sql.DB
}

// GetUser returns user by his id or error if such user not presented.
func (us *databaseUserStore) GetUser(id int) (*model.User, error) {
	rows, err := us.db.Query("select user_id, name, city  from Users where user_id = ?", id)
	defer rows.Close()

	if err != nil {
		Log(err)
		return nil, errors.New("troubles with base, my fault")
	}

	var name, city string

	for rows.Next() {
		err := rows.Scan(&id, &name, &city)
		if err != nil {
			Log(err)
			return nil, errors.New("troubles with base, my fault")
		}
		return &model.User{UserID: uint(id), Name: name, City: city}, nil

	}

	if err = rows.Err(); err != nil {
		Log(err)
		return nil, errors.New("troubles with base, my fault")
	}

	return nil, errors.New("no such id")
}

// GetAllUser returns all users presented in database
func (us *databaseUserStore) GetAllUser() []*model.User {
	rows, err := us.db.Query("select user_id, name, city  from Users")

	if err != nil {
		Log(err)
		return make([]*model.User, 0)
	}

	defer rows.Close()

	var id uint
	var name, city string
	result := make([]*model.User, 0)

	for rows.Next() {
		err := rows.Scan(&id, &name, &city)
		if err != nil {
			Log(err)
			return make([]*model.User, 0)
		}
		result = append(result, &model.User{UserID: id, Name: name, City: city})
	}

	if err = rows.Err(); err != nil {
		Log(err)
		return make([]*model.User, 0)
	}

	return result
}

// AddUser insert new user to db
func (us *databaseUserStore) AddUser(name string, city string) {
	_, err := us.db.Exec("INSERT INTO Users (name, city) VALUES (?, ?)", name, city)

	if err != nil {
		Log(err)
		return
	}
}

// Log saves info about errors which occurred during work with db
func Log(err error) {
	log.Println(err)
}

// DeleteUser not implemented
func (us *databaseUserStore) DeleteUser(id int) error {
	return errors.New("no such id")
}
