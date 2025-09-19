package main

import (
	"os" // for reading from files
	"strconv" // our old friend string conversion package
	"strings" // for working with strings
)


func ReadElectoralVotes(file string) map[string]uint {
	electoralVotes := make(map[string]uint)

	fileContents, err := os.ReadFile(file)
	check(err)

	giantString := string(fileContents)
	lines := strings.Split(giantString, "\n")
	for _,line := range lines {
		line := strings.Split(line, ",")
		str := line[0]
		vote, err := strconv.Atoi(line[1])
		check(err)
		electoralVotes[str] = uint(vote)
	}

	return electoralVotes
}

func ReadPollingData(file string) map[string]float64 {
	Candidate1Percentages := make(map[string]float64)

	fileContents,err := os.ReadFile(file)
	check(err)
	giantString := string(fileContents)
	lines := strings.Split(giantString, "\n")

	for _,line := range lines {
		line := strings.Split(line, ",")
		state := line[0]
		per, err := strconv.ParseFloat(line[1], 64)
		check(err)
		Candidate1Percentages[state] = per/100.0
	}

	return Candidate1Percentages
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}