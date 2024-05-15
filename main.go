package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var numbers = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	// Two-Dimensional array to store each row of the table as array
	var rows [][]string
	// Raw string table inserted by STDIN
	var table string
	// Returning Result
	var result int

	table = GetTableContent()
	rows = ConvertTableInto2DArray(table)

	// Iterate through each row find symbols and calculate the result
	for nthRow, row := range rows {
		for j := 0; j < len(row); j++ {
			if IsNumber(row[j]) {
				// List to store individual digits of a number
				var number []string

				// Iterate through the row until we find the end of the number
				for k := 0; j+k < len(row); k++ {
					if !IsNumber(row[j+k]) {
						break
					}
					number = append(number, row[j+k])
				}

				// Foreach number essentially calculate its neighbor,
				// if atleast one is found that means that number is
				// valid and we can count it and move on
				for n := j; n < len(number) + j; n++ {
					// Check for valid neighbors
					if hasValidNeighbor(rows, row, nthRow, n){
						result += ConvertNumber(number)
						break
					}
				}
				// This means that we skip over the number in the row in the next loop
				j += len(number) - 1
			}
		}
	}

	fmt.Printf("TOTAL: %v", result)
}

func hasValidNeighbor(rows [][]string, row []string, nthRow int, n int) bool {
	return 	(n != 0 && (row[n-1] != "." && !IsNumber(row[n-1]))) || // Left Neighbor
			(n != len(row)-1 && (row[n+1] != "." && !IsNumber(row[n+1]))) || // Right Neighbor
			(nthRow != 0 && (rows[nthRow-1][n] != "." && !IsNumber(rows[nthRow-1][n]))) || // Top Neighbor
			(nthRow != len(rows)-1 && (rows[nthRow+1][n] != "." && !IsNumber(rows[nthRow+1][n]))) || // Bottom Neighbor
			(nthRow != 0 && n != 0 && (rows[nthRow-1][n-1] != "." && !IsNumber(rows[nthRow-1][n-1]))) || // Top Left Neighbor
			(nthRow != 0 && n != len(row)-1 && (rows[nthRow-1][n+1] != "." && !IsNumber(rows[nthRow-1][n+1]))) || // Top Right Neighbor
			(nthRow != len(rows)-1 && n != 0 && (rows[nthRow+1][n-1] != "." && !IsNumber(rows[nthRow+1][n-1]))) || // Bottom-Left Neighbor
			(nthRow != len(rows)-1 && n != len(row)-1 && (rows[nthRow+1][n+1] != "." && !IsNumber(rows[nthRow+1][n+1]))) // Bottom-Right Neighbor
}

func GetTableContent() string {
	// Get the file from the STDIN
	if len(os.Args) != 2 {
		log.Fatal("Usage: program.exe <filename>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ConvertTableInto2DArray(table string) [][]string {
	var totalRows int
	var rows [][]string

	// Calculate the total of rows
	for _, char := range table {
		if char == '\n' {
			totalRows++
		}
	}

	// Calculate the width of each row
	width := len(table) / (totalRows + 1)

	// Create the two-dimensional array
	for i := 0; i <= totalRows; i++ {
		var array []string
		for j := 0; j < width; j++ {
			index := i*(width+1) + j
			if index < len(table) {
				array = append(array, string(table[index]))
			}
		}
		rows = append(rows, array)
	}

	return rows
}

func ConvertNumber(numberList []string) int {
	numberString := ""
	for _, char := range numberList {
		numberString += char
	}

	num, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatal(err)
	}

	return num
}

func IsNumber(val string) bool {
	for _, char := range numbers {
		if char == val {
			return true
		}
	}

	return false
}
