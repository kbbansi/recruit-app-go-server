package config

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

const recruitAppGo_DB = "recruit_app_go"
const recruitAppGo_User = "root"
const recruitAppGo_Password = "devionJave"

var Database *sql.DB
var message string

func DbConnect() (*sql.DB, error) {
	// connect to database

	// create the database connection string
	dataSourceName:= GetDbCredentials()

	//create database connection
	database, err:= sql.Open("mysql", dataSourceName)

	// handle errors with connection
	if err != nil {
		message = err.Error()
		fmt.Println(message)
		return nil, err
	}

	if err = database.Ping(); err != nil {
		message = err.Error()
		fmt.Println(message)
		return nil, err
	}
	database.SetMaxIdleConns(0)
	return database, nil
}

func GetDbCredentials() string {
	dbUser:= recruitAppGo_User
	dbPassword:= recruitAppGo_Password
	database:= recruitAppGo_DB
	return dbUser + ":" + dbPassword + "@/" + database
}