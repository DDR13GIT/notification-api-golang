package main

import (
	"log"
	"net/http"
	"notification-api-golang/database"
	"notification-api-golang/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	r := mux.NewRouter()

	r.HandleFunc("/notifications", handlers.CreateNotification).Methods("POST")
	r.HandleFunc("/notifications/{id}", handlers.GetNotification).Methods("GET")
	r.HandleFunc("/notifications", handlers.GetAllNotifications).Methods("GET")
	r.HandleFunc("/notifications/{id}", handlers.UpdateNotification).Methods("PUT")
	r.HandleFunc("/notifications/{id}", handlers.DeleteNotification).Methods("DELETE")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
