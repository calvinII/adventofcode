package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}

type coord struct {
	x, y int
}

func (c *coord) neighs() []coord {
	return []coord{
		coord{c.x - 1, c.y},
		coord{c.x + 1, c.y},
		coord{c.x, c.y - 1},
		coord{c.x, c.y + 1},
	}
}

func fields(s *map[coord]string, garden *map[coord]string) map[coord]string {
	result := make(map[coord]string)
	for k, _ := range *s {
		for _, co := range k.neighs() {
			if (*garden)[co]=="." {
				result[co] ="O"
			}
		}
	}
	return result
}

func main() {

	gpMatch := regexp.MustCompile(`\.`)
	rMatch := regexp.MustCompile(`#`)
	sMatch := regexp.MustCompile(`S`)
	start := coord{}
	foundStart := false
	garden := make(map[coord]string)
	startMap := make(map[coord]string)
	for i, l := range readLines() {
		for _, j := range gpMatch.FindAllStringIndex(l, -1) {
			garden[coord{i, j[0]}] = "."
		}
		for _, j := range rMatch.FindAllStringIndex(l, -1) {
			garden[coord{i, j[0]}] = "#"
		}
		if !foundStart {
			s := sMatch.FindStringIndex(l)
			if s != nil {
				start = coord{s[0], i}
				garden[start] = "."
				startMap[start] = "O"
				foundStart = true
			}
		}
	}

	for i :=1 ; i <=64 ; i++ {
		startMap = fields(&startMap, &garden)
	}
	fmt.Println(startMap)
	fmt.Println(garden)
	fmt.Println(start)
	fmt.Println("silver:",len(startMap))

}
