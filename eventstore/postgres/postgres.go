package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/eiji03aero/mskit"
	_ "github.com/lib/pq"
)

// TBD
// Load()

type DBOption struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type EventMap map[string]interface{}

type Client struct {
	DB       *sql.DB
	eventMap EventMap
}

func getDBUrl(opt DBOption) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		opt.User,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Name,
	)
}

// func New(opt DBOption, eventMap EventMap) (Repository, error) {
func New(opt DBOption) (mskit.Repository, error) {
	db, err := sql.Open("postgres", getDBUrl(opt))
	if err != nil {
		return nil, err
	}

	repo := &Client{
		DB: db,
		// eventMap: eventMap,
	}

	return repo, nil
}

func (c *Client) Save(event *mskit.Event) error {
	query := buildInsertStatement(
		"mskit_events",
		[]string{
			"event_type",
			"aggregate_type",
			"aggregate_id",
			"event_data",
		},
	)

	eventDataJson, err := json.Marshal(event.Data)
	if err != nil {
		return err
	}

	stmt, err := c.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.Type,
		event.AggregateType,
		event.AggregateID,
		string(eventDataJson),
	)
	if err != nil {
		return err
	}

	return nil
}
