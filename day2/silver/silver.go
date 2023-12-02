package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	red int
	green int
	blue int
}

type gameID struct {
	id int
	games []game
}

func successfulGame(games []game, toMatch game) bool {
	result := true
	for _,testGame := range games {
		if testGame.red > toMatch.red || testGame.blue > toMatch.blue || testGame.green > toMatch.green {
			fmt.Println(testGame)
			result = false
			break
		}
	}
	return result
}

func main() {
    readFile, err := os.Open("data.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var allGames []gameID
	sum := 0

	for fileScanner.Scan() {
		splitLine := strings.Split(fileScanner.Text(),":")
		gameNumber, _ := strconv.Atoi(strings.Join(regexp.MustCompile("[0-9]+").FindAllString(splitLine[0], -1), ""))
		var playedGames []game
		for _,selectedGame := range strings.Split(splitLine[1],";") {
			color := map[string]int {
				"red": 0,
				"blue": 0,
				"green": 0,
			}
			for _,element := range strings.Split(selectedGame,","){
				indexed := strings.Split(element," ")
				color[indexed[2]],_ = strconv.Atoi(indexed[1])
			}
			playedGames = append(playedGames, game { color["red"], color["green"], color["blue"] })
		}
		allGames = append(allGames, gameID { gameNumber,playedGames})


	}
	aocgame := game{12,13,14}
	for _,test := range allGames {
		fmt.Println(test.id)
		if successfulGame(test.games, aocgame) {
			
			sum += test.id
			fmt.Println(test.id,sum)
		}
	}

	fmt.Println(allGames,sum)

    readFile.Close()
}
