package cache

import "github.com/olusolaa/movieApi/models"

type Cache interface {
	Set(key string, value *[]models.Movie)
	Get(key string) *[]models.Movie
	SetCharacters(key string, value []models.Character)
	GetCharacters(key string) []models.Character
}
