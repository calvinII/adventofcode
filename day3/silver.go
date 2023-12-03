package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type numberType struct {
	number int
	length int
}

type numberMap map[int]numberType

type partMap map[int]string

func isGear(column int,line int, n []numberMap, lines int, linelength int) uint64 {
	var adjacentNumbers []int
	fmt.Println("testing",column,line)
	// left of the position?
	//1000-9999
	for leftlength := 4; leftlength > 0 ; leftlength-- {
		_,ok := n[line][column-leftlength] 
		if ok {
			if n[line][column-leftlength].length == leftlength {
				adjacentNumbers = append(adjacentNumbers, n[line][column-leftlength].number)
				break
			}	
		}	
			}
	// right of the position
	_,ok := n[line][column+1]
	if ok {
		adjacentNumbers = append(adjacentNumbers, n[line][column+1].number)
	}
	// line above
	if line >0 {
		for leftlength := 3; leftlength >= 0 ; leftlength-- {
			_,ok := n[line-1][column-leftlength] 
			if ok {
				if n[line-1][column-leftlength].length >= leftlength {
					adjacentNumbers = append(adjacentNumbers, n[line-1][column-leftlength].number)
					break
				}	
		}
	}
		//number starting right top
		_,ok := n[line-1][column+1]
		if ok {
			adjacentNumbers = append(adjacentNumbers, n[line-1][column+1].number)
		}
	
	}
	// line below
	if line < lines -1 {
		for leftlength := 3; leftlength >= 0 ; leftlength -- {
			_,ok := n[line+1][column-leftlength] 
			if ok {
				if n[line+1][column-leftlength].length >= leftlength {
					adjacentNumbers = append(adjacentNumbers, n[line+1][column-leftlength].number)
					break
				}	
		}
	}
		//number starting right top
		_,ok := n[line+1][column+1]
		if ok {
			adjacentNumbers = append(adjacentNumbers, n[line+1][column+1].number)
		}

	}
	if len(adjacentNumbers) == 2 {
		fmt.Println("Numbers:",adjacentNumbers)
		fmt.Println(column,line)
		return uint64(adjacentNumbers[0]*adjacentNumbers[1])
	} else {
		fmt.Println("No Gear",column,line)

		return 0
	}
}

func isAdjacentPart(n numberType, parts []partMap, lines int, linelength int, column int, line int) int {
	result := false
	// check line above and below
	// parts[line -1][column-1] bis parts[line-1][column+length]
	if line > 0 {
		for key := column - 1; key <= column+n.length; key++ {
			_, ok := parts[line-1][key]
			if ok {
				result = true
				break
			}
		}
	}
	// check left and right of number
	// parts[line][column-1] und parts[line][column+length]
	if column > 0 && !result {
		_, ok := parts[line][column-1]
		if ok {
			result = true
		}
	}

	if column+n.length < linelength-1 && !result {
		_, ok := parts[line][column+n.length]
		if ok {
			result = true
		}

	}
	// if line < lines -1
	// parts[line +1][column-1] bis parts[line+1][column+length]
	if line < lines-1 && !result {
		for key := column - 1; key <= column+n.length; key++ {
			_, ok := parts[line+1][key]
			if ok {
				result = true
				break
			}
		}
	}
	if result {
		return n.number
	} else {
		return 0
	}

}

func main() {

	var (
		numbers []numberMap
		parts   []partMap
		gears	[]partMap
		plan    [140]string
	)
	line := 0

	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		numberMatch := regexp.MustCompile(`[0-9]+`)
		partMatch := regexp.MustCompile(`[^0-9.]`)

		foundNumbers := numberMatch.FindAllString(fileScanner.Text(), -1)
		numberPositions := numberMatch.FindAllStringIndex(fileScanner.Text(), -1)
		foundParts := partMatch.FindAllString(fileScanner.Text(), -1)
		partPositions := partMatch.FindAllStringIndex(fileScanner.Text(), -1)
		nm := make(numberMap)
		pm := make(partMap)

		for index, num := range foundNumbers {
			convertedNumber, _ := strconv.Atoi(num)
			startPosition := numberPositions[index][0]
			length := len(num)
			nm[startPosition] = numberType{convertedNumber, length}
		}
		numbers = append(numbers, nm)

		for index, part := range foundParts {
			partChar := part
			partPosition := partPositions[index][0]
			pm[partPosition] = partChar
		}
		parts = append(parts, pm)

		plan[line] = fileScanner.Text()
		line += 1

	}
	fmt.Println(len(numbers), numbers)
	fmt.Println(len(parts), parts)
	for _,partLine := range parts {
		gm := make(partMap)
		for key,value := range partLine {
			if value == "*" {
				gm[key] = value
			}
		}
		gears = append(gears, gm)
	}
	sum := 0
	for line,nm := range numbers {
		for column,num := range nm {
			sum += isAdjacentPart(num,parts, 140, 140,column,line)
		}
	}
	fmt.Printf("Silver solution is: %d\n",sum)
	ratio := uint64(0)
	for line,gm := range gears {
		for column,_ := range gm {
			ratio += isGear(column, line, numbers, 140, 140)
		}
	}
	fmt.Printf("Gold Solution: %d",ratio)
	readFile.Close()
}
