package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

// direction in coords
type dir struct {
	x int
	y int
}

type pipe struct {
	c []byte
}

// https://rosettacode.org/wiki/Shoelace_formula_for_polygonal_area
// Hier berechnen wir die Fläche inclusive der Linien drumrum ...

func shoelace(pts []coord) int {
    sum := 0
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


func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}

var field [][]byte

func main() {
	var s coord
	var pipeLoop []coord

	pipes := make(map[byte]pipe)
	// the opposite
	op := map[byte]byte{'N': 'S', 
						'E': 'W', 
						'S': 'N', 
						'W': 'E'}
	pipes[byte('|')] = pipe{[]byte("SN")}
	pipes[byte('-')] = pipe{[]byte("WE")}
	pipes[byte('L')] = pipe{[]byte("NE")}
	pipes[byte('J')] = pipe{[]byte("NW")}
	pipes[byte('7')] = pipe{[]byte("WS")}
	pipes[byte('F')] = pipe{[]byte("SE")}

	// directing to next pipe
	dirs := make(map[byte]dir)

	dirs[byte('N')] = dir{0,-1}
	dirs[byte('E')] = dir{1,0}
	dirs[byte('S')] = dir{0,1}
	dirs[byte('W')] = dir{-1,0}


	for _, line := range readLines() {
		bytes := []byte(line)
		field = append(field, bytes)
	}
	fmt.Println(field)

	// find starting Point
	for i,l := range field {
		x := bytes.IndexByte(l,byte('S'))
		if  x != -1 {
			s.x =x
			s.y =i
			break
		}
	}
	fmt.Println(s)
	pipeLoop = append(pipeLoop, coord{s.x,s.y})

	// get the 4 connects around S and show the first match
	var (
		x int
		y int
	)
	var d byte
	for _,d = range []byte("NESW"){
		y = s.y+dirs[d].y
		x = s.x+dirs[d].x
		e := field[y][x]
		if pipes[e].c[0] == op[d] || pipes[e].c[1] == op[d] {
			break
		}
	}
	pipeLoop = append(pipeLoop, coord{x,y})
// x y op[d]

	steps :=1
	direction := d
	// idea we get the next pipe and eleminate the connected direction (at the end of the loop we have the next pipe and)
	for x != pipeLoop[0].x || y != pipeLoop[0].y {
		nd := pipes[field[y][x]]
		if nd.c[0] == op[direction] {
			direction = nd.c[1]
		} else {
			direction = nd.c[0]
		}
		x = x + dirs[direction].x
		y = y + dirs[direction].y
		fmt.Println(x,y,field[y][x])
		pipeLoop = append(pipeLoop, coord{x,y})
		steps += 1
	}
	fmt.Println(pipeLoop)
	fmt.Println("silver: ",steps/2)

	//https://de.wikipedia.org/wiki/Satz_von_Pick 
	// damit rechnen wir die Linie noch raus
    fmt.Println("Gold: ", shoelace(pipeLoop)-steps/2+1)

	
}
