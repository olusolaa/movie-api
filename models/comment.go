package models

import (
	"time"
)

type Comment struct {
	tableName struct{}  `pg:"movie_comments"`
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	MovieId   int       `json:"movie_id" pg:"movie_id"`
	IP        string    `json:"ip" pg:"ip"`
	Content   string    `json:"content" pg:"content" valid:"required,stringlength(2|500)~content: must be of appropriate length"`
	CreatedAt time.Time `json:"created_at" pg:"created_at"`
}
