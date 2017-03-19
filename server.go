package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"./connects"
	"./structures"
)


//Función Main
func main(){
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	mux := mux.NewRouter()
	mux.HandleFunc("/Chat_ms/Api/Message/{userId}", GetMessages).Methods("GET")
	mux.HandleFunc("/Chat_ms/Api/Message", AddMessage).Methods("POST")

	http.Handle("/", mux)
	log.Println("El servidor encuentra en el puerto 4005")
	log.Fatal(http.ListenAndServe(":4005", nil))
}


//Funciones GET
func GetMessages(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userId"]

	status := "success"
	var message string
	chat := connect.GetMessages(userId)

	if(chat.Id <= 0){
		status = "Error"
		message = "No existe este chat"
	}
	response := structures.Response{status, chat, message}

	json.NewEncoder(w).Encode(response)
}

//Funciones POST
func AddMessage(w http.ResponseWriter, r *http.Request) {
	//Message := Message{"Probando", 1, 2}
}