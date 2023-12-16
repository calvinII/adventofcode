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

func get_distance(race_time int64, record int64) int64 {
	count := int64(0)
	for press_time := int64(1); press_time <= race_time; press_time++ {
		travel := (race_time - press_time) * press_time
		if travel > record {
			count++
		}
	}
	return count
}

func main() {
	times := strings.Fields(readLines()[0])
	records := strings.Fields(readLines()[1])
	result :=int64(1)
	for i := 1; i < len(times); i++ {
		time,_ := strconv.ParseInt(times[i],10,64)
		record,_ := strconv.ParseInt(records[i],10,64)
		result *= get_distance(time,record)
	}
	fmt.Println("silver: ",result)
	new_time,_ := strconv.ParseInt(strings.Join(times[1:],""),10,64)
	new_record,_ := strconv.ParseInt(strings.Join(records[1:],""),10,64)

	fmt.Println(new_time,new_record)
	fmt.Println(get_distance(new_time,new_record))

}
