package main

import (
	"log"
	"os"
	"strings"
	"strconv"
)

// ReadBoardFromFile takes a filename as a string and reads in the data provided
// in this file, returning a game board.
func ReadBoardFromFile(filename string) GameBoard {
    // try reading in the file
    fileContents, err := os.ReadFile(filename)
    Check(err)

    // parse out the file contents as a slice of strings, one line per element
    giantString := string(fileContents)
    trimmedGiantString := strings.TrimSpace(giantString)
    lines := strings.Split(trimmedGiantString, "\n")

    // we know how many rows the GameBoard will have
    board := make(GameBoard, len(lines))

    // next, parse the data on each line and add to the GameBoard
    for i, currentLine := range lines {
        lineElements := strings.Split(currentLine, ",")

        // use a subroutine to set the values of the current row
        board[i] = SetRowValues(lineElements)
    }

    return board
}

// SetRowValues takes as input a slice of strings, where each element is a nonnegative integer.
// It parses these strings into a slice of integers.

// ReadRulesFromFile takes a filename as a string and reads the rule strings provided in this file.
// It stores the result in a map of strings to integers.

// Check takes as input a variable of error type.
// If the variable has any value other than nil, it panics.
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// SetRowValues takes as input a slice of strings,
// where each element is either "0" or "1".
// It returns a slice of boolean values of the same length,
// encoding "0" as false and "1" as true.
// SetRowValues takes as input a slice of strings,
// where each element is a nonnegative integer.
// It parses these strings into a slice of integers.
func SetRowValues(lineElements []string) []int {
    currentRow := make([]int, len(lineElements))
    //range through the current line elements and set values
    for i := range currentRow {
        val, err := strconv.Atoi(lineElements[i])
        Check(err)
        currentRow[i] = val
    }
    return currentRow
}

// ReadRulesFromFile takes a filename as a string and reads the rule strings provided in this file.
// It stores the result in a map of strings to integers.
func ReadRulesFromFile(filename string) map[string]int {
    // create the map to store the rules
    ruleStrings := make(map[string]int)

    // try reading in the file
    fileContents, err := os.ReadFile(filename)

    Check(err)

    // parse out the file contents as a slice of strings, one line per element
    giantString := string(fileContents)
    trimmedGiantString := strings.TrimSpace(giantString)
    lines := strings.Split(trimmedGiantString, "\n")

    // next, parse the data on each line and add to the ruleStrings map
    for _, currentLine := range lines {
        parts := strings.Split(currentLine, ":")
        if len(parts) == 2 {
            // first element is the neighborhood string, and the second element is the new state
            neighborhoodString := parts[0]
            newState, err := strconv.Atoi(parts[1])
            Check(err)
            ruleStrings[neighborhoodString] = newState
        }
    }

    return ruleStrings
}