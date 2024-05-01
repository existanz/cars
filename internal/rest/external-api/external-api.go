package externalapi

import (
	"cars/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetCarByRegNum(regNum string) (models.Car, error) {
	car := models.Car{}
	query := fmt.Sprintf("%s%s", os.Getenv("EXTERNAL_API_URL"), regNum)
	fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return car, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil || body == nil {
		return car, err
	}
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &car); err != nil {
		return car, err
	}
	if err := resp.Body.Close(); err != nil {
		return car, err
	}
	if resp.StatusCode != http.StatusOK {
		return car, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	fmt.Println("Last car:", car, "Error:", err)
	return car, nil
}
