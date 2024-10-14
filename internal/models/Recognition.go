package models

import "time"

type Recognition struct {
	Id         int       `json:"id"`
	Amount     int       `json:"amount"`
	Recipient  string    `json:"recipient"`
	Issuer     string    `json:"issuer"`
	Message    string    `json:"message"`
	DateIssued time.Time `json:"date_issued"`
}
