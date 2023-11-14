package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// type gameEntry struct {
// 	rank         int
// 	name         string
// 	platform     string
// 	year         int
// 	genre        string
// 	publisher    string
// 	na_sales     float32
// 	eu_sales     float32
// 	jp_sales     float32
// 	other_sales  float32
// 	global_sales float32
// }

const CONN_STRING string = "postgresql://postgres:dockerpw123@localhost:5432/searchEngineTest?sslmode=disable"

func createTable(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully.")
}

func insertData(db *sql.DB, rows [][]string, query string) {
	for _, row := range rows {
		_, err := db.Exec(query, row[0], row[1])

		if err != nil {
			log.Fatal(err)
		}
	}
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

	personTable.csvToSqlInit("data/simpleTest.csv", "people_simple")
	createTableQuery := personTable.createTable("name")
	fmt.Println(createTableQuery)
	createTable(db, createTableQuery)

	//TODO:
	insertData(db, personTable.rowData, personTable.insertQuery())
	var test string
	query := `
		SELECT name
		FROM people_simple
		LIMIT 1;
	`
	err = db.QueryRow(query).Scan(&test)

	if err != nil {
		log.Fatal("Failed to select person")
	}
	fmt.Println(test)
}
