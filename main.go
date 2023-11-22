package main

import "fmt"

func main() {
	//slices: are like arrays but the no of vars not defined

	nba_teams := []string{"Boston", "Lakers", "GSW"}

	fmt.Printf("He only know these %v NBA teams: %v \n", len(nba_teams), nba_teams)

	nba_teams = append(nba_teams, "76sers")

	fmt.Printf("He only know these %v NBA teams: %v \n", len(nba_teams), nba_teams)

	nba_teams = append(nba_teams, "SeaHawks")

	fmt.Printf("He only know these %v NBA teams: %v \n", len(nba_teams), nba_teams)

	fevteams := nba_teams[0:2]

	fmt.Printf("He only know these %v NBA teams: %v but he likes these %v more %v\n", len(nba_teams), nba_teams, len(fevteams), fevteams)

}
