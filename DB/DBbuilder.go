package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

//CreateDBconn  ...maybe unnecesary...
func CreateDBconn(connStr *string, DBname string, conn *sql.DB) error {
	var err error
	conn, err = sql.Open(DBname, *connStr)
	return err
}

//BuildDB - building database
func BuildDB() bool {
	connStr := "user=gotest password=gotest dbname=eksh"
	connStrBuild := "user=gotest password=gotest dbname=postgres"
	var connPostgres *sql.DB
	var connEksh *sql.DB

	if CreateDBconn(&connStr, "eksh", connEksh) != nil {
		log.Fatal("Could not connect to database eksh")
		log.Print("Trying to connect to database postgres")

		if CreateDBconn(&connStrBuild, "postgres", connPostgres) != nil {
			log.Fatal("Could not connect to database postgres")
			return false
		}
		_, err69 := connPostgres.Query("CREATE DATABASE eksh")

		if err69 != nil {
			log.Fatal("Could not create database eksh")
			return false
		}
	}

	//try to run migrations....

	return true
}
