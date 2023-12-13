package main

import (
	"fmt"
	"os"
	"strings"
)

type pattern []string

type mirror struct {
	p *pattern
	l int
}
func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}


func (p *pattern) find_horizontal_mirror() *mirror {
	for i, l := range *p {
		if i-1 > 0 {
			if l == (*p)[i-1] {
				fmt.Println(i,": ",l)
				m := mirror{p,i}
				return  &m
			}
		}
	}
	return &mirror{p,0}
}

func (m *mirror) test_horizontal_mirror() (bool,int) {
	if m.l == 0 {
		return false,0
	}
	for i,j := m.l,m.l-1 ; i < len(*(m.p))-1 && j >= 0 ; i,j = i+1,j-1 {
		fmt.Println((*(m.p))[i])
		fmt.Println((*(m.p))[j])
		if (*(m.p))[i] != (*(m.p))[j] {
			fmt.Println("falsch")
			return false,0
		}
	}
	fmt.Println("getestet:",m.l)
	return true,m.l
}

func (p *pattern) rotate() *pattern {
	np := make(pattern, 0)

	for i := 0 ; i < len((*p)[0])-1 ; i++ {
		var line strings.Builder
		 for _,j := range (*p) {
			line.WriteByte(j[i])
		 }
		 np = append(np, line.String())
	} 
	return &np
}

func main() {
	p := make([]pattern, 1)
	num := 0
	silver :=0
	for _, line := range readLines() {
		if line == "" {
			num += 1
			p = append(p, []string{})
			continue
		}
		p[num] = append(p[num], line)

	}
	for i, t := range p {
		m,l := (&t).find_horizontal_mirror().test_horizontal_mirror()
		if m {
			silver += (l*100)
			fmt.Println("hPattern: ",i,l*100)
		} else {
			m,l = (&t).rotate().find_horizontal_mirror().test_horizontal_mirror()
			if m {
				silver += l
				fmt.Println("vPattern: ",i,l)
			}
		}
	}

	fmt.Println(silver)
}
