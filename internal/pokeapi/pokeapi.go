package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func GetApiData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	body, err := io.ReadAll((res.Body))
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code %d\nMessage: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}

	return body, nil
}
