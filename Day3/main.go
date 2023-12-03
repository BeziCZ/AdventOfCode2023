package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"unicode"
)

func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}

func checkLines(lines []string, idx int) int {
	var lineSum int = 0
	for _, line := range lines {
		if line[idx] >= '0' && line[idx] <= '9'{
			var bStr string = ""
			var aStr string = ""
			for j := -2; j < 0; j++{
				if line[idx+j] == '.' && aStr != "" {
					bStr = ""
					break
				} else{
					bStr += string(line[idx+j])
				}
			}
			for k := 0; k < 3; k++ {
				if line[idx+k] == '.'{
					break
				}
				aStr += string(line[idx+k])
			}
			bStr = strings.TrimLeft(bStr,".")
			var str string = ""
			if isInt(bStr){
				str = bStr + aStr
			} else{
				str = aStr
			}
			num, _ := strconv.Atoi(str)
			lineSum += num
		} else if line[idx-1] >= '0' && line[idx-1] <= '9'{
			str := line[idx-3:idx]
			str = strings.Trim(str, ".")
			num, _ := strconv.Atoi(str)
			lineSum += num
		} else if line[idx+1]>= '0' && line[idx+1] <= '9'{
			str := line[idx+1:idx+4]
			str = strings.Trim(str, ".")
			num, _ := strconv.Atoi(str)
			lineSum += num
		}
	}
	return lineSum
}
func main() {
	fmt.Println("Opening file...")
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(raw), "\n")
	var totalSum int = 0
	for i := 1; i < len(lines)-2; i++ {
		fmt.Println(lines[i])
		for idx, c := range lines[i]{
			if !(c > '0' && c <= '9') && c != '.'{
				totalSum += checkLines(lines[i-1:i+2],idx)
			}
		}
		
	}
	fmt.Println(totalSum)
}