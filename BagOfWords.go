package main

import (
	"fmt"
	"os"
	//	"io"
	"encoding/csv"
	"github.com/kennygrant/sanitize"
	"regexp"
)

func main() {
	tsvFile, err := os.Open("../data/labeledTrainData.tsv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tsvFile.Close()

	tsvReader := csv.NewReader(tsvFile)
	tsvReader.Comma = '\t'
	tsvReader.LazyQuotes = true

	// Getting the header first, and building the set of map
	tsvLine, err := tsvReader.Read()
	if err != nil {
		fmt.Println(err)
	}
	header := make([]string, len(tsvLine))
	train := make(map[string][]string, len(tsvLine))
	for i := 0; i < len(tsvLine); i++ {
		header[i] = tsvLine[i]
		train[header[i]] = make([]string, 25000)
	}

	// Getting the a sample of 25000 records
	for i := 0; i < 25000; i++ {
		tsvLine, err = tsvReader.Read()
		if err != nil {
			fmt.Printf("Error reading line... %v\n", err)
		}

		for j := 0; j < len(tsvLine); j++ {
			// First we remove all HTML tags
			sanitizedLine := sanitize.HTML(tsvLine[j])

			// Then we remove all characters but letters...
			re := regexp.MustCompile("[^a-zA-Z]")
			lettersLine := re.ReplaceAllString(sanitizedLine, " ")

			train[header[j]][i] = lettersLine

		}

		fmt.Printf("Done %d/%d...\r", i, 25000)
	}

	for i := 0; i < 50; i++ {
		fmt.Println("=============================================================")
		fmt.Printf("Review nr. %d: %v\n", i+1, train[header[2]][i])
		fmt.Println("=============================================================")
	}
}
