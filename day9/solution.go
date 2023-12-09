package main

// !!!! don't try this at home, your computer can explode :O
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StringSlice struct {
	words []string
}

type NumberSlice struct {
	nums []int64
}

func (sl *StringSlice) SliceInt() *NumberSlice {
	var result NumberSlice
	for _, number := range sl.words {
		v, _ := strconv.ParseInt(number, 10, 64)
		result.nums = append(result.nums, v)
	}
	return &result
}

func readLines() []string {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawData), "\n")
}

func (n *NumberSlice) isZero() bool {
	for _, number := range n.nums {
		if number != 0 {
			return false
		}
	}
	return true
}

func (n *NumberSlice) distances() *NumberSlice {
	var result NumberSlice
	for i := 1; i < len(n.nums); i++ {
		result.nums = append(result.nums, n.nums[i]-n.nums[i-1])
	}
	return &result
}

func main() {
	result_silver := int64(0)
	result_gold := int64(0)
	for _, line := range readLines() {
		first_column := NumberSlice{}
		s := StringSlice{strings.Fields(line)}
		vs := s.SliceInt()
		last := vs.nums[len(vs.nums)-1]
		first_column.nums = append(first_column.nums, vs.nums[0])
		fmt.Println(vs)
		for !vs.isZero() {
			vs = vs.distances()
			first_column.nums = append(first_column.nums, vs.nums[0])
			fmt.Println(vs)
			last += vs.nums[len(vs.nums)-1]
		}
		result_silver += last

		v := int64(0)
		for num := len(first_column.nums)-1; num >0 ; num -- {
			v = first_column.nums[num-1]-v
		}
		result_gold += v
	}
	fmt.Println("Silver: ",result_silver)
	fmt.Println("Gold: ",result_gold)



}
