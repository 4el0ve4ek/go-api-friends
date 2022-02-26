package controller

import (
	Stores "phonebook-api/Stores"
)

// PhoneServer contains function which will work on each routes
type PhoneServer struct {
	store Stores.UserStore
}

// NewPhoneServer constructs PhoneServer
func NewPhoneServer() *PhoneServer {
	return &PhoneServer{Stores.NewTestStore()}
}
