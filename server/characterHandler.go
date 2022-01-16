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
// @Description Get all characters for a movie by movie id use the sort parameter
//to sort the results by name or height or gender, and the order parameter to order in assending or desending order
// @Produce  json
// @Param movie_id path int true "Movie ID"
// @Param sort_by path string false "Sort by field"
// @Param order path string false "Order"
// @Param filter_by path string false "Filter by field"
// @Success 200 {object} []models.Character
// @Failure 404 {object} models.ApiError
// @Failure 500 {object} models.ApiError
// @Router /api/v1/movies/{movie_id}/characters/{filter}/{sort_by}/{order} [get]
func (s *Server) GetCharacters() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortParam := c.Param("sort_by")
		filterParam := strings.Trim(c.Param("filter"), " ")
		orderParam := c.Param("order")
		movieId, err := strconv.Atoi(c.Param("movie_id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("Getting characters for movie id: ", movieId)
		log.Println("Sort by: ", len(sortParam))
		log.Println("Filter by: ", len(filterParam))
		log.Println("Order: ", len(orderParam))
		links, err := swapi.GetAllCharactersByMovieId(movieId)
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

		//if filterParam != "" {
		//	filteredList := []models.Character{}
		//	for _, character := range characters {
		//		if character.Gender == filterParam {
		//			filteredList = append(filteredList, character)
		//		}
		//	}
		//	characters = filteredList
		//}
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
		heightString, _ := fmt.Printf("%.2fcm | %2fft | %2finches", heightTotal, ft, inches)

		c.JSON(http.StatusOK, gin.H{"message": "user info retrieved successfully", "data": characters,
			"metadata": gin.H{"matching_characters": len(characters), "total_height": heightString}})
	}
}
