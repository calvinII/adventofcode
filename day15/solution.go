package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type lens string

type slot struct {
	fl    int
	order int
}

func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}

func (s lens) hash() int {
	number := 0
	sb := []byte(s)
	for _, c := range sb {
		number = ((int(c) + number) * 17) % 256
	}
	return number
}

func main() {
	s := strings.Split(readLines()[0], ",")
	var box [256]map[string]slot
	for i := 0; i < 256; i++ {
		box[i] = make(map[string]slot)
	}
	silver := 0
	for _, w := range s {
		silver += lens(w).hash()
	}
	fmt.Println("silver: ", silver)

	le := regexp.MustCompile(`=`)
	for i, w := range s {
		if le.MatchString(w) {
			l := strings.Split(w, "=")
			h := lens(l[0]).hash()
			n, _ := strconv.Atoi(l[1])
			e,ok := box[h][l[0]]
			if ok {
				e.fl = n
				box[h][l[0]] = e
			} else {
			box[h][l[0]] = slot{n, i}
			}
		} else {
			l := strings.Split(w, "-")
			h := lens(l[0]).hash()
			delete(box[h], l[0])
		}
	}
	// fmt.Println(box)
	var boxlist [256][]slot
	for i, b := range box {
		if len(b) != 0 {
			boxlist[i] = make([]slot, 0)
			for _, v := range b {
				boxlist[i] = append(boxlist[i], v)
			}
			sort.SliceStable(boxlist[i], func(y, z int) bool { return boxlist[i][y].order < boxlist[i][z].order })
		}
	}
	// fmt.Println(boxlist)

	gold := 0
	for i := 0; i < 256; i++ {
		if len(boxlist[i]) != 0 {
			for j, s := range boxlist[i] {
				v := (i + 1)*(j+1)*s.fl
				gold += v
			}
		}
	}
	fmt.Println("gold: ",gold)

	// convert maps to slices and sort by order

}
