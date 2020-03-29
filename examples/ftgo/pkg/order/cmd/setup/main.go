package main

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/eiji03aero/mskit/db/postgres"
	"github.com/eiji03aero/mskit/db/postgres/eventstore"
	"github.com/eiji03aero/mskit/db/postgres/sagastore"
	"github.com/eiji03aero/mskit/utils/logger"
)

func main() {
	dir := getDir()
	sqlFilePath := filepath.Join(dir, "./setup.sql")

	dbOption := postgres.DBOption{
		User:     "ftgo",
		Password: "ftgo123",
		Host:     "ftgo-order-postgres",
		Port:     "5432",
		Name:     "ftgo",
	}

	db, err := postgres.GetDB(dbOption)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = eventstore.InitializeDB(db)
	if err != nil {
		panic(err)
	}

	err = sagastore.InitializeDB(db)
	if err != nil {
		panic(err)
	}

	sqlFileContent, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		panic(err)
	}

	statements := strings.Split(string(sqlFileContent), ";")

	for _, s := range statements {
		result, err := db.Exec(s)
		logger.Println(s, result, err)
	}
}

func getDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}
