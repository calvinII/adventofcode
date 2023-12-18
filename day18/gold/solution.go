package main

import (
	"fmt"
	"os"
	"strconv"
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
	x int64
	y int64
}

type coords []coord

func (c *coord) new_coord(s int, d uint64) *coord {
	direction := map[int]coord{
		2: {-1, 0},
		1: {0, 1},
		3: {0, -1},
		0: {1, 0},
	}
	return &coord{c.x + direction[s].x*int64(d), c.y + direction[s].y*int64(d)}
}

func get_steps(l []string) (*coords, uint64) {
	s_coord := &coord{0, 0}
	var s coords

	steps := uint64(0)

	for _, l := range l {
		z := strings.Fields(l)
		p,_ := strconv.Atoi(z[2][7:8])
		d, _ := strconv.ParseUint(z[2][2:7],16,64)
		np := *s_coord.new_coord(p, d)
		s = append(s, np)
		steps += uint64(d)
		s_coord = &s[len(s)-1]
	}
	return &s, steps
}

// https://rosettacode.org/wiki/Shoelace_formula_for_polygonal_area
// Hier berechnen wir die Fläche inclusive der Linien drumrum ...

func shoelace(pts []coord) int64 {
	sum := int64(0)
	p0 := pts[len(pts)-1]
	for _, p1 := range pts {
		sum += p0.y*p1.x - p0.x*p1.y
		p0 = p1
	}
	if sum < 0 {
		sum = -sum
	}
	return sum / 2
}

func main() {
	s, st := get_steps(readLines())
	fmt.Println(s)
	fmt.Println("gold: ", uint64(shoelace(*s))+st/2+1)
	fmt.Println("len: ", st)
}
