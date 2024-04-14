package main

import (
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

func countLines(filename string) int {
	contents, err := readFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(contents, "\n")
	return len(lines) - 1
}

func main() {
	filename := os.Args[1]
	// fmt.Println(contents)
	fmt.Printf("%8d %s\n", countLines(filename), filename)
}
