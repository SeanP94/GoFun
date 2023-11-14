package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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

type csvToSql struct {
	columnTypeMap map[string]string // Used to store the strings of the column types
	rowData       [][]string
	columns       []string
	tableName     string
}

// csvToSqlInit will take in a filename of a csv file, adds values to all of csvToSql's variables.
func (c *csvToSql) csvToSqlInit(sqlFn string, tableName string) {
	csvFile := readCsv(sqlFn)
	c.rowData = csvFile[1:][:]
	c.columns = csvFile[0][:]
	c.columnTypeMap = map[string]string{}
	c.parseRows()
	c.tableName = tableName
}

// parseRows: takes in a [][]string rowData object. Creates the columnTypeMap for the csvToSql class.
// utilizes regex to map either string, int, or float to each column.
func (c *csvToSql) parseRows() {
	// Add items to the map
	for _, column := range c.columns {
		_, ok := c.columnTypeMap[column]
		if ok {
			log.Fatalf("Duplicate column name %v found in columns. Please correctRows", column)
		}
		c.columnTypeMap[column] = "TEXT"
	}

	intMatch := regexp.MustCompile(`^-*[1-9]+[0-9]*$`)
	floatMatch := regexp.MustCompile(`^-{0,1}[1-9]+[0-9]*\.[0-9]+`)

	for _, row := range c.rowData {
		for i, valueStr := range row {
			colKey := c.columns[i]
			value := []byte(valueStr)

			// Match the strings with regex to see if they fit a parameter
			foundInt := intMatch.Match([]byte(value))
			foundFloat := floatMatch.Match(value)

			if foundInt && c.columnTypeMap[colKey] == "TEXT" {
				c.columnTypeMap[colKey] = "INT"
			} else if foundFloat {
				c.columnTypeMap[colKey] = "NUMERIC(5, 2)"
			}

			// TODO: Dates later
			// Date and Datetime? Research postgres first
		}
	}
	fmt.Println(c.columnTypeMap)
}

// createTable will generate a create table  template call for the csv data.
func (c *csvToSql) createTable(primaryKey string) string {
	pk := slices.Contains(c.columns, primaryKey)
	if !pk {
		log.Fatalf("%v cannot be a primary key, as it is not a column.", primaryKey)
	}
	statement := []string{"DROP TABLE IF EXISTS " + c.tableName + ";\n\n CREATE TABLE IF NOT EXISTS " + c.tableName + " ("}
	for _, column := range c.columns {
		line := fmt.Sprintf("\t%v %v", column, c.columnTypeMap[column])
		if primaryKey == column {
			line += " PRIMARY KEY"
		}
		statement = append(statement, line+",")
	}
	// Remove the comma from the last value
	statement[len(statement)-1] = strings.TrimRight(statement[len(statement)-1], ",")
	return strings.Join(statement, "\n") + "\n);"
}

// insertQuery will create a template for
func (c csvToSql) insertQuery() string {
	statement := []string{fmt.Sprintf("INSERT INTO %v (", c.tableName)}
	vals := []string{}
	for i, col := range c.columns {
		// Could of just used strings.Join(c.columns, ", ") But needed the vals as well.
		// Look into a one line solution for vals.
		statement = append(statement, col+",")
		vals = append(vals, "$"+strconv.Itoa(i+1)+",")
	}
	// Remove the comma from the last value
	statement[len(statement)-1] = strings.TrimRight(statement[len(statement)-1], ",")
	vals[len(vals)-1] = strings.TrimRight(vals[len(vals)-1], ",")
	return strings.Join(statement, "") + ")\nVALUES (" + strings.Join(vals, "") + ");"

}
