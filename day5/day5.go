package main
// !!!! don't try this at home, your computer can explode :O
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type transformationline [][3]uint64

func main() {
	var maps []string
	data := make(map[string]transformationline)
	readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	var seeds []uint64
	seedstring := strings.Fields(strings.Split(fileScanner.Text(), ":")[1])
	fmt.Println("Länge: ", len(seeds))
	match := ""
	mapstring := regexp.MustCompile("(.*) map:")
	for fileScanner.Scan() {
		next := mapstring.FindString(fileScanner.Text())
		if next != "" {
			fmt.Println(next)
			match = mapstring.ReplaceAllString(fileScanner.Text(), "$1")
			maps = append(maps, match)
		} else {
			var value [3]uint64
			for i, v := range strings.Fields(fileScanner.Text()) {
				value[i], _ = strconv.ParseUint(v, 10, 64)
			}
			fmt.Println(value)
			data[match] = append(data[match], value)
		}

	}
	result := uint64(18446744073709551615)
	for i := 0; i < len(seedstring); i += 2 {
		s, _ := strconv.ParseUint(seedstring[i], 10, 64)
		l, _ := strconv.ParseUint(seedstring[i+1], 10, 64)
		for j := s; j < s+l; j++ {
			value := j
			for _, seedmap := range maps {
				for _, combination := range data[seedmap] {
					if combination[1] <= value && value < combination[1]+combination[2] {
						value += combination[0] - combination[1]
						break
					}
				}
			}
			if result > value {
				result = value
				fmt.Println("new value: ",result)
			}

		}
	}


	fmt.Println("Gold Result: ", result)
	readFile.Close()
}
