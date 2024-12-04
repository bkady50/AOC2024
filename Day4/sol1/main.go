package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func xmas(lines []string, i int, j int) int {
	count := 0

	if isUp(lines, i, j) {
		count++
	}

	if isDown(lines, i, j) {
		count++
	}

	if isLeft(lines, i, j) {
		count++
	}

	if isRight(lines, i, j) {
		count++
	}

	if isUpRight(lines, i, j) {
		count++
	}

	if isDownRight(lines, i, j) {
		count++
	}

	if isUpLeft(lines, i, j) {
		count++
	}

	if isDownLeft(lines, i, j) {
		count++
	}

	return count
}

func isUp(lines []string, i int, j int) bool {
	if i < 3 {
		return false
	}
	return lines[i - 1][j] == 'M' && lines[i - 2][j] == 'A' && lines[i - 3][j] == 'S'
}

func isDown(lines []string, i int, j int) bool {
	if i + 3 >= len(lines) {
		return false
	}
	return lines[i + 1][j] == 'M' && lines[i + 2][j] == 'A' && lines[i + 3][j] == 'S'
}

func isLeft(lines []string, i int, j int) bool {
	if j < 3 {
		return false
	}
	return lines[i][j - 1] == 'M' && lines[i][j - 2] == 'A' && lines[i][j - 3] == 'S'
}

func isRight(lines []string, i int, j int) bool {
	if j + 3 >= len(lines[i]) {
		return false
	}
	return lines[i][j + 1] == 'M' && lines[i][j + 2] == 'A' && lines[i][j + 3] == 'S'
}

func isUpRight(lines []string, i int, j int) bool {
	if i < 3 || j + 3 >= len(lines[i]) {
		return false
	}
	return lines[i - 1][j + 1] == 'M' && lines[i - 2][j + 2] == 'A' && lines[i - 3][j + 3] == 'S'
}

func isUpLeft(lines []string, i int, j int) bool {
	if i < 3 || j < 3 {
		return false
	}
	return lines[i - 1][j - 1] == 'M' && lines[i - 2][j - 2] == 'A' && lines[i - 3][j - 3] == 'S'
}

func isDownRight(lines []string, i int, j int) bool {
	if i + 3 >= len(lines)|| j + 3 >= len(lines[i]) {
		return false
	}
	//fmt.Printf("i: %d, j: %d\n", i, j)
	return lines[i + 1][j + 1] == 'M' && lines[i + 2][j + 2] == 'A' && lines[i + 3][j + 3] == 'S'
}

func isDownLeft(lines []string, i int, j int) bool {
	if i + 3 >= len(lines) || j < 3 {
		return false
	}
	return lines[i + 1][j - 1] == 'M' && lines[i + 2][j - 2] == 'A' && lines[i + 3][j - 3] == 'S'
}


func main () {
	file, err := os.Open("input.txt")
	if err != nil {
        panic(err)
    }
	defer file.Close()

	var lines []string
	
	scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
		lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	count := 0

	for i := range(len(lines)) {
		line := lines[i]

		for j := range(len(line)) {
			if line[j] == 'X' {
				count += xmas(lines, i, j)
			}

		}
	}

	fmt.Println(count)

}