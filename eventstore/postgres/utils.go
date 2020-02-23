package postgres

import (
	"fmt"
	"strings"
)

func buildInsertStatement(tableName string, columns []string) string {
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

func buildSelectStatement(tableName string, columns []string) string {
	query := "SELECT %s FROM %s"

	columnsFragment := strings.Join(columns, ", ")

	return fmt.Sprintf(
		query,
		columnsFragment,
		tableName,
	)
}
