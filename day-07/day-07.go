package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Ouvrir le fichier
	file, err := os.Open("puzzle-input.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	tachyon := make([]string, 0)
	start_index := 0

	scanner := bufio.NewScanner(file)
	// Lire la seule ligne
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Count(line, "S") > 0 {
			start_index = strings.Index(line, "S")
			fmt.Println("Start at", start_index)
			continue
		}

		if strings.Count(line, "^") == 0 {
			continue
		}

		tachyon = append(tachyon, line)
	}

	bottom := make([]int, len(tachyon[0]))
	bottom[start_index] = 1

	result_part1 := 0
	result_part2 := 0
	
	for _, v := range tachyon {
		for i, r := range v {
			if r != '^' {
				continue
			}

			fmt.Println("Splitter at", i )
			fmt.Println("Etat du Spliter:", bottom[i])

			if bottom[i] > 0 {
				result_part1++
				
				
				if i-1 >= 0 {
					bottom[i-1] += bottom[i]
					//result_part2++
				}
				if i+1 <= len(tachyon[0]) {
					bottom[i+1] += bottom[i]
					//result_part2++
				}

				bottom[i] = 0
			}
		}
	}

	for _, r := range bottom {
		result_part2 += r
	}

	fmt.Println(bottom)
	fmt.Println("Result Part1:", result_part1)
	fmt.Println("Result Part1:", result_part2)
}