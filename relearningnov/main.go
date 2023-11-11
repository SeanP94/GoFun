package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// readCsv reads in a fileName and returns the file in a [][]string format.
func readCsv(fn string) [][]string {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal("You gave me the wrong file idiot.")
	}
	defer f.Close() // Make sure you close the file.
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Idk why but its not working.")
	}
	return records
}

// Inserts a number between 0-4 into a slice.
func randInsert(myArr *[]int) {
	*myArr = append(*myArr, rand.Intn(5))
}

func main() {
	myArr := []int{}
	for i := 0; i < 5; i++ {
		randInsert(&myArr)
	}
	fmt.Println(myArr)
	for i, val := range myArr {
		fmt.Printf(" %-2d: %d\n", i, val)
	}
	for _, row := range readCsv("test.csv") {
		fmt.Println(row)
	}
}
