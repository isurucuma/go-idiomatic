package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {
	teamsNames := make([]string, 0, len(l.Teams))
	sort.Slice(l.Teams, func(i, j int) bool { return l.Wins[l.Teams[i].Name] > l.Wins[l.Teams[j].Name] })
	for _, Team := range l.Teams {
		teamsNames = append(teamsNames, Team.Name)
	}
	return teamsNames
}

func main() {
	// Create teams
	team1 := Team{
		Name:    "Team 1",
		Players: []string{"Player 1", "Player 2", "Player 3"},
	}
	team2 := Team{
		Name:    "Team 2",
		Players: []string{"Player 4", "Player 5", "Player 6"},
	}

	// Create league
	league := League{
		Teams: []Team{team1, team2},
		Wins:  make(map[string]int),
	}

	// Record match results
	league.MatchResult("Team 1", 3, "Team 2", 2)
	league.MatchResult("Team 2", 4, "Team 1", 1)
	league.MatchResult("Team 2", 3, "Team 1", 0)

	// Get ranking
	ranking := league.Ranking()
	fmt.Println("Ranking:")
	for i, team := range ranking {
		fmt.Printf("%d. %s\n", i+1, team)
	}
}
