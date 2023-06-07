package model

import (
	"time"
)

type Promotion struct {
	ID             int       `json:"id"`
	PID            string    `json:"pid"`
	Price          float64   `json:"price"`
	ExpirationDate time.Time `json:"expiration_date"`
}
