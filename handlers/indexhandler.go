package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mAmineChniti/FlixLib/components"
	"github.com/mAmineChniti/FlixLib/models"
	"github.com/mAmineChniti/FlixLib/pages"
	"github.com/mAmineChniti/FlixLib/utils"
)

const (
	apiHost      = "moviesdatabase.p.rapidapi.com"
	itemsPerPage = 10
)

func IndexHandler(c echo.Context) error {
	movieData, nextPage, err := fetchMoviesFromAPI(1)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch movies")
	}

	cardRows := components.CardsRow(movieData, nextPage)
	return utils.Render(c, pages.Index("Home", cardRows))
}

func LoadMoreHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	movieData, nextPage, err := fetchMoviesFromAPI(page)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch movies")
	}

	cardRows := components.CardsRow(movieData, nextPage)
	return utils.Render(c, cardRows)
}

func fetchMoviesFromAPI(page int) ([]models.MovieData, int, error) {
	url := "https://" + apiHost + "/titles?page=" + strconv.Itoa(page)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	var apiKey = os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("error fetching from API: %v, Status code: %d", err, res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading response body: %v", err)
	}

	var apiResponse models.APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, 0, fmt.Errorf("error unmarshalling response: %v", err)
	}

	movieData := make([]models.MovieData, len(apiResponse.Results))
	for i, movie := range apiResponse.Results {
		movieData[i] = models.MovieData{
			Title:    movie.TitleText.Text,
			ImageURL: movie.PrimaryImage.URL,
		}
	}

	var nextPage int
	if len(apiResponse.Results) < itemsPerPage {
		nextPage = 0
	} else {
		nextPage = page + 1
	}

	return movieData, nextPage, nil
}
