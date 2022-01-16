package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/movieApi/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

// @Summary Add comments
// @Description Add comments to a movie by movie id
// @Accept  json
// @Produce  json
// @Param comment body models.CommentRequest true "Comment"
// @Param movie_id path int true "MovieId"
// @Success 200 {object} models.Comment
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies/{movie_id}/comments [post]
func (s *Server) AddComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid movie id"})
			return
		}
		request := &models.CommentRequest{}
		if errs := s.decode(c, request); errs != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errs})
			return
		}
		comment := &models.Comment{
			MovieId:   movieID,
			Content:   request.Content,
			IP:        request.UserIPAdr,
			CreatedAt: time.Now(),
		}
		data, err := s.DB.AddComment(comment)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !s.increaseCommentCountInRedis(movieID) {
			log.Println("movie id not found in redis")
		}
		c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully!", "data": data})
	}
}

// @Summary Get comments
// @Description Get all comments for a movie by movie id
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Success 200 {object} models.Comment
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies/{movie_id}/comments [get]
// GetComments returns all comments of a movie
func (s *Server) GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := s.DB.GetComments(movieId)

		for i := 0; i < len(*data)/2; i++ {
			(*data)[i], (*data)[len(*data)-1-i] = (*data)[len(*data)-1-i], (*data)[i]
		}
		c.JSON(http.StatusOK, gin.H{"message": "Comments fetched successfully!", "data": data})
	}
}

// increaseCommentCountInRedis increases the comment count of corresponding movie in redis
func (s *Server) increaseCommentCountInRedis(movieID int) bool {
	var movies = s.Cache.Get("movies")
	if movies != nil {
		for idx, movie := range *movies {
			if movie.EpisodeId == movieID {
				temp := models.Movie{
					EpisodeId:    movie.EpisodeId,
					Title:        movie.Title,
					CommentCount: movie.CommentCount + 1,
					OpeningCrawl: movie.OpeningCrawl,
					ReleaseDate:  movie.ReleaseDate,
				}
				(*movies)[idx] = temp
				s.Cache.Set("movies", movies)
				fmt.Println("movie comment count increased", movies)
				return true
			}
		}
	}
	return false
}

// addCommentCountToMovies add the comment count for uncached movies.
func (s *Server) addCommentCountToMovies(movies *[]models.Movie) {
	for idx, movie := range *movies {
		count, _ := s.DB.CountComments(movie.EpisodeId)
		temp := models.Movie{
			EpisodeId:    movie.EpisodeId,
			Title:        movie.Title,
			CommentCount: count,
			OpeningCrawl: movie.OpeningCrawl,
			ReleaseDate:  movie.ReleaseDate,
		}
		(*movies)[idx] = temp
	}
}
