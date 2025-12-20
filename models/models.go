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

type ScanSummary struct {
	Id          string
	CreatedAt   time.Time
	ResultCount int
	FirstTest   time.Time
	LastTest    time.Time
}

type ScanResultQueryOptions struct {
	ScanID   *string
	Host     *string
	FromDate *time.Time
	ToDate   *time.Time
	Grade    *string // dentro del JSON (ej: "grade":"A")
	Limit    *int
	Offset   *int
}
