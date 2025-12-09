package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeRoll(rolls_map []string) (int, []string) {
	min_x := 0
	max_x := len(rolls_map[0])
	min_y := 0
	max_y := len(rolls_map)

	remove_counter := 0
	next_map := make([]string, len(rolls_map))
	copy(next_map, rolls_map)

	for y := min_y ; y < max_y ; y++ {
		for x := min_x ; x < max_x ; x++ {
			// est-ce un roll?
			if rolls_map[y][x] == '.' {
				continue
			}

			proximity_str := ""
			// fmt.Println("y, x", y, x)

			// creation d'un str de vérification
			if (y-1) >= min_y {
				proximity_str += rolls_map[y-1][max(x-1, min_x):min(x+2, max_x)]
			}
			proximity_str += rolls_map[y][max(x-1, min_x):min(x+2, max_x)]
			if (y+1) < max_y {
				proximity_str += rolls_map[y+1][max(x-1, min_x):min(x+2, max_x)]
			}

			// fmt.Println(proximity_str)
			// fmt.Println(strings.Count(proximity_str, "@"))

			// Moins de 4 voisins + soi meme
			if strings.Count(proximity_str, "@") < 5 {
				remove_counter++
				next_map[y] = next_map[y][:x] + "." + next_map[y][x+1:]
			}

		}
	}

	return remove_counter, next_map
}

func main() {

	// Ouvrir le fichier
	file, err := os.Open("puzzle-input.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rolls_map := make([]string, 0)
	next_map := make([]string, 0)

	// Lire la seule ligne
	for scanner.Scan() {
		line := scanner.Text()
		rolls_map = append(rolls_map, line)
	}

	result_part1 := 0
	result_part2 := 0

	result_part1, next_map = removeRoll(rolls_map)

	fmt.Println("Resultat Part1", result_part1)

	i := result_part1
	result_part2 += i
	for i > 0 {
		i, next_map = removeRoll(next_map)
		result_part2 += i
	}

	fmt.Println("Resultat Part2", result_part2)
}