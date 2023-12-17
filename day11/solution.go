package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type GalaxyCloud []string

type Galaxy struct {
	x int
	y int
}

func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}

func (g *GalaxyCloud) rotateGalaxy() *GalaxyCloud {
	ng := make(GalaxyCloud, 0)
	for i := 0; i < len((*g)[0]); i++ {
		var line strings.Builder
		for j := len(*g) - 1; j >= 0; j-- {
			line.WriteByte((*g)[j][i])
		}
		ng = append(ng, line.String())
	}
	return &ng
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	var (
		cloud       GalaxyCloud
		galaxies    []Galaxy
		c           *GalaxyCloud
		fillColumns []int
		fillLines   []int
	)

	no_galaxy_line := regexp.MustCompile(`^\.+$`)
	for i, l := range readLines() {
		cloud = append(cloud, l)
		if no_galaxy_line.MatchString(l) {
			fillLines = append(fillLines, i)
		}
	}
	// filler which is never reached
	fillLines = append(fillLines, 999)
	c = &cloud
	for i, l := range *c.rotateGalaxy() {
		if no_galaxy_line.MatchString(l) {
			fillColumns = append(fillColumns, i)
		}
	}
	// filler which is never reached
	fillColumns = append(fillColumns, 999)
	fmt.Println("lines:", fillLines)
	fmt.Println("columns: ", fillColumns)
	gm := regexp.MustCompile(`#`)
	fli := 0
	// for silver use padding := 1
	// for gold use padding := 999999
	padding := 999999
	for i, l := range *c {
		f := gm.FindAllStringIndex(l, -1)
		fmt.Println("line: ", i)
		if i == fillLines[fli] {
			fmt.Println("line filled: ", i)
			fli++
			fmt.Println("fli: ", fli)
		}
		for _, j := range f {
			fci := 0

			for {
				if j[0] >= fillColumns[fci] {
					fci++
					fmt.Println("fci: ", fci)

				} else {break}
			}
			galaxies = append(galaxies, Galaxy{j[0] + fci*padding, i + fli*padding})
		}
	}
	l := uint64(0)
	for i, g := range galaxies {
		for _, n := range galaxies[i+1:] {
			d := uint64(Abs(n.x-g.x) + Abs(n.y-g.y))
			l += d
			fmt.Println(d, " ", l)
		}
	}
}
