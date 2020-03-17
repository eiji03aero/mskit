package restaurant

import (
	"database/sql"
	"encoding/json"
	"fmt"

	restaurantdmn "order/domain/restaurant"

	"github.com/eiji03aero/mskit/db/postgres"
)

const (
	table = "json_data"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Save(restaurant restaurantdmn.Restaurant) (err error) {
	query := postgres.BuildInsertStatement(
		table,
		[]string{
			"id",
			"data",
		},
	)

	restaurantJson, err := json.Marshal(restaurant)
	if err != nil {
		return
	}

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		restaurant.Id,
		string(restaurantJson),
	)
	if err != nil {
		return
	}

	return nil
}

func (r *Repository) FindById(id string) (restaurant *restaurantdmn.Restaurant, err error) {
	query := postgres.BuildSelectStatement(
		table,
		[]string{
			"data",
		},
	)
	query = query + fmt.Sprintf(" WHERE id = $1")

	var dataStr string
	err = r.db.QueryRow(query, id).Scan(&dataStr)
	if err != nil {
		return
	}

	restaurant = &restaurantdmn.Restaurant{}
	err = json.Unmarshal([]byte(dataStr), restaurant)
	return
}
