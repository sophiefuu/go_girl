package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
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

type EventModel struct {
	DB *sql.DB
}

func (b EventModel) Insert(event *Event) error {
	query := `
		INSERT INTO events (title, category, description, location, date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at`

	args := []interface{}{event.Title, pq.Array(event.Category), event.Description, event.Location, event.Date}
	// return the auto generated system values to Go object
	return b.DB.QueryRow(query, args...).Scan(&event.ID, &event.CreatedAt)
}

func (b EventModel) Get(id int64) (*Event, error) {
	if id < 1 {
		return nil, errors.New("record not found")
	}

	query := `
		SELECT id, created_at, title, category, description, location, date
		FROM events
		WHERE id = $1`

	var event Event

	err := b.DB.QueryRow(query, id).Scan(
		&event.ID,
		&event.CreatedAt,
		&event.Title,
		pq.Array(&event.Category),
		&event.Description,
		&event.Location,
		&event.Date,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}

	return &event, nil
}

func (b EventModel) Update(event *Event) error {
	query := `
		UPDATE events title, category, description, location, date
		SET title = $1, category = $2, description = $3, location = $4, date = $5
		WHERE id = $6 AND version = $7
		RETURNING id`

	args := []interface{}{&event.Title,
		pq.Array(&event.Category),
		&event.Description,
		&event.Location,
		&event.Date}
	return b.DB.QueryRow(query, args...).Scan(&event.ID)
}

func (b EventModel) Delete(id int64) error {
	if id < 1 {
		return errors.New("record not found")
	}

	query := `
		DELETE FROM events
		WHERE id = $1`

	results, err := b.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (b EventModel) GetAll() ([]*Event, error) {
	query := `
	  SELECT * 
	  FROM events
	  ORDER BY id`

	rows, err := b.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.ID,
			&event.CreatedAt,
			&event.Title,
			pq.Array(&event.Category),
			&event.Description,
			&event.Location,
			&event.Date,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
