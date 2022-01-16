package models

type CommentRequest struct {
	UserIPAdr string `json:"user_ip"`
	Content   string `json:"content"`
}

type ApiError struct {
	Message string `json:"message"`
}
