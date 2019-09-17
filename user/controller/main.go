package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../dto"
	"../model"
	"../service"
	"github.com/gorilla/mux"
)

var user = &model.User{}
var users = []*model.User{}
var userDto = &dto.UserDto{}
var userCreateDto = &dto.UserCreateDto{}

const version string = "v1"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("api/v1/user", findAllUsers).Methods("GET")
	router.HandleFunc("api/v1/user/{id}", findUserByID).Methods("GET")
	router.HandleFunc("api/v1/user/{username}", findUserByUsername).Methods("GET")
	router.HandleFunc("api/v1/user", saveUser).Methods("POST")
	router.HandleFunc("api/v1/user/{id}", updateUser).Methods("UPDATE")
	router.HandleFunc("api/v1/user/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func findUserByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	conv, _ := strconv.Atoi(id) // converting param id to int
	service.FindByID(conv)
	json.NewEncoder(response).Encode(user)
	response.WriteHeader(http.StatusOK) //resp status
}

func findUserByUsername(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	vars := mux.Vars(request)
	username := vars["username"]
	service.FindByUsername(username)
	json.NewEncoder(response).Encode(user)
	response.WriteHeader(http.StatusOK) //resp status
}

func findAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	service.FindAll()
	json.NewEncoder(response).Encode(&users)
	response.WriteHeader(http.StatusOK) //resp status
}

func saveUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	service.Save(userCreateDto)
	json.NewEncoder(response).Encode(user)
	response.WriteHeader(http.StatusCreated) //resp status
}

func updateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	conv, _ := strconv.Atoi(id) // converting param id to int
	if conv == user.ID {
		service.Update(userDto) //gambiarra provisória

		//i still need to make the rest of the code

		json.NewEncoder(response).Encode(user) // encoding user to json
		response.WriteHeader(http.StatusOK)    //resp status
	} else {
		response.WriteHeader(http.StatusNotFound) //resp status
	}

}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	conv, _ := strconv.Atoi(id)
	if conv == user.ID {
		service.Delete(conv)
		response.WriteHeader(http.StatusOK) //resp status
	} else {
		response.WriteHeader(http.StatusNotFound) //resp status
	}
}
