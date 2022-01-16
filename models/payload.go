package models

type CommentRequest struct {
	UserIPAdr string `json:"user_ip"`
	Content   string `json:"content"`
}

type CommentResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []Comment `json:"response"`
}

type MovieResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    []Movie `json:"response"`
}

type CharacterResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []Character `json:"response"`
}

type ApiError struct {
	Status string `json:"status"`
	Data   string `json:"message"`
}
