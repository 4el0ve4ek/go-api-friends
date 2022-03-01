package model

// User model. Similar to column in database.
type User struct {
	UserID uint   `json:"UserID"`
	Name   string `json:"Name"`
	Status string `json:"Status"`
	City   string `json:"City"`
}
