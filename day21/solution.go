package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
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
		{c.x - 1, c.y},
		{c.x + 1, c.y},
		{c.x, c.y - 1},
		{c.x, c.y + 1},
	}
}

func fields(s *map[coord]string, garden *map[coord]string, width int, height int, padding coord) *map[coord]string {
	result := make(map[coord]string)
	for k, _ := range *s {
		for _, co := range k.neighs() {
			mCo := coord{(co.x + padding.x) % width, (co.y + padding.y) % height}
			// fmt.Println(co, ":", mCo, (*garden)[mCo])
			if (*garden)[mCo] == "." {
				result[co] = "O"
			}
		}
	}
	return &result
}

func main() {

	gpMatch := regexp.MustCompile(`\.`)
	rMatch := regexp.MustCompile(`#`)
	sMatch := regexp.MustCompile(`S`)
	start := coord{}
	foundStart := false
	garden := make(map[coord]string)
	startMap := make(map[coord]string)
	width := len(readLines()[0])
	height := len(readLines())
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
				start = coord{i, s[0]}
				garden[start] = "."
				startMap[start] = "O"
				foundStart = true
			}
		}
	}
	startTime := time.Now()
	sMap := &startMap
	steps := 524
	padding := coord{steps * width, steps * height}
	for i := 1; i <= steps; i++ {
		sMap = fields(sMap, &garden, width, height, padding)
		// fmt.Println(sMap)
	}
	elapsed := time.Since(startTime)

	// data is 131 x 131 
	// steps to walk: 26501365
	// 26501365/131 = 202300
	// 202300x202300 data fileds

	fmt.Println(startMap)
	fmt.Println(garden)
	fmt.Println(start)
	fmt.Println("silver:", len(*sMap))
	fmt.Println("elapsed: ", elapsed)

}
