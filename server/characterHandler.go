package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/movieApi/models"
	"github.com/olusolaa/movieApi/swapi"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// @Summary Get characters
// @Description Get all characters for a movie by movie id use the sort parameter to sort the results by name or height or gender, and the order parameter to order in assending or desending order eg /api/v1/movies/{movie_id}/characters?sort_by=height&filter_by=male&order=descending
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Param sort_by query string false "Sort by height or name or gender"
// @Param order query string false "ascending or descending order"
// @Param filter_by query string false "Filter by male or female or n/a or hermaphrodite"
// @Success 200 {object} []models.Character
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies/{movie_id}/characters [get]
func (s *Server) GetCharacters() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortParam := c.Query("sort_by")
		filterParam := strings.TrimSpace(c.Query("filter_by"))
		orderParam := c.Query("order")
		movieId, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		characters := s.Cache.GetCharacters("movie_" + strconv.Itoa(movieId) + "_characters")
		if characters == nil {
			links, err := swapi.GetAllCharactersByMovieId(movieId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			for _, link := range *links {
				info, _ := swapi.GetCharacterInfo(link)
				characters = append(characters, *info)
			}
			s.Cache.SetCharacters("movie_"+strconv.Itoa(movieId)+"_characters", characters)
		}

		if orderParam == "descending" {
			switch sortParam {
			case "name":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Name > characters[j].Name
				})
			case "height":
				sort.Slice(characters, func(i, j int) bool {
					heightI, _ := strconv.ParseFloat(characters[i].Height, 64)
					heightJ, _ := strconv.ParseFloat(characters[j].Height, 64)
					return heightI > heightJ
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
					heightI, _ := strconv.ParseFloat(characters[i].Height, 64)
					heightJ, _ := strconv.ParseFloat(characters[j].Height, 64)
					return heightI < heightJ
				})
			case "gender":
				sort.Slice(characters, func(i, j int) bool {
					return characters[i].Gender < characters[j].Gender
				})
			}
		}

		if filterParam == "male" || filterParam == "female" || filterParam == "n/a" || filterParam == "hermaphrodite" {
			filteredList := []models.Character{}
			for _, character := range characters {
				if character.Gender == filterParam {
					filteredList = append(filteredList, character)
				}
			}
			characters = filteredList
		}

		heightTotal := float64(0)
		for _, character := range characters {
			height, err := strconv.ParseFloat(character.Height, 64)
			if err != nil {
				log.Println(err)
				continue
			}
			heightTotal += height
		}
		ft := heightTotal / 34
		inches := heightTotal / 24.53
		heightString := fmt.Sprintf("%.2fcm || %.2fft || %.2finches", heightTotal, ft, inches)

		c.JSON(http.StatusOK, gin.H{"message": "user info retrieved successfully", "data": characters,
			"metadata": gin.H{"matching_characters": len(characters), "total_height": heightString}})
	}
}
