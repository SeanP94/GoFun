package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

const CONN_STRING string = "postgresql://postgres:dockerpw123@localhost:5432/searchEngineTest?sslmode=disable"

func createTable(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully.")
}

func insertData(db *sql.DB, table csvToSql) {
	// Inserts the data by row test.
	for _, row := range table.rowData {

		// These two are specific for personTable.
		age, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		networth, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(table.insertStatement, row[0], age, row[2], networth)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func selectData(db *sql.DB, table csvToSql) {
	var test string

	query := `
		SELECT name
		FROM people_simple
		LIMIT 1;
	`
	err := db.QueryRow(query).Scan(&test)

	if err != nil {
		log.Fatal("Failed to select person")
	}
	fmt.Println(test)
}

func main() {
	// ^^Above code has been used for learning. Do not delete until progress is a bit more consistent. ^^ //
	// Connect to db.
	db, err := sql.Open("postgres", CONN_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if DB is connected properly.
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	var personTable csvToSql

	personTable.csvToSqlInit("data/test.csv", "people")
	createTableQuery := personTable.createTable("name")

	// Tests (For learning)
	createTable(db, createTableQuery)
	insertData(db, personTable)
	selectData(db, personTable)
}
