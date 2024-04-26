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
	var rowsArray [][]string
	// Returning Result
	var result int
	// Total rows in the table
	var rowsTotal int
	// Raw string table inserted by STDIN
	var rawTable string

	// Get the file from the STDIN
	if len(os.Args) != 2 {
		fmt.Println("Usage: program.exe <filename>")
		return
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

	rawTable = string(content)

	// Calculate the total of rows
	for _, char := range rawTable {
		if char == '\n' {
			rowsTotal++
		}
	}

	// Calculate the width of each row
	width := len(rawTable) / (rowsTotal + 1)

	// Create the two-dimensional array
	for i := 0; i <= rowsTotal; i++ {
		var array []string
		for j := 0; j < width; j++ {
			index := i*(width+1) + j
			if index < len(rawTable) {
				array = append(array, string(rawTable[index]))
			}
		}
		rowsArray = append(rowsArray, array)
	}

	// Iterate through each row find symbols and calculate the result
	for nthRow, row := range rowsArray {

		// List to store individual digits of a number
		var number []string
		for j := 0; j < len(row); j++ {

			if row[j] == "." {
				continue
			}

			if isNumber(row[j]) {
				// Iterate through the row until we find the end of the number
				for k := 0; j+k < len(row); k++ {
					if !isNumber(row[j+k]) {
						break
					}
					number = append(number, row[j+k])
				}

				// Foreach number essentially calculate its neighbor,
				// if atleast one is found that means that number is
				// valid and we can count it and move on
				for n := 0; n < len(number); n++ {
					// We do this to essentially imitate the actual index in the row
					n += j

					// Check all neighbors
					if (n != 0 && (row[n-1] != "." && !isNumber(row[n-1]))) || // Left Neighbor
						(n != len(row)-1 && (row[n+1] != "." && !isNumber(row[n+1]))) || // Right Neighbor
						(nthRow != 0 && (rowsArray[nthRow-1][n] != "." && !isNumber(rowsArray[nthRow-1][n]))) || // Top Neighbor
						(nthRow != len(rowsArray)-1 && (rowsArray[nthRow+1][n] != "." && !isNumber(rowsArray[nthRow+1][n]))) || // Bottom Neighbor
						(nthRow != 0 && n != 0 && (rowsArray[nthRow-1][n-1] != "." && !isNumber(rowsArray[nthRow-1][n-1]))) || // Top Left Neighbor
						(nthRow != 0 && n != len(row)-1 && (rowsArray[nthRow-1][n+1] != "." && !isNumber(rowsArray[nthRow-1][n+1]))) || // Top Right Neighbor
						(nthRow != len(rowsArray)-1 && n != 0 && (rowsArray[nthRow+1][n-1] != "." && !isNumber(rowsArray[nthRow+1][n-1]))) || // Bottom-Left Neighbor
						(nthRow != len(rowsArray)-1 && n != len(row)-1 && (rowsArray[nthRow+1][n+1] != "." && !isNumber(rowsArray[nthRow+1][n+1]))) { // Bottom-Right Neighbor
						result += convertNumber(number)
						break
					}

					// Do this so we don't exceed the n < len(numbers)
					n -= j
				}

				// This means that we skip over the number in the row in the next loop
				j += len(number) - 1
				number = []string{}
			}
		}
	}

	fmt.Printf("TOTAL: %v", result)
}

func convertNumber(numberList []string) int {
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

func isNumber(val string) bool {
	for _, char := range numbers {
		if char == val {
			return true
		}
	}

	return false
}
