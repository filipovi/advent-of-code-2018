package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func retrieve(values []string) int {
	var result int
	sums := map[int]bool{0: false}

	for {
		for _, value := range values {
			value, _ := strconv.Atoi(value)
			result += value
			if sums[result] {
				return result
			}

			sums[result] = true
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var values []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := retrieve(values)

	fmt.Println(result)
}
