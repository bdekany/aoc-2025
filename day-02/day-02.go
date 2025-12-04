package main

import (
	"bufio"
	"fmt"
	"os"
 	"strconv"
	"strings"
)

func rangeToInt(a string) (int, int) {
	min, _ := strconv.Atoi(strings.Split(a, "-")[0])
	max, _ := strconv.Atoi(strings.Split(a, "-")[1])
	return min, max
}

func testMirror(a string) bool {
	max_len := len(a)
	k := max_len/2
	if a[0:k] != a[k:] {
		return false
	}

	return true
}

func testBlock(a string, size int) bool {
	max_len := len(a)

	if max_len < 6 || max_len == 7 {
		return false
	}

	for i:=0; i<(max_len-size); i+=size {
		start := i
		mid := i + size
		end := i + size * 2
		if a[start:mid] != a[mid:end] {
			return false
		}
	}
	return true
}

func testAll(a string) bool {

	if len(a) == 1 {
		return false
	}

	for i, _ := range a {
		if i+1 == len(a) {
			return true
		}

		if a[i] != a[i+1] {
			return false
		}
	}
	return true
}

func findInvalid(a, b int) (int, int) {
	count := 0
	count_part2 := 0
	for i := a ; i <= b ; i++ {
		iTxt := strconv.Itoa(i)
		if len(iTxt)%2 == 0 {
			if testMirror(iTxt) {
				count += i
				count_part2 += i
			} else if testBlock(iTxt, 2){
				count_part2 += i
			}
		} else {
			if testAll(iTxt){
				count_part2 += i
			} else if testBlock(iTxt, 3){
				count_part2 += i
			}
		}

	}
	return count, count_part2
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
	ranges := make([]string, 0)

	// Lire la seule ligne
	for scanner.Scan() {
		line := scanner.Text()
		ranges = strings.Split(line, ",")
	}

	// Trouver les invalides
	invalid_part1 := 0
	invalid_part2 := 0
	for _, r := range ranges {
		fmt.Println("Find for:", r)
		min, max := rangeToInt(r)
		miroir, all := findInvalid(min, max)
		invalid_part1 += miroir
		invalid_part2 += all
	}

	fmt.Println("Result 1:", invalid_part1)
	fmt.Println("Result 2:", invalid_part2)
}