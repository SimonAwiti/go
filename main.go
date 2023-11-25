package main

import (
	"fmt"
)

func main() {
	nba_teams := []string{"Lakera", "Celtics", "Mavs", "Suns"}

	nba_teams = append(nba_teams, "GSW")

	fev_teams := nba_teams[0:3]

	nba_teams = append(nba_teams, "76sers")

	statement := fmt.Sprintf("He is familiar with all the %v teams: %v, but then he like these %v teams: %v ", len(nba_teams), nba_teams, len(fev_teams), fev_teams)

	fmt.Println(statement)

}
