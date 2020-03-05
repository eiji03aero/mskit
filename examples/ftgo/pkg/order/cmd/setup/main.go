package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/eiji03aero/mskit/eventstore/postgres"
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

	db, err := postgres.InitializeDB(dbOption)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlFileContent, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		panic(err)
	}

	statements := strings.Split(string(sqlFileContent), ";")

	for _, s := range statements {
		result, err := db.Exec(s)
		log.Printf("result=%v err=%v \n", result, err)
	}
}

func getDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}