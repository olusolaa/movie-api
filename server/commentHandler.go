package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/movieApi/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) AddComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID, err := strconv.Atoi(c.Param("movie_id"))

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

		c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully!", "data": data})
	}
}

func (s *Server) GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId, err := strconv.Atoi(c.Param("movie_id"))
		data, err := s.DB.GetComments(movieId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Comments fetched successfully!", "data": data})
	}
}
