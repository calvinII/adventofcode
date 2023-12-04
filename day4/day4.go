package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	winningNumbers []int
	numbersOnCard  map[int]int
	matches        int
	stack          int
}

func main() {
	var lines []line

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	result := uint64(0)
	for fileScanner.Scan() {
		numberString := strings.Split(strings.Split(fileScanner.Text(), ":")[1], "|")
		var wn []int
		noc := make(map[int]int)
		for _, number := range strings.Fields(numberString[0]) {
			n, _ := strconv.Atoi(number)
			wn = append(wn, n)
		}
		for _, number := range strings.Fields(numberString[1]) {
			n, _ := strconv.Atoi(number)
			noc[n] = n
		}
		res := -1
		for _, t := range wn {
			_, ok := noc[t]
			if ok {
				res += 1
			}
		}
		if res >= 0 {
			result += 1 << uint64(res)
		}
		lines = append(lines, line{wn, noc, res, 1})
	}

	fmt.Println("Silver Result: ", result)

	for n, line := range lines {
		for i := n; i < n+line.matches+1; i++ {
			if i+1 < len(lines) {
				lines[i+1].stack += line.stack
			}
		}
	}
	cards :=0
	for _,line := range lines {
		cards += line.stack
	}
	fmt.Println("Gold Result: ", cards)
	readFile.Close()
}
