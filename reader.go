package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CsvLine struct {
	Column1 string
	Column2 string
	Column3 string
	Column4 string
	Column5 string
	Column6 string
}

func main() {
	readAll("result.csv")
	read("result.csv")
}

func readAll(filename string) {
	lines, err := ReadCsv(filename)
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
			Column2: line[1],
			Column3: line[2],
			Column4: line[3],
			Column5: line[4],
			Column6: line[5],
		}
		fmt.Println(data.Column1 + " " + data.Column2 + " " + data.Column3 + " " + data.Column4 + " " + data.Column5 + " " + data.Column6)
	}
	fmt.Println("\n")
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func read(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	//file could not have been opened so close.
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []string{}, err
		}
		fmt.Println(record)
	}
	return nil, nil
}