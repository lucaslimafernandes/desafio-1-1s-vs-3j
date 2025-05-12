package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Team struct {
	Name     string `json:"name"`
	Leader   bool   `json:"leader"`
	Projects []Project
}

type Log struct {
	Date   time.Time `json:"date"`
	Action string    `json:"action"`
}

type User struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Age     int32     `json:"age"`
	Score   int32     `json:"score"`
	Active  bool      `json:"active"`
	Country string    `json:"country"`
	Team    Team      `json:"team"`
	Logs    []Log     `json:"logs"`
}

func (l *Log) UnmarshalJSON(data []byte) error {
	type Alias Log
	aux := &struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(l),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", aux.Date)
	if err != nil {
		return err
	}

	l.Date = parsedDate
	return nil
}
