package main

import (
	"log"
	"os"
	"strings"
)

func ReadingFromFile(filename string) GameBoard {
	filecontents, err := os.ReadFile(filename)

	Check(err)
	giantString := string(filecontents)

	trimmedGiantString := strings.TrimSpace(giantString)

	lines := strings.Split(trimmedGiantString, "\n")

	board := make(GameBoard, len(lines))

	for i,currentline := range lines {
		
		linelements := strings.Split(currentline, ",")

		board[i] = SetRowLines(linelements)
	}
	return board
}

func SetRowLines(linelements []string) []bool {
	rowlength := len(linelements)

	currentRow := make([]bool, rowlength)

	for j,val := range linelements {
		if val == "0" {
			currentRow[j] = false
		} else if val == "1" {
			currentRow[j] = true
		} else {
			log.Fatal("error")
		}
	}
	return currentRow
}

func Check(e error) {
	if e!=nil {
		log.Fatal(e)
	}
}