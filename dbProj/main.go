package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"math"
	_ "github.com/lib/pq"
)

const CONN_STRING string = "postgresql://postgres:dockerpw123@localhost:5432/searchEngineTest?sslmode=disable"

type gameEntry struct {
	rank         int
	name         string
	platform     string
	year         int
	genre        string
	publisher    string
	na_sales     float32
	eu_sales     float32
	jp_sales     float32
	other_sales  float32
	global_sales float32
}

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

func printTopGame(db *sql.DB) {

	// var topGame gameEntry
	var publisher string
	query := `
		SELECT "Publisher"
		FROM products
		LIMIT 1;
	`
	err := db.QueryRow(query).Scan(&publisher)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The top game publisher is %v\n", publisher)
}

type TreeNode struct {
	Right *TreeNode
	value int
}

type DepthNode struct {
    currDepth int
    treeNode *TreeNode
}


func main() {
	var root *TreeNode
	stack := []DepthNode{DepthNode{1, root}}
	x,y := stack[len(stack)-1]
	fmt.Printf("%T, %T", x,y )
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

	// printTopGame(db)
	fmt.Println(math.Max(1,2))
	
}
