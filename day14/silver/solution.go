package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}

type rocks []string

func (r *rocks) rotateGalaxy() *rocks {
	nr := make(rocks, 0)
	for i := 0; i < len((*r)[0]); i++ {
		var line strings.Builder
		for j := len(*r) - 1; j >= 0; j-- {
			line.WriteByte((*r)[j][i])
		}
		nr = append(nr, line.String())
	}
	return &nr
}


func main() {
var m rocks
	for _,l := range readLines() {
		m = append(m, l)
	}
	// working in lines is easier
	mp := &m
	nr := mp.rotateGalaxy()


	silver:=0
	for _,l := range *nr {
		weight := 100
		for i:= len(l)-1; i >=0; i-- {
			if l[i] == '.' {
				continue
			}
			if l[i] == 'O' {
				silver +=weight
				weight--
			}
			if l[i] == '#' {
				weight = i
			}


		}

	}
	fmt.Println(silver)
}
