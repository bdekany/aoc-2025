package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
 	"strconv"
	"regexp"
)

func Strtoi(a string) []int {
	arr := make([]int, 0)
	for _, v := range strings.Split(a, " ") {
		tmp, _ := strconv.Atoi( string(v) )
		arr = append(arr, tmp)
	}
	return arr
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
	problems := make([][]int, 0)
	operators := make([]string, 0)

	worksheet := make([]string, 0)

	re := regexp.MustCompile(`\s+`)
	// Lire la seule ligne
	for scanner.Scan() {
		line := scanner.Text()
		// Stockage pour part2
		worksheet = append(worksheet, line)

		// Nettoyage des withspaces Part1
		line = strings.TrimSpace(line)
		line = re.ReplaceAllString(line, " ")
		if strings.Contains(line, "*") {
			operators = strings.Split(line, " ")
		} else {
			problems = append(problems, Strtoi(line))
		}
	}

	// fmt.Println(problems)
	// fmt.Println(operators)

	// Initial result, 0 pour les additions, 1 pour les *
	results := make([]int, 0)
	for _, o := range operators {
		if o == "*" {
			results = append(results, 1)
		} else {
			results = append(results, 0	)
		}
	}

	for _, v := range problems {
		for col, n := range v {
			if operators[col] == "*" {
				results[col] *= n
			} else {
				results[col] += n
			}
		}
	}

	// fmt.Println(results)

	result_part1 := 0
	for _, v := range results {
		result_part1 += v
	}

	fmt.Println("Resultat Part 1:", result_part1)

	size := len(worksheet)
	tmp := ""
	op := ""
	res := 0

	re_part2 := regexp.MustCompile(`^\s+$`)
	result_part2 := 0

	for j, _ := range worksheet[0] {

		for k := 0 ; k < size ; k++ {
			tmp = tmp + string(worksheet[k][j])
		}

		// Resultat complet
		if re_part2.MatchString(tmp) {
			// fmt.Println("Fin:", res)
			result_part2 += res
			op = ""
			res = 0
			continue
		}

		// Nouveau probleme
		if strings.ContainsAny(tmp, "*+") {
			if tmp[len(tmp)-1:] == "*" {
				op = "*"
				res = 1
			} else {
				op = "+"
				res = 0
			}
			tmp = tmp[:len(tmp)-1]
		}

		if op == "*" {
			z, _ := strconv.Atoi(strings.TrimSpace(tmp))
			res *= z
		} else {
			z, _ := strconv.Atoi(strings.TrimSpace(tmp))
			res += z
		}

		tmp = ""
	}

	result_part2 += res
	fmt.Println("Resultat Part 2:", result_part2)

}