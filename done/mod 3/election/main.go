package main

import (
	"fmt"
	"Math/rand"
)

func main() {
	fmt.Println("Let's simulate an election!")

	electoralVoteFile := "data/electoralVotes.csv"
	// pollFile := "data/earlyPolls.csv"
	// pollFile := "data/conventions.csv"
	pollFile := "data/debates.csv"

	electoralVotes := ReadElectoralVotes(electoralVoteFile)
	polls := ReadPollingData(pollFile)

	numTrials := 1000000
	marginOfError := 0.1

	probability1, probability2, ProbabilityTie := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)

	fmt.Println("Probability of Candidate 1 winning: ", probability1)
	fmt.Println("Probability of Candidate 2 winning: ", probability2)
	fmt.Println("Probability of tie: ", ProbabilityTie)
}

func SimulateMultipleElections(polls map[string]float64,
electoralVotes map[string]uint, numTrials int,
marginOfError float64) (float64,float64, float64) {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	for i:=0; i < numTrials; i++ {
		votes1, votes2 := SimulateOneElection(polls,electoralVotes,marginOfError)

		if votes1 > votes2 {
			winCount1++
		} else if votes1 < votes2 {
			winCount2++
		} else {
			tieCount++
		}
	}

	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	ProbabilityTie := float64(tieCount) / float64(numTrials)

	return probability1, probability2, ProbabilityTie
}

func SimulateOneElection(polls map[string]float64,
electoralVotes map[string]uint,
marginOfError float64) (uint,uint) {
	var collegeVotes1 uint = 0
	var collegeVotes2 uint = 0

	for state, pollingValue := range polls {
		numVotes := electoralVotes[state]

		adjustedPoll:= AddNoise(pollingValue, marginOfError)

		if adjustedPoll >= 0.5 {
			collegeVotes1 += numVotes
		} else {
			collegeVotes2 += numVotes
		}
	}
	return collegeVotes1, collegeVotes2
}

func AddNoise(pollingValue, marginOfError float64) float64 {
	x := rand.NormFloat64() 
	x/=2.0
	x*=marginOfError
	return x + pollingValue
}