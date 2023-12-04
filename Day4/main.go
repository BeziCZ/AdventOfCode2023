package main

import (
	"fmt"
	"os"
	//"bufio"
	"strings"
	"log"
	"slices"
	//"io"
	"regexp"
	"io/ioutil"
)

func getCardPoints(line string) int{
	numbers := strings.Split(line, "|")
	re := regexp.MustCompile("[0-9]+")
	winning := re.FindAllString(numbers[0],-1)
	random := re.FindAllString(numbers[1],-1)
	var worth int = 0 
	for _, win := range random{
		if slices.Contains(winning,win){ 
			
			if worth == 0{
				worth++
			} else {
				worth *= 2
			}
		}
	}
	return worth
}

func getNumberOfNextCards(line string) int{
	numbers := strings.Split(line, "|")
	re := regexp.MustCompile("[0-9]+")
	winning := re.FindAllString(numbers[0],-1)
	random := re.FindAllString(numbers[1],-1)
	var numOfNext int = 0 
	for _, win := range random{
		if slices.Contains(winning,win){
			numOfNext++
		}
	}
	return numOfNext
}
func main() {
	
	fmt.Println("Opening file...")
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	/* Used for part 1
	reader := bufio.NewReader(file)
	var totalValue int = 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil{
			if err == io.EOF {

				break
			} else {
				log.Fatal(err)
			}
		}
		split_line := strings.Split(line,":")
		totalValue += getCardPoints(split_line[1])
	}
	fmt.Println(totalValue)
	*/ 
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(raw), "\n")
	numsOfCards := make([]int,198,198)
	for i := 0; i < len(numsOfCards); i++ {
		numsOfCards[i]++
	}
	fmt.Println(len(numsOfCards))
	fmt.Println(len(lines))
	for idx, line := range lines{
		split_line := strings.Split(line,":")
		for i := 0; i < numsOfCards[idx]; i++{
			numOfNext := getNumberOfNextCards(split_line[1])
			for i := 1; i <= numOfNext; i++ {
				numsOfCards[idx+i]++ 
			}
		}
	}
	fmt.Println(numsOfCards)
	var totalNumberOfCards int = 0
	for _, n := range numsOfCards{
		totalNumberOfCards += n
	}
	fmt.Println(totalNumberOfCards)
}