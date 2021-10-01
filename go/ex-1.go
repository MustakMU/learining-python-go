package main

/**

Exercise - 1

1-Create employee data excel file with duplicate record  that contain the fallowing field
Emp id
Emp Name
Emp Salary
Write a python code to display duplicate employee record and count no of duplicate record available into a employee data file
**/

import (
	//"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("employee-det.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	duplicates := map[string]int{}
	totalDuplicates := 0

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := duplicates[record[0]]; ok {
			duplicates[record[0]] = duplicates[record[0]] + 1
			totalDuplicates++
		} else {
			duplicates[record[0]] = 1
		}
	}
	println("Total duplicates ", totalDuplicates)
	println("duplicates ", totalDuplicates)
}
