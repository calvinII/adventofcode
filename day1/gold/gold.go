package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)
 
func main() {
 
    readFile, err := os.Open("data.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
    var sum int64 = 0
    for fileScanner.Scan() {
        numberMap := make(map[string]string)
        numberMap["one"] = "o1e"
        numberMap["two"] = "t2o"
        numberMap["three"] = "t3e"
        numberMap["four"] = "f4r"
        numberMap["five"] = "f5e"
        numberMap["six"] = "s6x"
        numberMap["seven"] = "s7n"
        numberMap["eight"] = "e8t"
        numberMap["nine"] = "n9e"
        word := fileScanner.Text()
        for key, element := range numberMap {
            word = regexp.MustCompile(key).ReplaceAllString(word, element)
        }
		digits := regexp.MustCompile("[0-9]").FindAllString(word,-1)
	    number,_ := strconv.ParseInt(fmt.Sprintf("%s%s",digits[0],digits[len(digits)-1]),10,64)
		sum += number
		fmt.Println(sum, number)
    }
  
    readFile.Close()
}
