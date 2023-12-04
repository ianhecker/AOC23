package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ianhecker/aoc23/one"
)

func readLines(filepath string) []string {
	file, err := os.Open(filepath)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	solveOne()
}

func solveOne() {
	lines := readLines("./one/input.txt")

	var total int
	for _, line := range lines {

		i, err := one.ParseStringForDigit(line)
		check(err)

		fmt.Printf("parsed:'%s' and found digit:%d\n", line, i)

		total += i
	}

	fmt.Printf("total is: %d\n", total)
}
