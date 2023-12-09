package main

// !!!! don't try this at home, your computer can explode :O
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

func main() {
	silver :=0
	steps := make(map[string]map[byte]string)
	directions := readLines()[0]
	expr := regexp.MustCompile(`[A-Z]{3}`)
	for _, line := range readLines()[2:len(readLines())] {
		step := expr.FindAllString(line, -1)
		fmt.Println(step)
		if len(step) == 0 {
			continue
		}
		stepmap := map[byte]string{'L': step[1], 'R': step[2]}
		steps[step[0]] = stepmap
	}

	target := "AAA"
	fmt.Println("Starting with: ", target)
	for i := 0; i < len(directions); i++ {
		fmt.Println(target,steps[target])
		if target == "ZZZ" {
			break
		}
		fmt.Printf("%c",directions[i])
		silver += 1
		target = steps[target][directions[i]]
		if i == len(directions)-1 {
			i = -1
			fmt.Println()
		}

	}
	fmt.Println()
	fmt.Println("Directions: ", directions, len(directions))
	fmt.Println("Steps: ", steps)
	fmt.Println("Silver: ", silver)

}
