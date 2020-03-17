package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type DBOption struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetDBUrl(opt DBOption) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		opt.User,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Name,
	)
}

func GetDB(opt DBOption) (*sql.DB, error) {
	return sql.Open("postgres", GetDBUrl(opt))
}

func CreateTable(db *sql.DB, tableName string, columns []string) (err error) {
	_, err = db.Exec(BuildDropTableStatement(tableName))
	if err != nil {
		return
	}

	_, err = db.Exec(BuildCreateTableStatement(tableName, columns))
	if err != nil {
		return
	}

	return
}

func BuildDropTableStatement(tableName string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
}

func BuildCreateTableStatement(tableName string, columns []string) string {
	query := "CREATE TABLE %s (%s)"
	columnsStr := strings.Join(columns, ", ")
	return fmt.Sprintf(query, tableName, columnsStr)
}

func BuildInsertStatement(tableName string, columns []string) string {
	query := "INSERT INTO %s (%s) VALUES (%s)"

	columnsFragment := strings.Join(columns, ", ")

	placeholders := []string{}
	for i, _ := range columns {
		p := fmt.Sprintf("$%d", i+1)
		placeholders = append(placeholders, p)
	}
	placeholderFragment := strings.Join(placeholders, ", ")

	return fmt.Sprintf(
		query,
		tableName,
		columnsFragment,
		placeholderFragment,
	)
}

func BuildSelectStatement(tableName string, columns []string) string {
	query := "SELECT %s FROM %s"

	columnsFragment := strings.Join(columns, ", ")

	return fmt.Sprintf(
		query,
		columnsFragment,
		tableName,
	)
}
