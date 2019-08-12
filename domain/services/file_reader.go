package services

import (
	"encoding/csv"
	"io"
	"os"
)

//FileReader service to read lines from a csv file
type FileReader struct{}

//ReadFromFile expect a file path and return a array containing a array of string, and a error
func (f *FileReader) ReadFromFile(path string) (lines [][]string, err error) {
	csvfile, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, record)
	}

	return lines, err
}
