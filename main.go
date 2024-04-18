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

func countWords(filename string) int {
	contents, err := readFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	words := strings.Fields(contents)
	return len(words)
}

func countLines(filename string) int {
	contents, err := readFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(contents, "\n")
	return len(lines) - 1
}

func countCharacters(filename string) int {
	contents, err := readFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	return len(contents)
}

func main() {
	// fmt.Println(contents)
	linesFlag := flag.Bool("l", false, "Count lines")
	wordsFlag := flag.Bool("w", false, "Count words")
	charactersFlag := flag.Bool("c", false, "Count characters")

	flag.Parse()

	filename := flag.Arg(0)

	if !*linesFlag && !*wordsFlag && !*charactersFlag {
		fmt.Printf("%8d%8d%8d %s\n", countLines(filename), countWords(filename), countCharacters(filename), filename)
		return;
	}

	if *linesFlag {
		fmt.Printf("%8d", countLines(filename))
	}

	if *wordsFlag {
		fmt.Printf("%8d", countWords(filename))
	}

	if *charactersFlag {
		fmt.Printf("%8d", countCharacters(filename))
	}

	fmt.Printf(" %s\n", filename)

	// fmt.Printf("%8d%8d%8d %s\n", countLines(filename), countWords(filename), countCharacters(filename), filename)
	// fmt.Printf("%8d %s\n", countLines(filename), filename)
	// fmt.Printf("%8d %s\n", countWords(filename), filename)
	// fmt.Printf("%8d %s\n", countCharacters(filename), filename)
}
