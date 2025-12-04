package main

import (
	"bufio"
	"fmt"
	"os"
 	"strconv"
)

func indexOf(e int, data []int) int {
	for k, v := range data {
		if e == v {
			return k
		}
	}

	return -1 // not found
}

func Strtoi(a string) []int {
	arr := make([]int, 0)
	for _, v := range a {
		tmp, _ := strconv.Atoi( string(v) )
		arr = append(arr, tmp)
	}
	return arr
}

func findBatteries(b []int, size int) []int {

	batteriesFound := make([]int, 0)

	for size > 0 {
		// Looking for Batt to add
		for i := 9 ; i>0 ; i-- {
			index := indexOf(i, b)
			if index != -1 && index <= (len(b) - size) {
				batteriesFound = append(batteriesFound, i)
				b = b[index + 1:]
				break
			}
		}

		fmt.Println("Current Founds:", batteriesFound)

		size--
	}

	return batteriesFound


}

func convertBatteries(fb []int) int {
	str := ""
	for _, v := range fb {
		str = str + strconv.Itoa(v)
	}
	tmp, _ := strconv.Atoi(str)
	fmt.Println("Converted", tmp)
	return tmp
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
	banks := make([][]int, 0)

	// Lire la seule ligne
	for scanner.Scan() {
		line := scanner.Text()
		banks = append(banks, Strtoi(line))
	}

	fmt.Println(banks)

	result_part1 := 0
	result_part2 := 0
	for _, b := range banks {
		fb := findBatteries(b, 2)
		result_part1 += convertBatteries(fb)
		fb = findBatteries(b, 12)
		result_part2 += convertBatteries(fb)
	}

	fmt.Println(result_part1)
	fmt.Println(result_part2)

}