package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type Combination struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id           int
	Combinations []Combination
	Possible     bool
}

func parseInput(filename string) ([]Game, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(data, []byte("\n"))
	games := make([]Game, 0, len(lines))
	for _, l := range lines {
		strLine := string(l)
		game := Game{Possible: true}

		fmt.Sscanf(strLine, "Game %d:", &game.Id)
		strLine = strings.TrimPrefix(strLine, fmt.Sprintf("Game %d:", game.Id))

		strCombos := strings.Split(strLine, ";")
		game.Combinations = make([]Combination, len(strCombos))
		for i, combo := range strCombos {
			for _, pair := range strings.Split(combo, ",") {
				var value int
				var color string
				fmt.Sscanf(pair, "%d %s", &value, &color)
				switch color {
				case "red":
					{
						game.Combinations[i].Red = value
						if value > maxRed {
							game.Possible = false
						}
					}
				case "green":
					{
						game.Combinations[i].Green = value
						if value > maxGreen {
							game.Possible = false
						}
					}
				case "blue":
					{
						game.Combinations[i].Blue = value
						if value > maxBlue {
							game.Possible = false
						}
					}
				}
			}
		}

		games = append(games, game)
	}

	return games, nil
}

func part1(games []Game) int {
	sum := 0
	for _, game := range games {
		if game.Possible {
			sum += game.Id
		}
	}
	return sum
}

func part2(games []Game) int {
	comboOfMaxes := make([]Combination, len(games))

	for i, game := range games {
		for _, combo := range game.Combinations {
			if combo.Red > comboOfMaxes[i].Red {
				comboOfMaxes[i].Red = combo.Red
			}
			if combo.Green > comboOfMaxes[i].Green {
				comboOfMaxes[i].Green = combo.Green
			}
			if combo.Blue > comboOfMaxes[i].Blue {
				comboOfMaxes[i].Blue = combo.Blue
			}
		}
	}

	result := 0
	for _, combo := range comboOfMaxes {
		result += combo.Red * combo.Green * combo.Blue
	}

	return result
}

func main() {
	games, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(games))
	fmt.Println(part2(games))
}
