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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	gold := 0
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

	var target []string

	are := regexp.MustCompile(`..A`)
	zre := regexp.MustCompile(`..Z`)

	for key, _ := range steps {
		if are.MatchString(key) {
			target = append(target, key)
		}
	}

	fmt.Println("Starting with: ", target)
	var distances []int
	for _, t := range target {
		gold = 0
		for i := 0; i < len(directions); i++ {
			if zre.MatchString(t) {
				break
			}
			gold += 1
			t = steps[t][directions[i]]
			if i == len(directions)-1 {
				i = -1
			}
		}
		distances = append(distances, gold)

	}
	fmt.Println("Distances: ",distances)
	fmt.Println(LCM(distances[0],distances[1],distances...))

}
