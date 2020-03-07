package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
	_ "github.com/lib/pq"
)

type DBOption struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type Client struct {
	db            *sql.DB
	eventRegistry *mskit.EventRegistry
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

func GetDB(opt DBOption) (db *sql.DB, err error) {
	return sql.Open("postgres", getDBUrl(opt))
}

func InitializeDB(opt DBOption) (db *sql.DB, err error) {
	db, err = GetDB(opt)
	if err != nil {
		return
	}

	err = CreateTable(db, "mskit_events", []string{
		"id SERIAL PRIMARY KEY",
		"event_type VARCHAR NOT NULL",
		"aggregate_type VARCHAR NOT NULL",
		"aggregate_id VARCHAR NOT NULL",
		"event_data TEXT NOT NULL DEFAULT ''",
	})
	if err != nil {
		return
	}

	return
}

func New(opt DBOption, er *mskit.EventRegistry) (mskit.EventStore, error) {
	db, err := GetDB(opt)
	if err != nil {
		return nil, err
	}

	es := &Client{
		db:            db,
		eventRegistry: er,
	}

	return es, nil
}

func (c *Client) Save(event mskit.Event) error {
	query := BuildInsertStatement(
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

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.Type,
		event.AggregateType,
		event.AggregateId,
		string(eventDataJson),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Load(id string, aggregate mskit.Aggregate) error {
	_, aggregateName := utils.GetTypeName(aggregate)
	query := BuildSelectStatement(
		"mskit_events",
		[]string{
			"event_type",
			"event_data",
		},
	)
	query = query + fmt.Sprintf(" WHERE aggregate_id = $1 AND aggregate_type = $2")

	rows, err := c.db.Query(query, id, aggregateName)
	if err != nil {
		return err
	}

	for rows.Next() {
		var eventName string
		var eventData []byte
		if err := rows.Scan(&eventName, &eventData); err != nil {
			return err
		}

		eventPtr, err := c.eventRegistry.Get(eventName)
		if err != nil {
			return err
		}

		err = json.Unmarshal(eventData, eventPtr)
		if err != nil {
			return err
		}

		event := utils.DereferenceIfPtr(eventPtr)
		err = aggregate.Apply(event)
		if err != nil {
			return err
		}
	}

	return nil
}
