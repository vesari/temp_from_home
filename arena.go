package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

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

var reClientChanged = regexp.MustCompile("ClientUserinfoChanged: .+ n\\\\([^\\\\]+)")

func main() {
	content, err := ioutil.ReadFile("prova.txt")
	check(err)
	splitContent := strings.Split(string(content), "\n")
	currentGame := 0
	table := make([]game, 0)
	for _, line := range splitContent {
		var g *game
		if len(table) > 0 {
			g = &table[len(table)-1]
		}

		//fmt.Printf("current line: %v\n", line)
		if strings.Contains(line, "InitGame") {
			currentGame++
			fmt.Printf("Appending game with ID %v\n", currentGame)
			table = append(table, game{
				id:         currentGame,
				scoreboard: make(map[string]int),
			})
			//fmt.Printf("the current game is now %v\n", currentGame)
		}
		if strings.Contains(line, "killed") {

			g.kills = g.kills + 1
			fmt.Printf("Found a kill: %v, %v\n", g.id, g.kills)

		} else if m := reClientChanged.FindStringSubmatch(line); m != nil {
			name := m[1]
			fmt.Printf("Found name %v\n", name)
			if g != nil {
				foundPlayer := false
				for _, p := range g.players {
					if name == p {
						foundPlayer = true
						break
					}
				}
				if !foundPlayer {
					g.players = append(g.players, name)
					g.scoreboard[name] = 0
				}
			}
		}
	}
	for _, game := range table {
		fmt.Printf("Game: %v\n", game.id)
		fmt.Printf("  Kills: %v\n", game.kills)
		fmt.Printf("  Players: %v\n", formatPlayers(game))
		fmt.Printf("  Scoreboard:\n")
		for k, v := range game.scoreboard {
			fmt.Printf("    \"%v\": %v\n", k, v)
		}
	}

}
