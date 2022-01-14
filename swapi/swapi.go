package swapi

import (
	"encoding/json"
	"fmt"
	"github.com/olusolaa/movieApi/models"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const swapiBaseUrl = "https://swapi.dev/api"

//const swapiBaseUrl = "https://swapi-deno.azurewebsites.net/api/"

// GetAllMovies returns a list of movies from the swapi api
func GetAllMovies() (*[]models.Movie, error) {
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

	movies := &models.Movies{}
	//movies := &[]models.Movie{}
	if err := json.Unmarshal(body, movies); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body from swapi api")
	}

	return &movies.Results, nil
}

func GetAllCharactersByMovieId(movieId int) (*[]string, error) {
	url := fmt.Sprintf("%s/films/%d/", swapiBaseUrl, movieId)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get character from swapi api"))
		return nil, errors.Wrap(err, "failed to get character from swapi api")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body from swapi api")
	}

	var data models.CharacterLinks

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(errors.Wrap(err, "failed to unmarshal response body from swapi api"))
		return nil, errors.Wrap(err, "failed to unmarshal response body from swapi api")
	}
	return &data.Characters, nil
}

func GetCharacterInfo(link string) (*models.Character, error) {
	//url := fmt.Sprintf("%s/people/%s/", swapiBaseUrl, link)
	resp, err := http.Get(link)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get character from swapi api"))
		return nil, errors.Wrap(err, "failed to get character from swapi api")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to read response body from swapi api"))
		return nil, errors.Wrap(err, "failed to read response body from swapi api")
	}

	var data models.Character

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(errors.Wrap(err, "failed to unmarshal response body from swapi api"))
		return nil, errors.Wrap(err, "failed to unmarshal response body from swapi api")
	}
	return &data, nil
}
