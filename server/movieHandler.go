package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/movieApi/models"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// @Summary      Get all movies
// @Description  Get all movies in order of their release date from earliest to newest in the cache or from swapi if the cache is empty
// @Produce  json
// @Success 200 {object} models.Movie
// @Failure 400 {object} models.Response
// @Router /api/v1/movies [get]
func (s *Server) GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movies = s.Cache.Get("movies")
		if movies == nil {
			data, err := getMoviesFromSwapiApi()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			//result := data.Results
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

//const swapiBaseUrl = "https://swapi.dev/api"
const swapiBaseUrl = "https://swapi-deno.azurewebsites.net/api/"

// GetMovies returns a list of movies from the swapi api
func getMoviesFromSwapiApi() (*[]models.Movie, error) {
	url := fmt.Sprintf("%s/films/", swapiBaseUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get movies from swapi api")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body from swapi api")
	}

	//movies := &models.Movies{}
	movies := &[]models.Movie{}
	if err := json.Unmarshal(body, movies); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body from swapi api")
	}
	return movies, nil
}
