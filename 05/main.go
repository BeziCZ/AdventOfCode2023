package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func readFile(filename string) []string{
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseLines()
func getWantedSeeds(line string) []int{
	seeds := make([]int,0)
	seedNums := strings.Split(strings.Trim(strings.Split(line, ":")[1]," "), " ")
	for _, seed := range seedNums {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}
	return seeds
}
func main() {
	lines := readFile("testinput")
	wantedSeeds := getWantedSeeds(lines[0])
	fmt.Println(wantedSeeds)
	var seed2soil map[int]int

}