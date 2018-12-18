package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../day-02-p-01/input.txt")
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

	var nb int
	var result string
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if len(values[i]) != len(values[j]) {
				continue
			}

			var p string
			var nbp int
			for index := 0; index < len(values[i]); index++ {
				if values[i][index] == values[j][index] {
					nbp++
					p += string(values[i][index])
				}
				if nbp > nb {
					nb = nbp
					result = p
				}
			}
		}
	}
	fmt.Println(result)
}
