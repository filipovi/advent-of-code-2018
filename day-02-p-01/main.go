package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	var nbTows = 0
	var nbThrees = 0
	for _, value := range values {
		counts := make(map[string]int)
		for index := 0; index < len(value); index++ {
			counts[string(value[index])]++
		}
		firstTwo := false
		firstThree := false
		for _, value := range counts {
			if value == 2 && !firstTwo {
				nbTows++
				firstTwo = true
			}
			if value == 3 && !firstThree {
				nbThrees++
				firstThree = true
			}
		}
	}

	fmt.Printf("%d twos\n", nbTows)
	fmt.Printf("%d threes\n", nbThrees)
	fmt.Println(nbThrees * nbTows)
}
