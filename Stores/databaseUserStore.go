package Stores

import (
	"database/sql"
	"errors"
	"go-api-friends/model"
	"log"
)

// an implementation of UserStore but with database.
// query written for mysql database
type databaseUserStore struct {
	db *sql.DB
}

// GetUser returns user by his id or error if such user not presented.
func (us *databaseUserStore) GetUser(id int) (*model.User, error) {
	var name, status, city string
	err := us.db.QueryRow("select name, status, city  from Users where user_id = ?", id).Scan(&name, &status, &city)

	if err != nil {
		if err == sql.ErrNoRows {
			Log(err)
			return nil, errors.New("no such id")
		}
		Log(err)
		return nil, errors.New("troubles with base, my fault")
	}

	return &model.User{UserID: uint(id), Name: name, Status: status, City: city}, nil
}

// GetAllUser returns all users presented in database
func (us *databaseUserStore) GetAllUser() []*model.User {
	rows, err := us.db.Query("select user_id, name, status,city  from Users")

	if err != nil {
		Log(err)
		return make([]*model.User, 0)
	}

	defer rows.Close()

	var id uint
	var name, status, city string
	result := make([]*model.User, 0)

	for rows.Next() {
		err := rows.Scan(&id, &name, &status, &city)
		if err != nil {
			Log(err)
			return make([]*model.User, 0)
		}
		result = append(result, &model.User{UserID: id, Name: name, Status: status, City: city})
	}

	if err = rows.Err(); err != nil {
		Log(err)
		return make([]*model.User, 0)
	}

	return result
}

// AddUser insert new user to db
func (us *databaseUserStore) AddUser(name string, password string) {
	_, err := us.db.Exec("INSERT INTO Users (name, password, status, city) VALUES (?, ?, '', '')", name, password)

	if err != nil {
		Log(err)
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

func (us *databaseUserStore) UpdateUser(user *model.User) {
	_, err := us.db.Exec("UPDATE Users SET city=?, status=? WHERE user_id=?", user.City, user.Status, user.UserID)
	if err != nil {
		Log(err)
	}
}

func (us *databaseUserStore) ValidateUser(username string, password string) *model.User {
	query := us.db.QueryRow("SELECT user_id, name, status, city FROM Users WHERE name=? AND password=?", username, password)
	var id uint
	var status, city string
	if err := query.Scan(&id, &username, &status, &city); err != nil {
		return nil
	}
	return &model.User{UserID: id, Name: username, Status: status, City: city}
}
