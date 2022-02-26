package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"phonebook-api/model"
	"strconv"
)

// GetUsers returns to sender all saved users using json format.
func (ps *PhoneServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, ps.store.GetAllUser())
}

// GetUserById returns to sender user with id from request.
func (ps *PhoneServer) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := ps.store.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(w, user)
}

// writes to response json of object v.
func jsonResponse(w http.ResponseWriter, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// AddUser adds new user with name and city from body.
func (ps *PhoneServer) AddUser(w http.ResponseWriter, r *http.Request) {
	type RequestUser struct {
		Name string `json:"name"`
		City string `json:"city"`
	}

	var user RequestUser
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ps.store.AddUser(user.Name, user.City)
}

// GetUserFromCity returns users from certain city.
func (ps *PhoneServer) GetUserFromCity(w http.ResponseWriter, r *http.Request) {
	users := ps.store.GetAllUser()
	requiredCity := mux.Vars(r)["city"]
	result := make([]*model.User, 8)
	for _, user := range users {
		if user.City == requiredCity {
			result = append(result, user)
		}
	}
	jsonResponse(w, result)
}

// AddRelations will make two users friends.
func (ps *PhoneServer) AddRelations(w http.ResponseWriter, r *http.Request) {

}
