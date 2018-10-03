package main

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kovetskiy/ko"
	"gopkg.in/yaml.v2"
)

func makeDBConfig() (*mysql.Config, error) {
	var dbConfig mysql.Config
	var err = ko.Load("database.yaml", &dbConfig, yaml.Unmarshal)
	if err != nil {
		return nil, errors.New("[ERROR]: Unable to parse configuration file!")
	}
	dbConfig.AllowNativePasswords = true
	return &dbConfig, nil
}

//ConnectToDB connects to mysql Database using configuration
//specified in database.yaml
func ConnectToDB() (*sqlx.DB, error) {
	var db *sqlx.DB

	conf, err := makeDBConfig()

	if err != nil {
		return nil, err
	}

	db, err = sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, errors.New("[ERROR]: Unable to connect to Database!")
	}
	return db, nil
}
