package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	// Function to read a CSV file and return the records
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}

	return records
}

func writeCsvFile(filePath string, records [][]string) {
	// Function to write records to a new CSV file
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	// Write all records to the new CSV file
	w.WriteAll(records)
}

func main() {
	// Read records from CSV file
	records := readCsvFile(os.Args[1])

	// Isolate the header
	header := records[:1]
	// Remove record list header
	records = records[1:]

	// Sort records first by index 1 (second element)
	sort.SliceStable(records, func(i, j int) bool {
		return records[i][1] < records[j][1]
	})

	// Then sort the records by index 0 (first element), ignoring case
	sort.SliceStable(records, func(i, j int) bool {
		return strings.ToLower(records[i][0]) < strings.ToLower(records[j][0])
	})

	// Create a new recordset, including the header
	newfile := append(header, records...)

	// Write the sorted records to a new CSV file
	writeCsvFile(os.Args[2], newfile)
}
