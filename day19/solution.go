package main

import (
	"fmt"
	"os"
	"regexp"
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

type rules map[string][]string
type elements []map[string]int

func get_rules(l string, r *rules) {
	re := regexp.MustCompile(`(.*){(.*)}`)
	s := re.FindAllStringSubmatch(l, -1)
	conds := strings.Split(s[0][2], ",")
	(*r)[s[0][1]] = conds
}

func get_elements(l string, e *elements) {
	im := make(map[string]int)
	re := regexp.MustCompile(`{(.*)}`)
	s := strings.Split(re.FindAllStringSubmatch(l, -1)[0][1], ",")
	for _,i := range s {
		x := strings.Split(i,"=")
		v,_ := strconv.Atoi(x[1])
		im[x[0]]=v
	}
	*e = append(*e, im)
}

func comp(v1, v2 int , c string) bool {
	if c == "<" {
		return v1 < v2
	} else {
		return v1 > v2
	}
}

func evaluate_condition(e map[string]int, c []string) string {
	for _,i := range c {
		sc := strings.Split(i,":")
		if len(sc) == 1 {
			return sc[0]
		} else {
			k := sc[0][0:1]
			c := sc[0][1:2]
			v,_ := strconv.Atoi(sc[0][2:])
			if comp(e[k],v,c) {return sc[1]}
		}
	}
	return "FAIL"
}

func main() {
	ir := make(rules)
	var ie elements
	isRule := true
	for _, l := range readLines() {
		if l == "" {
			isRule = false
			continue
		}
		if isRule {
			get_rules(l, &ir)
		} else {
			get_elements(l, &ie)
		}
	}
fmt.Println(ie)
	result := 0
	for _,e := range ie {
		fmt.Println("checking: ",e)
		r := "in" 
		for {
			r = evaluate_condition(e,ir[r])
			if r == "A" {
				for _,v := range e {
					result +=v
					
				}
				break
			}
			if r == "R" {break}
		}

	}
	fmt.Println(result)
}
