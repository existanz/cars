package externalapi

import (
	"cars/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func GetCarByRegNum(regNum string) (models.Car, error) {
	car := models.Car{}
	query := fmt.Sprintf("%s%s", os.Getenv("EXTERNAL_API_URL"), regNum)
	slog.Debug("Send request to: ", "URL", query)
	resp, err := http.Get(query)
	if err != nil {
		return car, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil || body == nil {
		return car, err
	}
	slog.Debug("Response: ", "body", string(body))
	if err := json.Unmarshal(body, &car); err != nil {
		return car, err
	}
	if err := resp.Body.Close(); err != nil {
		return car, err
	}
	if resp.StatusCode != http.StatusOK {
		return car, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return car, nil
}
