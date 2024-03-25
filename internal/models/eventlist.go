package models

import (
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"title"`
	Category    []string  `json:"category"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
}

type EventResponse struct {
	Event *Event `json:"book"`
}

type EventsResponse struct {
	Event *[]Event `json:"books"`
}

type EventlistModel struct {
	Endpoint string
}
