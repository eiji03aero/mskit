package sagastore

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/postgres"
	"github.com/eiji03aero/mskit/utils"
)

const (
	TableName = "mskit_saga_instances"
)

type Client struct {
	db *sql.DB
}

func InitializeDB(opt postgres.DBOption) (db *sql.DB, err error) {
	db, err = postgres.GetDB(opt)
	if err != nil {
		return
	}

	err = postgres.CreateTable(db, TableName, []string{
		"id VARCHAR PRIMARY KEY",
		"saga_state smallint NOT NULL",
		"data TEXT NOT NULL DEFAULT ''",
	})
	if err != nil {
		return
	}

	return
}

func New(opt postgres.DBOption) (mskit.SagaInstanceRepository, error) {
	db, err := postgres.GetDB(opt)
	if err != nil {
		return nil, err
	}

	repository := &Client{
		db: db,
	}

	return repository, nil
}

func (c *Client) Save(si SagaInstance) error {
	query := postgres.BuildInsertStatement(
		TableName,
		[]string{
			"id",
			"saga_state",
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
		si.SagaState,
		string(siDataJson),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Load(id string, sagaInstance *mskit.SagaInstance) error {
	var dataStr string
	query := postgres.BuildSelectStatement(
		TableName,
		[]string{
			"saga_state",
			"data",
		},
	)
	query = query + fmt.Sprintf(" WHERE id = $1")

	err := c.db.QueryRow(query, id).
		Scan(&sagaInstance.SagaState, &dataStr)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(dataStr), &sagaInstance.Data)
	if err != nil {
		return err
	}

	sagaInstance.Id = id

	return nil
}
