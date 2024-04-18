package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) (string, error) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(fileData), nil
}

func countWords(fileContents string) int {
	words := strings.Fields(fileContents)
	return len(words)
}

func countLines(fileContents string) int {
	lines := strings.Split(fileContents, "\n")
	return len(lines) - 1
}

func countCharacters(fileContents string) int {
	return len(fileContents)
}

func main() {
	// fmt.Println(contents)
	linesFlag := flag.Bool("l", false, "Count lines")
	wordsFlag := flag.Bool("w", false, "Count words")
	charactersFlag := flag.Bool("c", false, "Count characters")

	flag.Parse()

	files := flag.Args()

	var totalLines, totalWords, totalCharacters int
	for _, filename := range files {
		fileContents, err := readFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		if !*linesFlag && !*wordsFlag && !*charactersFlag {
			fmt.Printf("%8d%8d%8d %s\n", countLines(fileContents), countWords(fileContents), countCharacters(fileContents), filename)
			continue
		}

		if *linesFlag {
			lines := countLines(fileContents)
			totalLines += lines
			fmt.Printf("%8d", lines)
		}

		if *wordsFlag {
			words := countWords(fileContents)
			totalWords += words
			fmt.Printf("%8d", words)
		}

		if *charactersFlag {
			characters := countCharacters(fileContents)
			totalCharacters += characters
			fmt.Printf("%8d", characters)
		}

		fmt.Printf(" %s\n", filename)
	}

	if *linesFlag {
		fmt.Printf("%8d", totalLines)
	}

	if *wordsFlag {
		fmt.Printf("%8d", totalWords)
	}

	if *charactersFlag {
		fmt.Printf("%8d", totalCharacters)
	}

	fmt.Printf(" total\n")
}
