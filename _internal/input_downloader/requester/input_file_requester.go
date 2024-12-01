package requester

import (
	"fmt"
	"io"
	"net/http"
)

type Requester struct {
	SessionToken *string
}

// Get retrieves the Advent of Code puzzle input for the given day and year.
func (r *Requester) Get(day int, year int) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	fmt.Printf("Retrieving puzzle input for day %d, year %d...\n", day, year)
	fmt.Println("URL:", url)

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return make([]byte, 0), err
	}

	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: *r.SessionToken,
	})

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return make([]byte, 0), err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
			return
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return make([]byte, 0), fmt.Errorf("received non-200 status code: %d | details: %v", response.StatusCode, response.Status)
	}

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return make([]byte, 0), err
	}

	// Trim last line
	bodyBytes = bodyBytes[:len(bodyBytes)-len("\n")]

	return bodyBytes, nil
}
