package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	Events *[]Event `json:"books"`
}

type EventlistModel struct {
	Endpoint string
}

func (m *EventlistModel) GetAll() (*[]Event, error) {
	resp, err := http.Get(m.Endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var eventsResp EventsResponse
	err = json.Unmarshal(data, &eventsResp)
	if err != nil {
		return nil, err
	}

	return eventsResp.Events, nil
}
