package main

import (
	"bufio"
	"strings"
	//"fmt"
	"log"
	"os"
	"math"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func isSafe(report []string) bool {
	first_increasing := false
	cur_increasing := false

	for level := range(len(report)) - 1 {
		first, err := strconv.ParseFloat(report[level], 64)
		check(err)

		if (first == -1) {
			continue
		}
	
		second, err := strconv.ParseFloat(report[level+1], 64)
		check(err)

		if level == 0 {
			first_increasing = first < second
			cur_increasing = first < second
		} else {
			cur_increasing = first < second
		}
	
		diff := math.Abs(first - second)
		if (diff == 0 || diff < 1 || diff > 3) {
			return false
		}
		
		if first_increasing != cur_increasing {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
    check(err)
	defer file.Close()

	var reports []string
	
	scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
		reports = append(reports, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	number_safe := 0

	for i:= range len(reports) {
		safe := false
		report := strings.Split(reports[i], " ")

		safe = isSafe(report)

		if !safe {
			for level := range len(report) {
				val := report[level]
				report[level] = "-1"
				//log.Println(report)

				trunc_report := make([]string, len(report) - 1)

				i := 0
				j := 0
				for i < len(report) {
					if report[i] != "-1" {
						//log.Printf("i: %d, report[i]: %s", i, report[i])
						trunc_report[j] = report[i]
						j++
						//log.Println(trunc_report)
					}
					i++
				}
				safe = isSafe(trunc_report) || safe
				report[level] = val
			}
		}

		if safe {
			number_safe++
		}
	}

	log.Println(number_safe)

}