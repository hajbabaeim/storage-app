package model

import "time"

type Promotion struct {
	ID             string
	Price          float64
	ExpirationDate time.Time
}
