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

// Inserts a number between 0-4 into a slice.
// func randInsert(myArr *[]int) {
// 	*myArr = append(*myArr, rand.Intn(5))
// }

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// type TreeNode struct {
// 	Right *TreeNode
// }

// type DepthNode struct {
// 	currDepth int
// 	treeNode  *TreeNode
// }

const CONN_STRING string = "postgresql://postgres:dockerpw123@localhost:5432/searchEngineTest?sslmode=disable"

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
		log.Fatal("Failed to select publisher")
	}
	fmt.Printf("The top game publisher is %v\n", publisher)
}

func main() {
	/*
			var root *TreeNode
		stack := []DepthNode{{1, root}, {3, root}}
		stack = append(stack, DepthNode{2, &TreeNode{root}})

		x := stack[len(stack)-1]

		fmt.Printf("%v, %T, %T\n", x.treeNode, x.currDepth, x.treeNode)
		// Test for opening csv files. It works.

		db, err := sql.Open("postgres", CONN_STRING)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()
		// Check if DB is connected properly.
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}

		// printTopGame(db)
		fmt.Println(math.Max(1, 2))
		printTopGame(db)
		fmt.Println("\n\n\n\n")

	*/

	// ^^Above code has been used for learning. Do not delete until progress is a bit more consistent. ^^ //
	var personTable csvToSql

	personTable.csvToSqlInit("data/test.csv")
	createTableQuery := personTable.createTable("people", "name")
	fmt.Println(createTableQuery)
}
