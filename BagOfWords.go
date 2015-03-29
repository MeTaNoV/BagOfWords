package main

import (
	"fmt"
	"os"
	//	"io"
	"encoding/csv"
	"github.com/MeTaNoV/snowball"
	"github.com/kennygrant/sanitize"
	"regexp"
	"strings"
)

const MAX = 2500 // 25000 at the end

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
		train[header[i]] = make([]string, MAX)
	}

	// Step 1: we extract all review, remove HTML tags and unwanted character
	for i := 0; i < MAX; i++ {
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

			// Then we make all lowercase
			train[header[j]][i] = strings.ToLower(lettersLine)
		}

		fmt.Printf("Step 1: Done %5d/%5d...\r", i+1, MAX)
	}

	fmt.Println("")

	// Step 2: we stemm the words removing the stop words
	stemmedReviews := make([][]string, 0)
	countStopped := 0
	words := 0

	for i := 0; i < MAX; i++ {
		stemmedWords := make([]string, 0)
		wordsToStem := strings.Split(train[header[2]][i], " ")

		for j := range wordsToStem {
			if nok, _ := snowball.IsStopWord(wordsToStem[j], "english"); !nok {
				wordToStem, _ := snowball.Stem(wordsToStem[j], "english", true)
				stemmedWords = append(stemmedWords, wordToStem)
			} else {
				countStopped++
			}
		}

		words += len(stemmedWords)
		stemmedReviews = append(stemmedReviews, stemmedWords)

		fmt.Printf("Step 2: Done %5d/%5d... (%7d words, %7d stopped)\r", i+1, MAX, len(wordsToStem), countStopped)
	}

	fmt.Println("")
	fmt.Printf("Cleaned reviews: %5d, Total words stemmed: %7d", len(stemmedReviews), words)

	// Step 3: Create the afm file to be used for test with CloudForest

	fmt.Println("\nEND!")
}
