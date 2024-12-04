package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func xmas(lines []string, i int, j int) bool {
	// up M down S
	if upMdownS(lines, i, j) {
		return true
	}

	// up S down M
	if upSdownM(lines, i, j) {
		return true
	}

	// left M right S
	if leftMrightS(lines, i, j) {
		return true
	}

	// left S right M
	if leftSrightM(lines, i, j) {
		return true
	}

	return false
}

func upMdownS(lines []string, i int, j int) bool {
	//fmt.Printf("i: %d, j: %d\n", i, j)
	if i < 1 || i + 1 >= len(lines) || j < 1 || j + 1 >= len(lines[i]) {
		return false
	}
	return lines[i - 1][j - 1] == 'M' && lines[i - 1][j + 1] == 'M' && lines[i + 1][j - 1] == 'S' && lines[i + 1][j + 1] == 'S'
}

func upSdownM(lines []string, i int, j int) bool {
	if i < 1 || i + 1 >= len(lines) || j < 1 || j + 1 >= len(lines[i]) {
		return false
	}
	return lines[i - 1][j - 1] == 'S' && lines[i - 1][j + 1] == 'S' && lines[i + 1][j - 1] == 'M' && lines[i + 1][j + 1] == 'M'
}

func leftMrightS(lines []string, i int, j int) bool {
	if i < 1 || i + 1 >= len(lines) || j < 1 || j + 1 >= len(lines[i]) {
		return false
	}
	return lines[i - 1][j - 1] == 'S' && lines[i - 1][j + 1] == 'M' && lines[i + 1][j - 1] == 'S' && lines[i + 1][j + 1] == 'M'
}

func leftSrightM(lines []string, i int, j int) bool {
	if i < 1 || i + 1 >= len(lines) || j < 1 || j + 1 >= len(lines[i]) {
		return false
	}
	return lines[i - 1][j - 1] == 'M' && lines[i - 1][j + 1] == 'S' && lines[i + 1][j - 1] == 'M' && lines[i + 1][j + 1] == 'S'
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
			if line[j] == 'A' {
				if xmas(lines, i, j) {
					count++
				}
			}

		}
	}

	fmt.Println(count)

}