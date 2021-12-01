package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var previousNumber int64 = -1
	increments := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		number, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if number > previousNumber && previousNumber != -1 {
			increments += 1
		}

		previousNumber = number

		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Increments", increments)
}
