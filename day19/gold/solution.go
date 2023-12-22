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

type (
	Rule struct {
		category string
		operator string
		value    int
		target   string
	}
	Rules map[string][]Rule

	Range struct {
		min, max int
	}

	Input map[string]Range
)

// first input with range inside rule, second input with range outside rule 
func calc(i Input,rule) (Input,Input) {
	
}

func get_rules(l string, r *Rules) {
	re := regexp.MustCompile(`(.*){(.*)}`)
	s := re.FindAllStringSubmatch(l, -1)
	conds := strings.Split(s[0][2], ",")
	var cr []Rule
	for _, c := range conds {
		ce := strings.Split(c, ":")
		if len(ce) == 1 {
			cr = append(cr, Rule{"x", "<", 9999, ce[0]})
		} else {
			v, _ := strconv.Atoi(ce[0][2:])
			cr = append(cr, Rule{ce[0][0:1], ce[0][1:2], v, ce[1]})
		}
	}
	(*r)[s[0][1]] = cr
}

func main() {
	ir := make(Rules)
	for _, l := range readLines() {
		if l == "" {
			break
		}
		get_rules(l, &ir)
	}
	fmt.Println(ir)
}
