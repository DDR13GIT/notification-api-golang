package handlers

import (
	"encoding/json"
	"net/http"
	"notification-api-golang/database"
	"notification-api-golang/models"
	"strconv"
	"github.com/gorilla/mux"
)

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	var notification models.Notification
	json.NewDecoder(r.Body).Decode(&notification)

	result, err := database.DB.Exec("INSERT INTO notifications (message, user_id, is_read) VALUES (?, ?, ?)",
		notification.Message, notification.UserID, notification.IsRead)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	notification.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notification)
}

func GetNotification(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var notification models.Notification
	err := database.DB.QueryRow("SELECT id, message, user_id, is_read FROM notifications WHERE id = ?", id).
		Scan(&notification.ID, &notification.Message, &notification.UserID, &notification.IsRead)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notification)
}

func UpdateNotification(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var notification models.Notification
	json.NewDecoder(r.Body).Decode(&notification)
	notification.ID = id

	_, err := database.DB.Exec("UPDATE notifications SET message = ?, user_id = ?, is_read = ? WHERE id = ?",
		notification.Message, notification.UserID, notification.IsRead, notification.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notification)
}

func DeleteNotification(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := database.DB.Exec("DELETE FROM notifications WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
