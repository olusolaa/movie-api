package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/movieApi/models"
	"github.com/olusolaa/movieApi/swapi"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func (s *Server) GetCharacters() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortParam := c.Query("sort_by")
		filterParam := strings.Trim(c.Query("filter"), "\\s+")
		orderParam := c.Query("order")
		movieID, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid movie id"})
			return
		}
		links, err := swapi.GetAllCharactersByMovieId(movieID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var characters []models.Character
		for _, link := range *links {
			info, _ := swapi.GetCharacterInfo(link)
			characters = append(characters, *info)
		}

		if orderParam == "descending" {
			switch sortParam {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Name > characters[j].Name
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Height > characters[j].Height
				})
			case "gender":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Gender > characters[j].Gender
				})
			}
		} else {
			switch sortParam {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Name < characters[j].Name
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Height < characters[j].Height
				})
			case "gender":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Gender < characters[j].Gender
				})
			}
		}

		if filterParam != "" {
			filteredList := []models.Character{}
			for _, character := range characters {
				if character.Gender == filterParam {
					filteredList = append(filteredList, character)
				}
			}
			characters = filteredList
		}
		c.JSON(http.StatusOK, gin.H{"message": "user info retrieved successfully", "data": characters})
	}
}
