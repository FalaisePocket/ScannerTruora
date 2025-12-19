package models

import "time"

type Scan struct {
	Id        string
	Host      []string
	FirstTest time.Time
	LastTest  time.Time
	CreatedAt time.Time
}

type ScanResult struct {
	Id        int
	ScanID    string
	URL       string
	Result    string
	CreatedAt time.Time
}
