package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, file := GetData("input.txt")
	defer file.Close()
	sumOfValidGames := 0
	sumOfPowers := 0

	for data.Scan() {
		gameNumber, listOfCubes := getRGB(data.Text())
		var gameCubes []game
		for _, turn := range listOfCubes {
			game := parseTurn(turn)
			gameCubes = append(gameCubes, game)
		}
		constraint := game{
			red:   12,
			blue:  14,
			green: 13,
		}
		isGameValid := validateGame(constraint, gameCubes)

		if isGameValid {
			sumOfValidGames += gameNumber
		}
		// p2

		minCube := getMinCubesForGame(gameCubes)
		power := getPower(minCube)
		sumOfPowers += power

	}
	fmt.Printf("sumOfValidGames: %v\n", sumOfValidGames)
	fmt.Printf("sumOfPowers: %v\n", sumOfPowers)
}

func getPower(min game) int {
	return min.blue * min.green * min.red
}

func getMinCubesForGame(games []game) game {
	var min game

	for _, turn := range games {
		if turn.red > min.red {
			min.red = turn.red
		}
		if turn.blue > min.blue {
			min.blue = turn.blue
		}
		if turn.green > min.green {
			min.green = turn.green
		}
	}

	return min
}

func validateGame(constraint game, games []game) bool {
	for _, turn := range games {
		if turn.blue > constraint.blue {
			return false
		}
		if turn.green > constraint.green {
			return false
		}
		if turn.red > constraint.red {
			return false
		}
	}
	return true
}

func parseTurn(turn string) game { // converts the given string as a game struct
	var g game
	turn = strings.TrimSpace(turn)
	cubes := strings.Split(turn, ",")
	for _, cube := range cubes {
		cube = strings.TrimSpace(cube)
		qtyColour := strings.Split(cube, " ")
		qty, _ := strconv.Atoi(qtyColour[0])

		colour := qtyColour[1]
		switch colour {
		case "blue":
			g.blue = qty
		case "red":
			g.red = qty
		case "green":
			g.green = qty
		}
	}

	return g
}

func getRGB(game string) (int, []string) { // parse gamestring
	gameRGBsplit := strings.Split(game, ":")
	gameNum := gameRGBsplit[0]
	gameCubes := gameRGBsplit[1]
	numSplit := strings.Split(gameNum, " ")
	num := numSplit[1]
	numInt, _ := strconv.Atoi(num)
	gameCubes = strings.TrimSpace(gameCubes)
	cubes := strings.Split(gameCubes, ";")
	return numInt, cubes
}

type game struct {
	red   int
	blue  int
	green int
}

func GetData(filename string) (*bufio.Scanner, *os.File) {
	readfile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner, readfile
}
