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

func diff_line(l1,l2 string) int {
	diff := 0
	for i := 0 ; i < len(l1); i++ {
		if l1[i] != l2[i] {diff++}
	}
	if diff == 1{
		fmt.Println("diffed line",l1)
		fmt.Println("diffed line",l2)
	}
	return diff
}

func (p *pattern) find_mirror() (bool,int) {
	smudge := false
	for i, l := range *p {
		fmt.Println("Zeile: ",l)
		if i >= 1 {
			if diff_line(l,(*p)[i-1]) <=1 {
				smudge = diff_line(l,(*p)[i-1]) == 1
				fmt.Println(i, ": ", l)
				ret := true
				for m,n := i, i-1; m <= len(*p)-1 && n >= 0 ; m,n = m+1,n-1 {
					fmt.Println("testing(",len(*p),"): ",m,"|",n)
					if diff_line((*p)[m],(*p)[n]) >1 {
						ret = false
						break
					} else if diff_line((*p)[m],(*p)[n]) == 1 {
						smudge = true
					}
				}
				if ret && smudge { return ret,i} 	
			}
		}
	}
	return false,0
}

// func (p *pattern) find_horizontal_mirror() *mirror {
// 	for i, l := range *p {
// 		if i-1 > 0 {
// 			if l == (*p)[i-1] {
// 				fmt.Println(i, ": ", l)
// 				m := mirror{p, i}
// 				return &m
// 			}
// 		}
// 	}
// 	return &mirror{p, 0}
// }

// func (m *mirror) test_horizontal_mirror() (bool, int) {
// 	if m.l == 0 {
// 		return false, 0
// 	}
// 	for i, j := m.l, m.l-1; i < len(*(m.p))-1 && j >= 0; i, j = i+1, j-1 {
// 		fmt.Println((*(m.p))[i])
// 		fmt.Println((*(m.p))[j])
// 		if (*(m.p))[i] != (*(m.p))[j] {
// 			fmt.Println("falsch")
// 			return false, 0
// 		}
// 	}
// 	fmt.Println("getestet:", m.l)
// 	return true, m.l
// }

func (p *pattern) rotate() *pattern {
	for _,l := range *p {
	fmt.Println(l)
	}
	np := make(pattern, 0)
	fmt.Println("rotating")
	for i := 0; i < len((*p)[0]); i++ {
		var line strings.Builder
		for j := len(*p)-1 ; j >=0 ; j--  {
			line.WriteByte((*p)[j][i])
		}
		np = append(np, line.String())
	}
	for _,l := range np {
		fmt.Println(l)
		}
	
	return &np
}

func main() {
	p := make([]pattern, 1)
	num := 0
	silver := 0
	for _, line := range readLines() {
		if line == "" {
			num += 1
			p = append(p, []string{})
			continue
		}
		p[num] = append(p[num], line)

	}
	fmt.Println("Anzahl: ", len(p))
	for i, t := range p {
		fmt.Println("Pattern: ",i)
		m, l := (&t).find_mirror()
		if m {
			silver += (l * 100)
			fmt.Println("hPattern: ", i, l*100)
		} else {
			m, l = (&t).rotate().find_mirror()
			if m {
				silver += l
				fmt.Println("vPattern: ", i, l)
			}

		}
	}

	fmt.Println(silver)
}
