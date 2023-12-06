package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

func readFile(filename string) []string{
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(raw), "\n")
	return lines
}

func main(){
	lines := readFile("./testinput")
	
}