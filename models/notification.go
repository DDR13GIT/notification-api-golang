package models

type Notification struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
	IsRead  bool   `json:"is_read"`
}