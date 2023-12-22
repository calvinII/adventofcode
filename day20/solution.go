package main

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

type connector struct {
	c string
	v int
}

type conn struct {
	sources *[]connector
	kind    string
	state int
	targets *[]connector
}

func parse(sb conn, out conn, c map[string]conn, steps int) uint64 {
	result := uint64(0) 
	for i := 1; i < steps ; i++ {

	}
}

func main() {
	connLine := make(map[string]conn)
	re := regexp.MustCompile(`,`)
	for _, l := range readLines() {
		s := strings.Fields(re.ReplaceAllString(l, ""))
		fmt.Println(s)
		kind := s[0][0:1]
		name := s[0][1:]
		tnames := new([]connector)
		snames := new([]connector)
		for _,i := range s[2:] {
			*tnames = append(*tnames, connector{i,0})
		}
		connLine[name] = conn{snames,kind,0,tnames}
	}

	for k, c := range connLine {
		for _, t := range *c.targets {
			_, ok := connLine[t.c]
			if ok {
				*connLine[t.c].sources = append(*connLine[t.c].sources, connector{k,0})
			} else {
				s0 := new([]connector)
				t0 := new([]connector)
				connLine[t.c] = conn{s0,"o",0,t0}
				fmt.Println("Ausgang: ",t)
			}
		}
	}

	for k, c := range connLine {
		fmt.Println(k, ": ", c.sources, " | ", c.kind, " | ", c.state, " | ", c.targets)
	}
}
