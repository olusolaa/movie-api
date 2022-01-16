package models

type CommentRequest struct {
	Content string `json:"content"`
}

type ApiError struct {
	Message string `json:"message"`
}
