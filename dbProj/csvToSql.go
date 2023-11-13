package main

import (
	// "regexp"
	"encoding/csv"
	"log"
	"os"
)

/*
PROJECT TODO:
	Build a csv -> SQL parser.
	--- --- --- --- --- --- ---
	Build a struct that is meant for you to parse csv files into Postgresql.

	Notes:

		readInCsv() # This should happen at the creation of the table.....
		# Reads in a csv file and gets all data types mapped.
		- I believe csv files read out as string or at least that's the format Id like to follow.
		  - Parse lines utilize regex (Look into golang regex funs)
		  - $##.##, ##.##.... float32
		  - ##... ##,###.... $###.... int
		  - ##/##/## Date ? (Look into this so you can val == a date type, or pass parameter that column is datetype )
		  - Else string.

		generateTable()
		# Creates a Create Table string with parameters saved.
		- Pass in table name, unique id, array of nonNullables.

		printStruct()
		# prints a string in the terminal to copy/paste the module struct.
		- Not 100% if this is needed, but assuming if the table is created we might need to have a struct module for
		  database migration/safe querying.

		loadTable()
		# Should take a db object, loads all the saved items in the table that was generated in readInCsv()

		Notes:
			Maybe have some simple query globals ready to test??
*/

// Helper Functions //
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

// type csvToSql struct {
// 	columnTypeMap  map[string]string // Used to store the strings of the column types
// }
