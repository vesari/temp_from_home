package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	// "strings"
)

// func search() map[string]int {
//
//
// 	games := strings.Split(string(content), "InitGame")
// 	//qui devo trovare un modo di fare una mappa che classifichi ogni partita con le relative statistiche
// 	//playerCounts := strings.Count(string(games[1]), "ClientUserinfoChanged: 2 n\")
//
//
// 	//devo trovare un modo per identificare i giocatori per nome ed assegnargli statistiche personali poi da sommare
// 	scoreboard := make(map[string]int)
// 	//scoreboard ["Player1"] = 0
// 	for i, game := range games {
//
//     killCounts := strings.Count(string(games[i]), "killed")
//
//     logs := strings.Split(game, "\n")
//
//
// 		playerLogs := []string{}
//
// 		for _, log := range logs {
// 			if strings.Contains(log, "ClientUserinfoChanged:") {
// 				playerLogs = append(playerLogs, log)
// 			}
//
// 			return playerLogs
// 		}
// 	}
//
// 	return scoreboard
//
// }

func formatPlayers(g game) string {
	playersList := ""
	for _, player := range g.players {
		playersList = fmt.Sprintf("%v\"%v\"", playersList, player)
	}
	return playersList
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type game struct {
	id         int
	kills      int
	players    []string
	scoreboard map[string]int
}

func main() {
	content, err := ioutil.ReadFile("prova.txt")
	check(err)
	splitContent := strings.Split(string(content), "\n")
	currentGame := 0
	table := make([]game, 0)
	for _, line := range splitContent {
		//fmt.Printf("current line: %v\n", line)
		if strings.Contains(line, "InitGame") {
			currentGame++
			fmt.Printf("Appending game with ID %v\n", currentGame)
			table = append(table, game{
				id: currentGame,
			})
			//fmt.Printf("the current game is now %v\n", currentGame)
		}
		if strings.Contains(line, "killed") {
			g := &table[len(table)-1]
			g.kills = g.kills + 1
			fmt.Printf("Found a kill: %v, %v\n", g.id, g.kills)

		}
	}
	for _, game := range table {
		fmt.Printf("Game: %v\n", game.id)
		fmt.Printf("  Kills: %v\n", game.kills)
		fmt.Printf("  Players: %v\n", formatPlayers(game))
		fmt.Printf("  Scoreboard:\n")
		for k, v := range game.scoreboard {
			fmt.Printf("\"%v\": %v\n", k, v)
		}
	}

}
