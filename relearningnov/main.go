package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"math/rand"
	"os"

	_ "github.com/lib/pq"
)

const CONN_STRING string = "postgresql://postgres:dockerpw123@localhost:5432/searchEngineTest?sslmode=disable"

// readCsv reads in a fileName and returns the file in a [][]string format.
func readCsv(fn string) [][]string {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal("You gave me the wrong file idiot: \n\t- ", err)
	}
	defer f.Close() // Make sure you close the file.
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Idk why but its not working. \n\t- ", err)
	}
	return records
}

// Inserts a number between 0-4 into a slice.
func randInsert(myArr *[]int) {
	*myArr = append(*myArr, rand.Intn(5))
}

func main() {
	// Test for opening csv files. It works.
	readCsv("data/test.csv")

	db, err := sql.Open("postgres", CONN_STRING)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// Check if DB is connected properly.
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}
