package main

import (
	"fmt"
	"log"
	"net/http"
	"phonebook-api/controller"

	"github.com/gorilla/mux"
)

// fun
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	server := controller.NewPhoneServer()
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/user", server.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", server.GetUserById).Methods("GET")
	router.HandleFunc("/user", server.AddUser).Methods("POST")
	router.HandleFunc("/city/{city}", server.GetUserFromCity).Methods("GET")
	router.HandleFunc("/friend", server.AddRelations).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
