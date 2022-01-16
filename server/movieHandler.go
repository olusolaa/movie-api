package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/movieApi/swapi"
	"log"
	"net/http"
	"sort"
)

// @Summary      Get all movies
// @Description  Get all movies in order of their release date from earliest to newest in the cache or from swapi if the cache is empty
// @Produce  json
// @Success 200 {object} models.MovieResponse
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies [get]
func (s *Server) GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movies = s.Cache.Get("movies")
		if movies == nil {
			data, err := swapi.GetAllMovies()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			result := *data
			sort.Slice(result, func(i, j int) bool {
				return result[i].ReleaseDate > result[j].ReleaseDate
			})
			movies = &result
			s.addCommentCountToMovies(movies)
			s.Cache.Set("movies", movies)
			log.Println("Movies added to cache")
		}
		c.JSON(http.StatusOK, movies)
	}
}
