package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func formatPlayers(g game) string {
	playersList := ""
	for i, player := range g.players {
		prefix := ""
		if i > 0 {
			prefix = ", "
		}
		playersList = fmt.Sprintf("%v%v\"%v\"", playersList, prefix, player)
	}
	return playersList
}

func initializePlayer(g *game, name string) {
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
var reKill = regexp.MustCompile("Kill: .+: (.+) killed (.+) by (.+)")

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
		} else if m := reClientChanged.FindStringSubmatch(line); m != nil {
			name := m[1]
			// fmt.Printf("Found name %v\n", name)
			if g != nil {
				initializePlayer(g, name)
			}
		} else if m := reKill.FindStringSubmatch(line); m != nil {
			g.kills = g.kills + 1
			killer := m[1]
			victim := m[2]
			//mode := m[3]
			if g != nil {
				fmt.Printf("Found kill, killer: %v, victim: %v\n", killer, victim)
				initializePlayer(g, victim)
				if killer != "<world>" {
					initializePlayer(g, killer)
					if killer != victim {
						g.scoreboard[killer]++
					}

				} else {
					g.scoreboard[victim]--

				}

			}

		}
	}
	for _, game := range table {
		fmt.Printf("Game: %v\n", game.id)
		fmt.Printf("  Kills: %v\n", game.kills)
		fmt.Printf("  Players: %v\n", formatPlayers(game))
		fmt.Printf("  Scoreboard:\n")
		players := make([]string, len(game.players))
		copy(players, game.players)
		sort.Slice(players, func(a, b int) bool {
			playerNameA := players[a]
			playerNameB := players[b]
			return game.scoreboard[playerNameA] > game.scoreboard[playerNameB]
		})
		for _, p := range players {
			fmt.Printf("    \"%v\": %v\n", p, game.scoreboard[p])
		}
	}

}
