package lib

import (
	"encoding/csv"
	"log"
	"os"
)

func CsvScan(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
