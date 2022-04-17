package model

import "time"

type IPAddress struct {
	IP string `json:"ip"`
}

type Geolocation struct {
	Country     string    `json:"country"`
	CountryCode string    `json:"countrycode"`
	Date        time.Time `json:"date"`
	IPAddress   string    `json:"ipAddress"`
	Platform    string    `json:"platform"`
	Page        string    `json:"page"`
}
