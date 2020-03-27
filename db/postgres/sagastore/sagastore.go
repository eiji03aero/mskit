package sagastore

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/postgres"
)

const (
	TableName = "mskit_saga_instances"
)

type Client struct {
	db *sql.DB
}

func InitializeDB(db *sql.DB) (err error) {
	err = postgres.CreateTable(db, TableName, []string{
		"id VARCHAR PRIMARY KEY",
		"step_index smallint NOT NULL",
		"state smallint NOT NULL",
		"data TEXT NOT NULL DEFAULT ''",
	})
	if err != nil {
		return
	}

	return
}

func New(opt postgres.DBOption) (mskit.SagaStore, error) {
	db, err := postgres.GetDB(opt)
	if err != nil {
		return nil, err
	}

	repository := &Client{
		db: db,
	}

	return repository, nil
}

func (c *Client) Save(si *mskit.SagaInstance) error {
	query := postgres.BuildInsertStatement(
		TableName,
		[]string{
			"id",
			"step_index",
			"state",
			"data",
		},
	)

	siDataJson, err := json.Marshal(si.Data)
	if err != nil {
		return err
	}

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		si.Id,
		si.StepIndex,
		si.State,
		string(siDataJson),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Update(si *mskit.SagaInstance) error {
	query := postgres.BuildUpdateStatement(
		TableName,
		[]string{
			"step_index",
			"state",
			"data",
		},
	)
	query += " WHERE id = $4"

	siDataJson, err := json.Marshal(si.Data)
	if err != nil {
		return err
	}

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		si.StepIndex,
		si.State,
		string(siDataJson),
		si.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Load(id string, si *mskit.SagaInstance) error {
	query := postgres.BuildSelectStatement(
		TableName,
		[]string{
			"state",
			"data",
		},
	)
	query = query + fmt.Sprintf(" WHERE id = $1")

	err := c.db.QueryRow(query, id).
		Scan(&si.State, &si.Data)
	if err != nil {
		return err
	}

	si.Id = id

	return nil
}
