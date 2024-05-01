package models

type Car struct {
	Id     int    `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year,omitempty"`
	Owner  People `json:"owner"`
}

type CarDTO struct {
	RegNum string    `json:"regNum"`
	Mark   string    `json:"mark"`
	Model  string    `json:"model"`
	Year   int       `json:"year,omitempty"`
	Owner  PeopleDTO `json:"owner"`
}
