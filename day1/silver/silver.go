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
		digits := regexp.MustCompile("[0-9]").FindAllString(fileScanner.Text(),-1)
	    number,_ := strconv.ParseInt(fmt.Sprintf("%s%s",digits[0],digits[len(digits)-1]),10,64)
		sum += number
		fmt.Println(sum, number)
    }
  
    readFile.Close()
}
