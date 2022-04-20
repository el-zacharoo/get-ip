package model

import "time"

type Geolocation struct {
	ID          string    `json:"id"`
	Country     string    `json:"country"`
	CountryCode string    `json:"countrycode"`
	Date        time.Time `json:"date"`
	IPAddress   string    `json:"ipAddress"`
	Platform    string    `json:"platform"`
	Page        string    `json:"page"`
}

type Page struct {
	Data    []Geolocation `json:"data"`
	Matches int64         `json:"matches"`
}
